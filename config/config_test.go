package config_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/ironzhang/x-pearls/config"
)

type Service struct {
	Name string
	Host string
}

type DB struct {
	Hostname string
	Username string
	Password string
}

type Sub struct {
	I8  int8
	I16 int16
	U8  uint8
	U16 uint16
}

type Misc struct {
	I32 int32
	I64 int64
	U32 uint32
	U64 uint64
	F32 float32
	F64 float64
	Sub Sub
}

type Config struct {
	Environment string
	Service     Service
	DB          DB
	Misc        Misc
}

var example = Config{
	Environment: "test",
	Service: Service{
		Name: "config",
		Host: "127.0.0.1:2000",
	},
	DB: DB{
		Hostname: "localhost:3306",
		Username: "root",
		Password: "123456",
	},
	Misc: Misc{
		I32: 32,
		I64: 64,
		U32: 32,
		U64: 64,
		F32: 3.14,
		F64: 3.1415926,
		Sub: Sub{
			I8:  8,
			I16: 16,
			U8:  8,
			U16: 16,
		},
	},
}

func TestJSONConfig(t *testing.T) {
	filename := "test.json"
	got, want := Config{}, example
	if err := config.JSON.WriteToFile(filename, want); err != nil {
		t.Fatalf("write to file: %v", err)
	}
	defer os.Remove(filename)
	if err := config.JSON.LoadFromFile(filename, &got); err != nil {
		t.Fatalf("load from file: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%+v != %+v", got, want)
	}
	//fmt.Printf("%+v\n", got)
}

func TestTOMLConfig(t *testing.T) {
	filename := "test.cfg"
	got, want := Config{}, example
	if err := config.TOML.WriteToFile(filename, want); err != nil {
		t.Fatalf("write to file: %v", err)
	}
	defer os.Remove(filename)
	if err := config.TOML.LoadFromFile(filename, &got); err != nil {
		t.Fatalf("load from file: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%+v != %+v", got, want)
	}
	//fmt.Printf("%+v\n", got)
}

func TestConfig(t *testing.T) {
	tests := []struct {
		filename  string
		defaultLW config.LoadWriter
	}{
		{"test.config.json", config.JSON},
		{"test.config.conf", config.TOML},
	}
	for _, tt := range tests {
		got, want := Config{}, example
		config.Default = tt.defaultLW
		if err := config.WriteToFile(tt.filename, want); err != nil {
			t.Fatalf("%s: write to file: %v", tt.filename, err)
		}
		defer os.Remove(tt.filename)
		if err := config.LoadFromFile(tt.filename, &got); err != nil {
			t.Fatalf("%s: load from file: %v", tt.filename, err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%+v != %+v", got, want)
		}
		//fmt.Printf("%+v\n", got)
	}
}
