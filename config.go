package main

import (
	"github.com/Netflix/go-env"
	"time"
)

type Config struct {
	TrexAPIAddress  string        `env:"TREX_API_ADDRESS,required=true"`
	TrexWorkerName  string        `env:"TREX_WORKER_NAME,required=true"`
	ListenPort      int           `env:"LISTEN_PORT,default=49152"`
	CollectInterval time.Duration `env:"COLLECT_INTERVAL,default=60s"`
}

func NewConfig() (*Config, error) {
	var config Config

	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
