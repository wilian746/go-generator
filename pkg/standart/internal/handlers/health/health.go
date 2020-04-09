package health

import (
	"encoding/json"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/adapter"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/handlers"
	"net/http"
)

type Handler struct {
	handlers.Interface
	Repository adapter.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Repository: repository,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if !h.Repository.Health() {
		w.WriteHeader(http.StatusInternalServerError)
		bytes, _ := json.Marshal("Relational database not alive")
		_, _ = w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal("Service OK")
	_, _ = w.Write(bytes)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *Handler) Options(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
