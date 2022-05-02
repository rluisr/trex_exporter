package main

import (
	"github.com/Netflix/go-env"
)

type Config struct {
	ListenPort int `env:"LISTEN_PORT,default=49152"`
}

func NewConfig() (*Config, error) {
	var config Config

	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
