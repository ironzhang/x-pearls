package testutil

import (
	"encoding/json"
	"fmt"
)

type Endpoint struct {
	Name string
	Addr string
}

func (p *Endpoint) Node() string {
	return p.Name
}

func (p *Endpoint) String() string {
	return fmt.Sprintf("Name: %s, Addr: %s", p.Name, p.Addr)
}

func (p *Endpoint) Equal(a interface{}) bool {
	ep := a.(*Endpoint)
	return *p == *ep
}

func (p *Endpoint) Marshal() (string, error) {
	data, err := json.Marshal(p)
	return string(data), err
}

func (p *Endpoint) Unmarshal(s string) error {
	return json.Unmarshal([]byte(s), p)
}
