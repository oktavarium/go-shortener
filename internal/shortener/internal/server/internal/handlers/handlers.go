package handlers

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oktavarium/go-shortener/internal/shortener/internal/server/internal/models"
	"github.com/oktavarium/go-shortener/internal/shortener/internal/server/internal/storage"
)

type Handlers struct {
	storage  storage.Storage
	baseAddr string
}

func NewHandlers(s storage.Storage, baseAddr string) Handlers {
	return Handlers{
		storage:  s,
		baseAddr: baseAddr,
	}
}

func (h *Handlers) CreateURL(w http.ResponseWriter, r *http.Request) {
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
	shortName += h.baseAddr
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortName))
	w.Header().Set("Context-Type", "text/plain")
}

func (h *Handlers) GetURL(w http.ResponseWriter, r *http.Request) {
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

func (h *Handlers) GetJSONURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var incomingData models.IncomingData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&incomingData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var outcomingData models.OutcomingData

	shortName := base64.StdEncoding.EncodeToString([]byte(incomingData.URL))
	h.storage.Save(shortName, incomingData.URL)
	shortName += h.baseAddr

	outcomingData.Result = shortName
	encoder := json.NewEncoder(w)
	err = encoder.Encode(&outcomingData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
