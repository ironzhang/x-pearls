package testutil

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ironzhang/x-pearls/govern"
)

func Wait(r *Refresher, n int) error {
	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		if len(r.Endpoints) == n {
			return nil
		}
	}
	return errors.New("timeout")
}

func TestDriver(t *testing.T, d govern.Driver) {
	sv := "TestService"

	var r Refresher
	c := d.NewConsumer(sv, &Endpoint{}, r.Refresh)
	defer c.Close()

	endpoints := []govern.Endpoint{
		&Endpoint{Name: "node0", Addr: "localhost:2000"},
		&Endpoint{Name: "node1", Addr: "localhost:2001"},
		&Endpoint{Name: "node2", Addr: "localhost:2002"},
		&Endpoint{Name: "node3", Addr: "localhost:2003"},
	}
	for _, ep := range endpoints {
		x := ep
		p := d.NewProvider(sv, 10*time.Second, func() govern.Endpoint { return x })
		defer p.Close()
	}

	if err := Wait(&r, len(endpoints)); err != nil {
		t.Fatalf("%s: wait: %v", d.Name(), err)
	}

	if got, want := r.Endpoints, endpoints; !reflect.DeepEqual(got, want) {
		t.Fatalf("%s: endpoints: got %v, want %v", d.Name(), got, want)
	} else {
		t.Logf("%s: endpoints: got %v", d.Name(), got)
	}
}
