package testutil

import "github.com/ironzhang/x-pearls/govern"

type Refresher struct {
	Count     int
	Endpoints []govern.Endpoint
}

func (t *Refresher) Refresh(endpoints []govern.Endpoint) {
	t.Count++
	t.Endpoints = endpoints
}
