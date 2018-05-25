package stub

const DriverName = "stub"

/*
type provider struct {
	dir  string
	done chan struct{}
}

func newProvider(dir string, interval time.Duration, f govern.GetEndpointFunc) *provider {
	if f == nil {
		panic("govern.GetEndpointFunc is nil")
	}

	ch := make(chan struct{})
	go func(done <-chan struct{}) {
		t := time.NewTicker(interval)
		defer t.Stop()

		f()
		for {
			select {
			case <-t.C:
				f()
			case <-done:
				return
			}
		}
	}(ch)

	return &provider{dir: dir, done: ch}
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

type Consumer struct {
	dir       string
	endpoints []govern.Endpoint
}

func NewConsumer(dir string, endpoints []govern.Endpoint, refresh govern.RefreshEndpointsFunc) *Consumer {
	if refresh != nil {
		refresh(endpoints)
	}
	return &Consumer{
		dir:       dir,
		endpoints: endpoints,
	}
}

func (p *Consumer) Driver() string {
	return DriverName
}

func (p *Consumer) Directory() string {
	return p.dir
}

func (p *Consumer) Close() error {
	return nil
}

func (p *Consumer) GetEndpoints() []govern.Endpoint {
	return p.endpoints
}

type Driver struct {
	namespace string
	endpoints []govern.Endpoint
}

func NewDriver(namespace string, endpoints []govern.Endpoint) *Driver {
	return &Driver{namespace: namespace, endpoints: endpoints}
}

func (p *Driver) Name() string {
	return DriverName
}

func (p *Driver) Namespace() string {
	return p.namespace
}

func (p *Driver) NewProvider(service string, interval time.Duration, f govern.GetEndpointFunc) govern.Provider {
	return newProvider(p.dir(service), interval, f)
}

func (p *Driver) NewConsumer(service string, endpoint govern.Endpoint, f govern.RefreshEndpointsFunc) govern.Consumer {
	return NewConsumer(p.dir(service), p.endpoints, f)
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
	c := config.(Config)
	return NewDriver(namespace, c.Endpoints), nil
}

func init() {
	govern.Register(DriverName, Open)
}
*/
