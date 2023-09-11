package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oktavarium/go-shortener/internal/shortener/internal/server/internal/handlers"
	"github.com/oktavarium/go-shortener/internal/shortener/internal/server/internal/storage"
)

type Server struct {
	router *chi.Mux
	addr   string
}

func NewServer(addr string) Server {
	server := Server{
		router: chi.NewRouter(),
		addr:   addr,
	}

	storage := storage.NewStorage()
	handlers := handlers.NewHandlers(storage)

	server.router.Use(middleware.Logger)
	server.router.Post("/", handlers.CreateUrl)
	server.router.Get("/{id}", handlers.GetUrl)

	return server
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(s.addr, s.router)
}
