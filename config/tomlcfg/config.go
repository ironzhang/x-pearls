package tomlcfg

import (
	"os"

	"github.com/BurntSushi/toml"
)

type config struct{}

func (config) LoadFromFile(filename string, cfg interface{}) error {
	_, err := toml.DecodeFile(filename, cfg)
	return err
}

func (config) WriteToFile(filename string, cfg interface{}) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := toml.NewEncoder(f)
	enc.Indent = "\t"
	return enc.Encode(cfg)
}

var TOML config
