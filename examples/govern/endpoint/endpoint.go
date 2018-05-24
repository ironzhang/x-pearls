package endpoint

import "encoding/json"

type Endpoint struct {
	Name  string
	Addrs []string
	Load  float64
}

func (p *Endpoint) Node() string {
	return p.Name
}

func (p *Endpoint) String() string {
	data, _ := json.Marshal(p)
	return string(data)
}

func (p *Endpoint) Equal(a interface{}) bool {
	ep := a.(*Endpoint)
	if p.Name != ep.Name {
		return false
	}
	if len(p.Addrs) != len(ep.Addrs) {
		return false
	}
	for i := range p.Addrs {
		if p.Addrs[i] != ep.Addrs[i] {
			return false
		}
	}
	if p.Load != ep.Load {
		return false
	}
	return true
}
