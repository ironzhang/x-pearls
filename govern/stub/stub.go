package stub

import (
	"sync"

	"github.com/ironzhang/x-pearls/govern"
)

type stub struct {
	endpoints govern.Endpoints
	emu       sync.RWMutex
	eps       []govern.Endpoint
	smu       sync.Mutex
	subs      map[string]govern.RefreshEndpointsFunc
}

func (p *stub) AddEndpoint(ep govern.Endpoint) {
	if p.endpoints.Add(ep) {
		p.doRefresh(p.endpoints.SortList())
	}
}

func (p *stub) RemoveEndpoint(node string) {
	if p.endpoints.Remove(node) {
		p.doRefresh(p.endpoints.SortList())
	}
}

func (p *stub) AddSubscriber(token string, f govern.RefreshEndpointsFunc) {
	p.smu.Lock()
	p.subs[token] = f
	p.smu.Unlock()
}

func (p *stub) RemoveSubscriber(token string) {
	p.smu.Lock()
	delete(p.subs, token)
	p.smu.Unlock()
}

func (p *stub) GetEndpoints() []govern.Endpoint {
	p.emu.RLock()
	defer p.emu.RUnlock()
	return p.eps
}

func (p *stub) doRefresh(eps []govern.Endpoint) {
	p.emu.Lock()
	p.eps = eps
	p.emu.Unlock()

	p.smu.Lock()
	subs := make([]govern.RefreshEndpointsFunc, len(p.subs))
	for _, s := range p.subs {
		subs = append(subs, s)
	}
	p.smu.Unlock()

	for _, f := range subs {
		f(eps)
	}
}
