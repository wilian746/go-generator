package product

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/adapter"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/controllers/product"
	EntityProduct "github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities/product"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/handlers"
	Rules "github.com/wilian746/gorm-crud-generator/pkg/standart/internal/rules"
	RulesProduct "github.com/wilian746/gorm-crud-generator/pkg/standart/internal/rules/product"
	HttpStatus "github.com/wilian746/gorm-crud-generator/pkg/standart/utils/http"
	"net/http"
)

type Handler struct {
	handlers.Interface

	Controller product.Interface
	Rules      Rules.Interface
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
	productBody, err := h.getBodyAndValidate(r)
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

	productBody, err := h.getBodyAndValidate(r)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}

	if err := h.Controller.Update(ID, productBody); err != nil {
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
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusNoContent(w, r)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}

func (h *Handler) getBodyAndValidate(r *http.Request) (productBody *EntityProduct.Product, err error) {
	body, err := h.Rules.ConvertIoReaderToStruct(r.Body, productBody)
	if err != nil {
		return &EntityProduct.Product{}, errors.New("body is required")
	}

	productBody, err = EntityProduct.InterfaceToModel(body)
	if err != nil {
		return &EntityProduct.Product{}, errors.New("error on convert body to model")
	}

	return productBody, h.Rules.Validate(productBody)
}
