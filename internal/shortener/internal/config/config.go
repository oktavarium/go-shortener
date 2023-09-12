package config

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Addr     string `env:"SERVER_ADDRESS"`
	BaseAddr string `env:"BASE_URL"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	flag.StringVar(&cfg.Addr, "a", "localhost:8080", "http-server address and port")
	flag.StringVar(&cfg.BaseAddr, "b", "http://localhost:8080/", "base redirect address")
	flag.Parse()

	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("error on reading env parameters: %w", err)
	}

	if cfg.BaseAddr[len(cfg.BaseAddr)-1] != '/' {
		cfg.BaseAddr += "/"
	}

	if len(flag.Args()) > 0 {
		return cfg, fmt.Errorf("too many flags in bash")
	}

	return cfg, nil
}
