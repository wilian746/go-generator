package product

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/wilian746/go-generator/pkg/standart-gorm/internal/controllers/product"
	"github.com/wilian746/go-generator/pkg/standart-gorm/internal/handlers"
	RulesProduct "github.com/wilian746/go-generator/pkg/standart-gorm/internal/rules/product"
	HttpStatus "github.com/wilian746/go-generator/pkg/standart-gorm/internal/utils/http"
	"github.com/wilian746/go-generator/pkg/standart-gorm/pkg/repository/adapter"
	"net/http"
)

type Handler struct {
	handlers.Interface

	Controller product.Interface
	Rules      *RulesProduct.Rules
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Controller: product.NewController(repository),
		Rules:      RulesProduct.NewRules(),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "ID") != "" {
		h.getOne(w, r)
	} else {
		h.getAll(w, r)
	}
}

func (h *Handler) getOne(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}

	response, err := h.Controller.ListOne(ID)
	if err != nil {
		if err.Error() == adapter.ErrRecordNotFound.Error() {
			HttpStatus.StatusNotfound(w, r, err)
			return
		}
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, response)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	response, err := h.Controller.ListAll()
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, response)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	productBody, err := h.Rules.ConvertIoReaderToProduct(r.Body)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("body is required"))
		return
	}
	ID, err := h.Controller.Create(productBody)
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, errors.New("error when create"))
		return
	}

	HttpStatus.StatusOK(w, r, map[string]interface{}{"id": ID.String()})
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}

	productBody, err := h.Rules.ConvertIoReaderToProduct(r.Body)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}

	if err := h.Controller.Update(ID, productBody); err != nil {
		if err.Error() == adapter.ErrRecordNotFound.Error() {
			HttpStatus.StatusNotfound(w, r, err)
			return
		}
		HttpStatus.StatusInternalServerError(w, r, errors.New("error when create"))
		return
	}

	HttpStatus.StatusNoContent(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}

	if err := h.Controller.Remove(ID); err != nil {
		if err.Error() == adapter.ErrRecordNotFound.Error() {
			HttpStatus.StatusNotfound(w, r, err)
			return
		}
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusNoContent(w, r)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}
