package config

import "flag"

type Config struct {
	Addr     string
	BaseAddr string
}

func LoadConfig() Config {
	var cfg Config
	flag.StringVar(&cfg.Addr, "a", "localhost:8080", "http-server address and port")
	flag.StringVar(&cfg.BaseAddr, "b", "http://localhost:8080/", "base redirect address")
	flag.Parse()
	if cfg.BaseAddr[len(cfg.BaseAddr)-1] != '/' {
		cfg.BaseAddr += "/"
	}
	return cfg
}
