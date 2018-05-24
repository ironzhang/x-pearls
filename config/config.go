package config

import (
	"encoding/json"
	"io/ioutil"
)

type LoadWriter interface {
	LoadFromFile(filename string, cfg interface{}) error
	WriteToFile(filename string, cfg interface{}) error
}

type jsoncfg struct{}

func (jsoncfg) LoadFromFile(filename string, cfg interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, cfg); err != nil {
		return err
	}
	return nil
}

func (jsoncfg) WriteToFile(filename string, cfg interface{}) error {
	data, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(filename, data, 0666); err != nil {
		return err
	}
	return nil
}

var (
	JSON    jsoncfg
	Default LoadWriter = JSON
)

func LoadFromFile(filename string, cfg interface{}) error {
	return Default.LoadFromFile(filename, cfg)
}

func WriteToFile(filename string, cfg interface{}) error {
	return Default.WriteToFile(filename, cfg)
}
