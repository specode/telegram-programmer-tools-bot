package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var cfg *Config

type Config struct {
	Bot struct {
		Token string
		Debug bool
	}
}

func initConfig(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	cfg = &Config{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return err
	}
	return nil
}
