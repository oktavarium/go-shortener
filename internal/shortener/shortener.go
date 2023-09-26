package shortener

import (
	"fmt"

	"github.com/oktavarium/go-shortener/internal/shortener/internal/config"
	"github.com/oktavarium/go-shortener/internal/shortener/internal/server"
)

func Run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed on loading config: %w", err)
	}
	s, err := server.NewServer(cfg.Addr, cfg.BaseAddr, cfg.LogLevel)
	if err != nil {
		return fmt.Errorf("failed creating new server: %w", err)
	}
	return s.ListenAndServe()
}
