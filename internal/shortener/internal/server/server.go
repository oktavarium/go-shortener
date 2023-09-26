package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oktavarium/go-shortener/internal/shortener/internal/server/internal/handlers"
	"github.com/oktavarium/go-shortener/internal/shortener/internal/server/internal/logger"
	"github.com/oktavarium/go-shortener/internal/shortener/internal/server/internal/storage"
)

type Server struct {
	router *chi.Mux
	addr   string
}

func NewServer(addr, baseAddr string, logLevel string) (Server, error) {
	server := Server{
		router: chi.NewRouter(),
		addr:   addr,
	}

	storage := storage.NewStorage()
	handlers := handlers.NewHandlers(storage, baseAddr)
	err := logger.InitLogger(logLevel)
	if err != nil {
		return server, fmt.Errorf("failed on logger init: %w", err)
	}

	server.router.Use(logger.LogMiddleware)
	server.router.Post("/", handlers.CreateURL)
	server.router.Get("/{id}", handlers.GetURL)

	return server, err
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(s.addr, s.router)
}
