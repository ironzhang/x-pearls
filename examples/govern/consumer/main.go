package main

import (
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/ironzhang/x-pearls/examples/govern/endpoint"
	"github.com/ironzhang/x-pearls/govern"
	"github.com/ironzhang/x-pearls/govern/etcdv2"
	"github.com/ironzhang/x-pearls/zlog"
)

type Options struct {
	Level int
}

func (o *Options) Parse() {
	flag.IntVar(&o.Level, "level", int(zlog.INFO), "log level")
	flag.Parse()
}

func main() {
	var opts Options
	opts.Parse()
	zlog.Default.SetLevel(zlog.Level(opts.Level))

	d, err := govern.Open(etcdv2.DriverName, "test", client.Config{Endpoints: []string{"http://127.0.0.1:2379"}})
	if err != nil {
		zlog.Fatalw("open", "error", err)
	}
	defer d.Close()
	defer time.Sleep(time.Second)

	p := d.NewConsumer("ac-test", &endpoint.Endpoint{}, refresh)
	defer p.Close()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func refresh(endpoints []govern.Endpoint) {
	zlog.Infow("refresh", "endpoints", endpoints)
}
