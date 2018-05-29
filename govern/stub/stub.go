package memory

import (
	"sync"

	"github.com/ironzhang/x-pearls/govern"
)

type stub struct {
	mu          sync.Mutex
	list        []govern.Endpoint
	endpoints   govern.Endpoints
	subscribers map[string]govern.RefreshEndpointsFunc
}

func newStub() *stub {
	return &stub{
		endpoints:   make(govern.Endpoints),
		subscribers: make(map[string]govern.RefreshEndpointsFunc),
	}
}

func (p *stub) AddEndpoint(ep govern.Endpoint) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.endpoints.Add(ep) {
		p.doRefresh(p.endpoints.SortList())
	}
}

func (p *stub) RemoveEndpoint(node string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.endpoints.Remove(node) {
		p.doRefresh(p.endpoints.SortList())
	}
}

func (p *stub) AddSubscriber(token string, f govern.RefreshEndpointsFunc) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.subscribers[token] = f
}

func (p *stub) RemoveSubscriber(token string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.subscribers, token)
}

func (p *stub) GetEndpoints() []govern.Endpoint {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.list
}

func (p *stub) doRefresh(eps []govern.Endpoint) {
	p.list = eps
	for _, f := range p.subscribers {
		f(eps)
	}
}
