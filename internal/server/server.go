package server

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var storage = map[string]string{}

func Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		shortName := base64.StdEncoding.EncodeToString(body)
		storage[shortName] = string(body)
		shortName = "http://localhost:8080/" + base64.StdEncoding.EncodeToString(body)
		w.Write([]byte(shortName))
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Context-Type", "text/plain")
		fmt.Println(storage)
	})

	router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		ctx := chi.RouteContext(r.Context())
		id := ctx.URLParam("id")
		if len(id) == 0 {
			fmt.Println("1")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		v, ok := storage[id]
		if !ok {
			fmt.Println("2")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Location", v)
		w.WriteHeader(http.StatusTemporaryRedirect)

	})

	return http.ListenAndServe(":8080", router)
}
