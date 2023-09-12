package shortener

import (
	"github.com/oktavarium/go-shortener/internal/shortener/internal/config"
	"github.com/oktavarium/go-shortener/internal/shortener/internal/server"
)

func Run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	s := server.NewServer(cfg.Addr, cfg.BaseAddr)
	return s.ListenAndServe()
}
