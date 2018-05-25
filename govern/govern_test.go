package govern_test

import (
	"reflect"
	"testing"

	"github.com/ironzhang/x-pearls/govern"
)

type TestEndpoint struct {
	Name string
}

func (p *TestEndpoint) Node() string {
	return p.Name
}

func (p *TestEndpoint) String() string {
	return p.Name
}

func (p *TestEndpoint) Equal(a interface{}) bool {
	return *p == *a.(*TestEndpoint)
}

func TestEndpointsAdd(t *testing.T) {
	m := govern.Endpoints{
		"node0": &TestEndpoint{"node0"},
		"node1": &TestEndpoint{"node1"},
		"node2": &TestEndpoint{"node2"},
		"node3": &TestEndpoint{"node3"},
	}

	tests := []struct {
		ep   *TestEndpoint
		want bool
	}{
		{ep: &TestEndpoint{"endpoint0"}, want: true},
		{ep: &TestEndpoint{"endpoint1"}, want: true},
		{ep: &TestEndpoint{"endpoint1"}, want: false},
		{ep: &TestEndpoint{"node0"}, want: false},
	}
	for i, tt := range tests {
		if got, want := m.Add(tt.ep), tt.want; got != want {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}

	result := govern.Endpoints{
		"node0":     &TestEndpoint{"node0"},
		"node1":     &TestEndpoint{"node1"},
		"node2":     &TestEndpoint{"node2"},
		"node3":     &TestEndpoint{"node3"},
		"endpoint0": &TestEndpoint{"endpoint0"},
		"endpoint1": &TestEndpoint{"endpoint1"},
	}
	if got, want := m.SortList(), result.SortList(); !reflect.DeepEqual(got, want) {
		t.Errorf("result: got %v, want %v", got, want)
	} else {
		t.Logf("result: got %v", got)
	}
}
