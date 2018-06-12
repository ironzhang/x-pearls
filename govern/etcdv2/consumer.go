package etcdv2

import (
	"context"
	"sync"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/ironzhang/x-pearls/govern"
	"github.com/ironzhang/x-pearls/govern/etcdv2/etcdapi"
	"github.com/ironzhang/x-pearls/log"
)

type consumer struct {
	api     etcdapi.API
	dir     string
	refresh govern.RefreshEndpointsFunc

	endpoints govern.Endpoints
	done      chan struct{}
	mu        sync.RWMutex
	list      []govern.Endpoint
}

func newConsumer(api client.KeysAPI, dir string, endpoint govern.Endpoint, refresh govern.RefreshEndpointsFunc) *consumer {
	return new(consumer).init(api, dir, endpoint, refresh)
}

func (c *consumer) init(api client.KeysAPI, dir string, endpoint govern.Endpoint, refresh govern.RefreshEndpointsFunc) *consumer {
	c.api.Init(api, endpoint)
	c.dir = dir
	c.refresh = refresh
	c.endpoints = make(govern.Endpoints)
	c.done = make(chan struct{})
	go c.watching(c.done)
	return c
}

func (c *consumer) Driver() string {
	return DriverName
}

func (c *consumer) Directory() string {
	return c.dir
}

func (c *consumer) Close() error {
	close(c.done)
	return nil
}

func (c *consumer) GetEndpoints() []govern.Endpoint {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.list
}

func (c *consumer) watching(done <-chan struct{}) {
	log.Infow("start watch", "dir", c.dir)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-done
		cancel()
	}()

	eps, index, err := c.listEndpoints(ctx)
	if err != nil {
		log.Infow("stop watch", "dir", c.dir)
		return
	}
	c.setup(eps)

	w := c.api.Watcher(c.dir, index)
	for {
		event, err := c.watchNext(ctx, w)
		if err != nil {
			break
		}
		c.update(event)
	}

	log.Infow("stop watch", "dir", c.dir)
}

func (c *consumer) listEndpoints(ctx context.Context) ([]govern.Endpoint, uint64, error) {
	const min, max = time.Second, 60 * time.Second

	delay := min
	for {
		eps, index, err := c.api.Get(ctx, c.dir)
		if err == nil {
			return eps, index, nil
		} else if e, ok := err.(client.Error); ok && e.Code == client.ErrorCodeKeyNotFound {
			return nil, index, nil
		} else if err == context.Canceled {
			return nil, 0, err
		} else {
			log.Warnw("list endpoints", "dir", c.dir, "delay", delay, "error", err)
			time.Sleep(delay)
			if delay *= 2; delay > max {
				delay = max
			}
		}
	}
}

func (c *consumer) watchNext(ctx context.Context, w *etcdapi.Watcher) (etcdapi.Event, error) {
	const min, max = 5 * time.Millisecond, time.Second

	delay := min
	for {
		event, err := w.Next(ctx)
		if err == nil {
			return event, nil
		} else if err == context.Canceled {
			return event, err
		} else {
			log.Warnw("watch next", "dir", c.dir, "delay", delay, "error", err)
			time.Sleep(delay)
			if delay *= 2; delay > max {
				delay = max
			}
		}
	}
}

func (c *consumer) setup(eps []govern.Endpoint) {
	log.Debugw("setup", "dir", c.dir, "endpoints", eps)
	for _, ep := range eps {
		c.endpoints.Add(ep)
	}
	c.doRefresh(c.endpoints.SortList())
}

func (c *consumer) update(event etcdapi.Event) {
	log.Debugw("update", "dir", c.dir, "event", event)
	switch event.Action {
	case "set", "update":
		if c.endpoints.Add(event.Endpoint) {
			c.doRefresh(c.endpoints.SortList())
		}
	case "delete", "expire":
		if c.endpoints.Remove(event.Name) {
			c.doRefresh(c.endpoints.SortList())
		}
	}
}

func (c *consumer) doRefresh(eps []govern.Endpoint) {
	c.mu.Lock()
	c.list = eps
	c.mu.Unlock()
	if c.refresh != nil {
		c.refresh(eps)
	}
}
