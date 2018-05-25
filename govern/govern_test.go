package govern_test

import (
	"reflect"
	"testing"

	"github.com/ironzhang/x-pearls/govern"
	"github.com/ironzhang/x-pearls/govern/testutil"
)

func TestEndpointsAdd(t *testing.T) {
	m := govern.Endpoints{
		"node0": &testutil.Endpoint{Name: "node0"},
		"node1": &testutil.Endpoint{Name: "node1"},
		"node2": &testutil.Endpoint{Name: "node2"},
		"node3": &testutil.Endpoint{Name: "node3"},
	}

	tests := []struct {
		ep   *testutil.Endpoint
		want bool
	}{
		{ep: &testutil.Endpoint{Name: "endpoint0"}, want: true},
		{ep: &testutil.Endpoint{Name: "endpoint1"}, want: true},
		{ep: &testutil.Endpoint{Name: "endpoint1"}, want: false},
		{ep: &testutil.Endpoint{Name: "node0"}, want: false},
	}
	for i, tt := range tests {
		if got, want := m.Add(tt.ep), tt.want; got != want {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}

	result := govern.Endpoints{
		"node0":     &testutil.Endpoint{Name: "node0"},
		"node1":     &testutil.Endpoint{Name: "node1"},
		"node2":     &testutil.Endpoint{Name: "node2"},
		"node3":     &testutil.Endpoint{Name: "node3"},
		"endpoint0": &testutil.Endpoint{Name: "endpoint0"},
		"endpoint1": &testutil.Endpoint{Name: "endpoint1"},
	}
	if got, want := m.SortList(), result.SortList(); !reflect.DeepEqual(got, want) {
		t.Errorf("result: got %v, want %v", got, want)
	} else {
		t.Logf("result: got %v", got)
	}
}
