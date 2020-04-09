package product

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/wilian746/gorm-crud-generator/"
	"github.com/wilian746/gorm-crud-generator/internal/controllers/product"
	"github.com/wilian746/gorm-crud-generator/internal/handlers"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/adapter"
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
	w.Header().Set("Content-Type", "application/json")

	if chi.URLParam(r, "ID") != "" {
		h.getOne(w, r)
	} else {
		h.getAll(w, r)
	}
}

func (h *Handler) getOne(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		bytes, _ := json.Marshal(errors.New("ID is not uuid valid"))
		_, _ = w.Write(bytes)
		return
	}

	response, err := h.Controller.ListOne(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		bytes, _ := json.Marshal(err)
		_, _ = w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal(response)
	_, _ = w.Write(bytes)
}

func (h *Handler) getAll(w http.ResponseWriter, _ *http.Request) {
	response, err := h.Controller.ListAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		bytes, _ := json.Marshal(err)
		_, _ = w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal(response)
	_, _ = w.Write(bytes)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productBody, err := h.getBodyAndValidate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		bytes, _ := json.Marshal(err)
		_, _ = w.Write(bytes)
		return
	}

	ID, err := h.Controller.Create(productBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		bytes, _ := json.Marshal(errors.New("error when update"))
		_, _ = w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal(ID)
	_, _ = w.Write(bytes)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		bytes, _ := json.Marshal(errors.New("ID is not uuid valid"))
		_, _ = w.Write(bytes)
		return
	}

	productBody, err := h.getBodyAndValidate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		bytes, _ := json.Marshal(err)
		_, _ = w.Write(bytes)
		return
	}

	if err := h.Controller.Update(ID, productBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		bytes, _ := json.Marshal(errors.New("error when update"))
		_, _ = w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		bytes, _ := json.Marshal(errors.New("ID is not uuid valid"))
		_, _ = w.Write(bytes)
		return
	}

	if err := h.Controller.Remove(ID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		bytes, _ := json.Marshal(err)
		_, _ = w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) Options(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) getBodyAndValidate(r *http.Request) (*EntityProduct.Product, error) {

	body, err := h.Rules.ConvertIoReaderToStruct(r.Body)
	if err != nil {
		return &EntityProduct.Product{}, errors.New("body is required")
	}

	productBody, err := EntityProduct.InterfaceToModel(body)
	if err != nil {
		return &EntityProduct.Product{}, errors.New("error on convert body to model")
	}

	return productBody, h.Rules.Validate(productBody)
}
