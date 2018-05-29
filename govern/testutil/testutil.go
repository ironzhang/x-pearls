package testutil

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ironzhang/x-pearls/govern"
)

var Endpoints = []govern.Endpoint{
	&Endpoint{Name: "node0", Addr: "localhost:2000"},
	&Endpoint{Name: "node1", Addr: "localhost:2001"},
	&Endpoint{Name: "node2", Addr: "localhost:2002"},
	&Endpoint{Name: "node3", Addr: "localhost:2003"},
}

func waitProvider(p govern.Provider, node string, list func(string) ([]govern.Endpoint, error)) (govern.Endpoint, error) {
	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		eps, err := list(p.Directory())
		if err != nil {
			return nil, err
		}
		for _, ep := range eps {
			if ep.Node() == node {
				return ep, nil
			}
		}
	}
	return nil, errors.New("timeout")
}

func waitConsumer(c govern.Consumer, n int) error {
	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		if len(c.GetEndpoints()) == n {
			return nil
		}
	}
	return errors.New("timeout")
}

func waitRefresher(r *Refresher, n int) error {
	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		if len(r.Endpoints) == n {
			return nil
		}
	}
	return errors.New("timeout")
}

func TestProvider(t *testing.T, p govern.Provider, ep govern.Endpoint, list func(string) ([]govern.Endpoint, error)) {
	val, err := waitProvider(p, ep.Node(), list)
	if err != nil {
		t.Fatalf("%s: wait provider: %v", p.Driver(), err)
	}
	if got, want := val, ep; !got.Equal(want) {
		t.Fatalf("%s: endpoint: got %v, want %v", p.Driver(), got, want)
	} else {
		t.Logf("%s: endpoint: got %v", p.Driver(), got)
	}
}

func TestConsumer(t *testing.T, c govern.Consumer, register func(string, govern.Endpoint) error) {
	for _, ep := range Endpoints {
		if err := register(c.Directory(), ep); err != nil {
			t.Fatalf("%s: register: %v", c.Driver(), err)
		}
	}
	if err := waitConsumer(c, len(Endpoints)); err != nil {
		t.Fatalf("%s: wait consumer: %v", c.Driver(), err)
	}
	if got, want := c.GetEndpoints(), Endpoints; !reflect.DeepEqual(got, want) {
		t.Fatalf("%s: endpoints: got %v, want %v", c.Driver(), got, want)
	} else {
		t.Logf("%s: endpoints: got %v", c.Driver(), got)
	}
}

func TestConsumerRefresh(t *testing.T, c govern.Consumer, r *Refresher, register func(string, govern.Endpoint) error) {
	for _, ep := range Endpoints {
		if err := register(c.Directory(), ep); err != nil {
			t.Fatalf("%s: register: %v", c.Driver(), err)
		}
	}
	if err := waitRefresher(r, len(Endpoints)); err != nil {
		t.Fatalf("%s: wait refresher: %v", c.Driver(), err)
	}
	if got, want := r.Endpoints, Endpoints; !reflect.DeepEqual(got, want) {
		t.Fatalf("%s: endpoints: got %v, want %v", c.Driver(), got, want)
	} else {
		t.Logf("%s: endpoints: got %v", c.Driver(), got)
	}
}

func TestDriver(t *testing.T, d govern.Driver) {
	sv := "TestService"

	var r Refresher
	c := d.NewConsumer(sv, &Endpoint{}, r.Refresh)
	defer c.Close()

	for _, ep := range Endpoints {
		x := ep
		p := d.NewProvider(sv, 10*time.Second, func() govern.Endpoint { return x })
		defer p.Close()
	}

	if err := waitRefresher(&r, len(Endpoints)); err != nil {
		t.Fatalf("%s: wait refresher: %v", d.Name(), err)
	}
	if got, want := r.Endpoints, Endpoints; !reflect.DeepEqual(got, want) {
		t.Fatalf("%s: endpoints: got %v, want %v", d.Name(), got, want)
	} else {
		t.Logf("%s: endpoints: got %v", d.Name(), got)
	}
}
