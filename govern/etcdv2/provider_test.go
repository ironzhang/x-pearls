package etcdv2

import (
	"context"
	"testing"
	"time"

	"github.com/ironzhang/x-pearls/govern"
	"github.com/ironzhang/x-pearls/govern/etcdv2/etcdapi"
	"github.com/ironzhang/x-pearls/govern/testutil"
	"github.com/ironzhang/x-pearls/govern/testutil/etcdv2util"
)

func TestProvider(t *testing.T) {
	defer time.Sleep(500 * time.Millisecond)

	ep := &testutil.Endpoint{Name: "node0", Addr: "localhost:2000"}
	api := etcdv2util.NewTestKeysAPI()

	p := newProvider(api, "/TestProvider", 10*time.Second, func() govern.Endpoint { return ep })
	defer p.Close()

	list := func(dir string) ([]govern.Endpoint, error) {
		eapi := etcdapi.NewAPI(api, &testutil.Endpoint{})
		eps, _, err := eapi.Get(context.Background(), dir)
		return eps, err
	}
	testutil.TestProvider(t, p, ep, list)
}
