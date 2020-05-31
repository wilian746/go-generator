package health

import (
	"errors"
	_ "github.com/wilian746/go-generator/pkg/standart-gorm/internal/entities/health" // import used in swagger
	"github.com/wilian746/go-generator/pkg/standart-gorm/internal/handlers"
	HttpStatus "github.com/wilian746/go-generator/pkg/standart-gorm/internal/utils/http"
	"github.com/wilian746/go-generator/pkg/standart-gorm/pkg/repository/adapter"
	"net/http"
)

var ErrRelationalNotOK = errors.New("{ERROR_HEALTH} Relational database not alive")

type Handler struct {
	handlers.Interface
	Repository adapter.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Repository: repository,
	}
}

// @Tags Health
// @Security ApiKeyAuth
// @Description Check if Health  of service it's OK!
// @ID health
// @Accept  json
// @Produce  json
// @Success 200 {object} health.ResponseHealth
// @Failure 500 {object} http.ResponseError
// @Router /health [get]
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if !h.Repository.Health() {
		HttpStatus.StatusInternalServerError(w, r, ErrRelationalNotOK)
		return
	}

	HttpStatus.StatusOK(w, r, "Service OK")
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}
