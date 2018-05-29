package stub

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/ironzhang/x-pearls/govern"
)

const DriverName = "stub"

type provider struct {
	stub *stub
	dir  string
	done chan<- struct{}
}

func newProvider(stub *stub, dir string, interval time.Duration, f govern.GetEndpointFunc) *provider {
	if f == nil {
		panic("govern.GetEndpointFunc is nil")
	}

	ch := make(chan struct{})
	go func(done <-chan struct{}) {
		t := time.NewTicker(interval)
		defer t.Stop()

		stub.AddEndpoint(f())
		for {
			select {
			case <-t.C:
				stub.AddEndpoint(f())
			case <-done:
				stub.RemoveEndpoint(f().Node())
				return
			}
		}
	}(ch)

	return &provider{stub: stub, dir: dir, done: ch}
}

func (p *provider) Driver() string {
	return DriverName
}

func (p *provider) Directory() string {
	return p.dir
}

func (p *provider) Close() error {
	close(p.done)
	return nil
}

type consumer struct {
	stub  *stub
	dir   string
	token string
}

func newConsumer(stub *stub, dir string, refresh govern.RefreshEndpointsFunc) *consumer {
	token := ""
	if refresh != nil {
		token = fmt.Sprint(rand.Int())
		stub.AddSubscriber(token, refresh)
	}
	return &consumer{
		stub:  stub,
		dir:   dir,
		token: token,
	}
}

func (p *consumer) Driver() string {
	return DriverName
}

func (p *consumer) Directory() string {
	return p.dir
}

func (p *consumer) Close() error {
	if p.token != "" {
		p.stub.RemoveSubscriber(p.token)
	}
	return nil
}

func (p *consumer) GetEndpoints() []govern.Endpoint {
	return p.stub.GetEndpoints()
}

type Driver struct {
	namespace string
	mu        sync.Mutex
	stubs     map[string]*stub
}

func NewDriver(namespace string) *Driver {
	return &Driver{namespace: namespace, stubs: make(map[string]*stub)}
}

func (p *Driver) Name() string {
	return DriverName
}

func (p *Driver) Namespace() string {
	return p.namespace
}

func (p *Driver) NewProvider(service string, interval time.Duration, f govern.GetEndpointFunc) govern.Provider {
	p.mu.Lock()
	stub, ok := p.stubs[service]
	if !ok {
		stub = newStub()
		p.stubs[service] = stub
	}
	p.mu.Unlock()
	return newProvider(stub, p.dir(service), interval, f)
}

func (p *Driver) NewConsumer(service string, endpoint govern.Endpoint, f govern.RefreshEndpointsFunc) govern.Consumer {
	p.mu.Lock()
	stub, ok := p.stubs[service]
	if !ok {
		stub = newStub()
		p.stubs[service] = stub
	}
	p.mu.Unlock()
	return newConsumer(stub, p.dir(service), f)
}

func (p *Driver) Close() error {
	return nil
}

func (p *Driver) dir(service string) string {
	return fmt.Sprintf("/%s/%s", p.namespace, service)
}

type Config struct {
	Endpoints []govern.Endpoint
}

func Open(namespace string, config interface{}) (govern.Driver, error) {
	return NewDriver(namespace), nil
}

func init() {
	govern.Register(DriverName, Open)
}
