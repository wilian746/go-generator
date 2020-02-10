package product

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/adapter"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/controllers/product"
	EntityProduct "github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities/product"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/handlers"
	"net/http"
	"strconv"
)

type Handler struct {
	handlers.Interface

	Controller product.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Controller: product.NewController(repository),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if chi.URLParam(r, "ID") != "" {
		ID, err := strconv.ParseUint(chi.URLParam(r, "ID"), 10, 32)
		if err != nil || uint(ID) <= uint(0) {
			w.WriteHeader(http.StatusBadRequest)
			bytes, _ := json.Marshal(errors.New("ID is required bigger then zero"))
			_, _ = w.Write(bytes)
			return
		}
		response, err := h.Controller.ListOne(uint(ID))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			bytes, _ := json.Marshal(err)
			_, _ = w.Write(bytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		bytes, _ := json.Marshal(response)
		_, _ = w.Write(bytes)
		return
	}

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
	body, err := EntityProduct.ConvertIoReaderToStruct(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		bytes, _ := json.Marshal(errors.New("body is required"))
		_, _ = w.Write(bytes)
		return
	}
	ID, err := h.Controller.Create(body)
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
	ID, err := strconv.ParseUint(chi.URLParam(r, "ID"), 10, 32)
	if err != nil || uint(ID) <= uint(0) {
		w.WriteHeader(http.StatusBadRequest)
		bytes, _ := json.Marshal(errors.New("ID is required bigger then zero"))
		_, _ = w.Write(bytes)
		return
	}
	body, err := EntityProduct.ConvertIoReaderToStruct(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		bytes, _ := json.Marshal(errors.New("body is required"))
		_, _ = w.Write(bytes)
		return
	} else if err := h.Controller.Update(uint(ID), body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		bytes, _ := json.Marshal(errors.New("error when update"))
		_, _ = w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID, err := strconv.ParseUint(chi.URLParam(r, "ID"), 10, 32)
	if err != nil || uint(ID) <= uint(0) {
		w.WriteHeader(http.StatusBadRequest)
		bytes, _ := json.Marshal(errors.New("ID is required bigger then zero"))
		_, _ = w.Write(bytes)
		return
	}

	if err := h.Controller.Remove(uint(ID)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		bytes, _ := json.Marshal(err)
		_, _ = w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}