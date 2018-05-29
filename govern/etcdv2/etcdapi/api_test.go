package etcdapi

import (
	"context"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/ironzhang/x-pearls/govern/testutil"
	"github.com/ironzhang/x-pearls/govern/testutil/etcdv2util"
)

func TestParseName(t *testing.T) {
	tests := []struct {
		key  string
		name string
	}{
		{
			key:  "/",
			name: "",
		},
		{
			key:  "/foo",
			name: "foo",
		},
		{
			key:  "/foo/bar",
			name: "bar",
		},
		{
			key:  "/foo/",
			name: "",
		},
	}
	for _, tt := range tests {
		name, err := parseName(tt.key)
		if err != nil {
			t.Fatalf("%q: parse name: %v", tt.key, err)
		}
		if got, want := name, tt.name; got != want {
			t.Errorf("%q: parse name: got %q, want %q", tt.key, got, want)
		} else {
			t.Logf("%q parse name: got %q", tt.key, got)
		}
	}
}

func TestAPISet(t *testing.T) {
	api := NewAPI(etcdv2util.NewTestKeysAPI(), &testutil.Endpoint{})

	endpoints := []testutil.Endpoint{
		{Name: "node1"},
		{Name: "node2"},
		{Name: "node3"},
	}
	for _, ep := range endpoints {
		if err := api.Set(context.Background(), "/TestAPISet", &ep, 10*time.Second); err != nil {
			t.Fatalf("set: %v", err)
		}
	}

	eps, index, err := api.Get(context.Background(), "/TestAPISet")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	sort.Slice(eps, func(i, j int) bool { return eps[i].Node() < eps[j].Node() })

	for i, ep := range eps {
		if got, want := *ep.(*testutil.Endpoint), endpoints[i]; got != want {
			t.Fatalf("%d: endpoint: got %v, want %v", i, got, want)
		}
	}
	t.Logf("index: %v, endpoints: %v", index, eps)
}

func TestAPIDel(t *testing.T) {
	api := NewAPI(etcdv2util.NewTestKeysAPI(), &testutil.Endpoint{})

	endpoints := []testutil.Endpoint{
		{Name: "node1"},
		{Name: "node2"},
		{Name: "node3"},
	}
	for _, ep := range endpoints {
		if err := api.Set(context.Background(), "/TestAPIDel", &ep, 10*time.Second); err != nil {
			t.Fatalf("set: %v", err)
		}
	}

	eps, _, err := api.Get(context.Background(), "/TestAPIDel")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got, want := len(eps), len(endpoints); got != want {
		t.Fatalf("count: got %v, want %v", got, want)
	}

	for _, ep := range eps {
		if err := api.Del(context.Background(), "/TestAPIDel", ep.Node()); err != nil {
			t.Fatalf("del: %v", err)
		}
	}

	eps, _, err = api.Get(context.Background(), "/TestAPIDel")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got, want := len(eps), 0; got != want {
		t.Fatalf("count: got %v, want %v", got, want)
	}
}

func TestWatcher(t *testing.T) {
	api := NewAPI(etcdv2util.NewTestKeysAPI(), &testutil.Endpoint{})

	events := []Event{
		{Action: "set", Name: "node1", Endpoint: &testutil.Endpoint{Name: "node1"}},
		{Action: "set", Name: "node2", Endpoint: &testutil.Endpoint{Name: "node2"}},
		{Action: "set", Name: "node1", Endpoint: &testutil.Endpoint{Name: "node1"}},
		{Action: "delete", Name: "node1", Endpoint: nil},
		{Action: "delete", Name: "node2", Endpoint: nil},
		{Action: "set", Name: "node2", Endpoint: &testutil.Endpoint{Name: "node2"}},
	}

	go func() {
		for _, event := range events {
			switch event.Action {
			case "set":
				api.Set(context.Background(), "/TestWatcher", event.Endpoint, 10*time.Second)
			case "delete":
				api.Del(context.Background(), "/TestWatcher", event.Name)
			}
		}
	}()

	w := api.Watcher("/TestWatcher", 0)
	for i, want := range events {
		got, err := w.Next(context.Background())
		if err != nil {
			t.Fatalf("%d: next: %v", i, err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("%d: event: got %v, want %v", i, got, want)
		}
		t.Logf("%d: event: %v", i, got)
	}
}
