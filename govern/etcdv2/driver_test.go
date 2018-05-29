package etcdv2

import (
	"testing"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/ironzhang/x-pearls/govern"
	"github.com/ironzhang/x-pearls/govern/testutil"
)

func OpenTestDriver(namespace string) govern.Driver {
	driver, err := Open(namespace, client.Config{Endpoints: []string{"http://127.0.0.1:2379"}})
	if err != nil {
		panic(err)
	}
	return driver
}

func TestDriver(t *testing.T) {
	defer time.Sleep(500 * time.Millisecond)

	d := OpenTestDriver("TestDriver")
	defer d.Close()

	testutil.TestDriver(t, d)
}
