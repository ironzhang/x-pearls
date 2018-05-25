package memory

type Endpoint struct {
	Name string
}

func (p *Endpoint) Node() string {
	return p.Name
}

func (p *Endpoint) String() string {
	return p.Name
}

func (p *Endpoint) Equal(a interface{}) bool {
	return *p == *a.(*Endpoint)
}

/*
func TestProvider(t *testing.T) {
	var count int
	f := func() govern.Endpoint {
		count++
		return nil
	}

	p := newProvider("/TestProvider", 100*time.Millisecond, f)
	time.Sleep(200*time.Millisecond + 10*time.Millisecond)
	p.Close()

	if got, want := count, 3; got != want {
		t.Errorf("got %v, want %v", got, want)
	} else {
		t.Logf("got %v", got)
	}
}

func TestConsumer(t *testing.T) {
	var endpoints []govern.Endpoint
	f := func(eps []govern.Endpoint) {
		endpoints = eps
	}

	tests := []struct {
		endpoints []govern.Endpoint
	}{
		{endpoints: nil},
		{endpoints: []govern.Endpoint{
			&Endpoint{Name: "n0"},
			&Endpoint{Name: "n1"},
		}},
	}
	for i, tt := range tests {
		c := NewConsumer("/TestConsumer", tt.endpoints, f)
		if got, want := endpoints, tt.endpoints; !reflect.DeepEqual(got, want) {
			t.Fatalf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
		if got, want := c.GetEndpoints(), tt.endpoints; !reflect.DeepEqual(got, want) {
			t.Fatalf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}

func TestDriverNewProvider(t *testing.T) {
	var count int
	f := func() govern.Endpoint {
		count++
		return &Endpoint{Name: "node0"}
	}

	d := NewDriver("test", nil)
	p := d.NewProvider("service", 100*time.Millisecond, f)
	time.Sleep(200*time.Millisecond + 10*time.Millisecond)
	p.Close()

	if got, want := count, 3; got != want {
		t.Errorf("got %v, want %v", got, want)
	} else {
		t.Logf("got %v", got)
	}
}

func TestDriverNewConsumer(t *testing.T) {
	var endpoints []govern.Endpoint
	f := func(eps []govern.Endpoint) {
		endpoints = eps
	}

	tests := []struct {
		endpoints []govern.Endpoint
	}{
		{endpoints: nil},
		{endpoints: []govern.Endpoint{
			&Endpoint{Name: "n0"},
			&Endpoint{Name: "n1"},
		}},
	}
	for i, tt := range tests {
		d := NewDriver("test", tt.endpoints)
		c := d.NewConsumer("service", nil, f)
		if got, want := endpoints, tt.endpoints; !reflect.DeepEqual(got, want) {
			t.Fatalf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
		if got, want := c.GetEndpoints(), tt.endpoints; !reflect.DeepEqual(got, want) {
			t.Fatalf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}

func TestDriverOpen(t *testing.T) {
	d, err := Open("test", Config{})
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer d.Close()
}
*/
