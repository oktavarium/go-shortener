package handlers

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oktavarium/go-shortener/internal/shortener/internal/server/internal/storage"
)

type Handlers struct {
	storage storage.Storage
}

func NewHandlers(s storage.Storage) Handlers {
	return Handlers{
		storage: s,
	}
}

func (h *Handlers) CreateUrl(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	//server will close this body
	// we close bopdy manually only in response after client.Do
	//defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortName := base64.StdEncoding.EncodeToString(body)
	h.storage.Save(shortName, string(body))
	shortName = "http://localhost:8080/" + base64.StdEncoding.EncodeToString(body)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortName))
	w.Header().Set("Context-Type", "text/plain")
}

func (h *Handlers) GetUrl(w http.ResponseWriter, r *http.Request) {
	ctx := chi.RouteContext(r.Context())
	id := ctx.URLParam("id")
	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	v, ok := h.storage.Get(id)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", v)
	w.WriteHeader(http.StatusTemporaryRedirect)

}
