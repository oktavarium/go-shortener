package shortener

import (
	"github.com/oktavarium/go-shortener/internal/shortener/internal/server"
)

func Run() error {
	s := server.NewServer("localhost:8080")
	return s.ListenAndServe()
}
