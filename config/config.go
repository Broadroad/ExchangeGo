package config

import (
	"encoding/json"
	"io/ioutil"
)

type HuobiConfig struct {
	Apikey    string
	Secret    string
	Frequence float64
}

type FCoinConfig struct {
	Apikey    string
	Secret    string
	Frequence float64
}

type Config struct {
	Huobi *HuobiConfig
	FCoin *FCoinConfig
}

func NewConfig(path string) *Config {
	var c Config

	var b []byte
	var err error
	if b, err = ioutil.ReadFile(path); err != nil {
		panic(err)
	}

	if err = json.Unmarshal(b, &c); err != nil {
		panic(err)
	}

	return &c
}
