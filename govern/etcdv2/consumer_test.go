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

func TestConsumer(t *testing.T) {
	defer time.Sleep(500 * time.Millisecond)

	r := &testutil.Refresher{}
	api := etcdv2util.NewTestKeysAPI()

	c1 := newConsumer(api, "/TestConsumerGetEndpoints", &testutil.Endpoint{}, nil)
	defer c1.Close()

	c2 := newConsumer(api, "/TestConsumerRefresh", &testutil.Endpoint{}, r.Refresh)
	defer c2.Close()

	register := func(dir string, ep govern.Endpoint) error {
		eapi := etcdapi.NewAPI(api, &testutil.Endpoint{})
		return eapi.Set(context.Background(), dir, ep, 10*time.Second)
	}
	testutil.TestConsumer(t, c1, register)
	testutil.TestConsumerRefresh(t, c2, r, register)
}
