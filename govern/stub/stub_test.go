package memory

import (
	"reflect"
	"testing"

	"github.com/ironzhang/x-pearls/govern"
	"github.com/ironzhang/x-pearls/govern/testutil"
)

func TestStub(t *testing.T) {
	var r1, r2 testutil.Refresher

	s := newStub()
	s.AddSubscriber("1", r1.Refresh)
	s.AddSubscriber("2", r2.Refresh)

	endpoints := []govern.Endpoint{
		&testutil.Endpoint{Name: "node1", Addr: "1.0"},
		&testutil.Endpoint{Name: "node0", Addr: "0.0"},
		&testutil.Endpoint{Name: "node0", Addr: "0.1"},
	}
	result := []govern.Endpoint{
		&testutil.Endpoint{Name: "node0", Addr: "0.1"},
		&testutil.Endpoint{Name: "node1", Addr: "1.0"},
	}

	for _, ep := range endpoints {
		s.AddEndpoint(ep)
	}

	if got, want := r1.Count, 3; got != want {
		t.Errorf("r1.Count: got %v, want %v", got, want)
	}
	if got, want := r1.Endpoints, result; !reflect.DeepEqual(got, want) {
		t.Errorf("r1.Endpoints: got %v, want %v", got, want)
	}
	if got, want := r2.Count, 3; got != want {
		t.Errorf("r2.Count: got %v, want %v", got, want)
	}
	if got, want := r2.Endpoints, result; !reflect.DeepEqual(got, want) {
		t.Errorf("r2.Endpoints: got %v, want %v", got, want)
	}
}
