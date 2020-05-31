package product

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	ControllersProduct "github.com/wilian746/go-generator/pkg/standart-gorm/internal/controllers/product"
	_ "github.com/wilian746/go-generator/pkg/standart-gorm/internal/entities"         // import used in swagger
	_ "github.com/wilian746/go-generator/pkg/standart-gorm/internal/entities/product" // import used in swagger
	"github.com/wilian746/go-generator/pkg/standart-gorm/internal/handlers"
	RulesProduct "github.com/wilian746/go-generator/pkg/standart-gorm/internal/rules/product"
	HttpStatus "github.com/wilian746/go-generator/pkg/standart-gorm/internal/utils/http"
	"github.com/wilian746/go-generator/pkg/standart-gorm/pkg/repository/adapter"
	"net/http"
)

type Handler struct {
	handlers.Interface
	Controller ControllersProduct.Interface
	Rules      *RulesProduct.Rules
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Controller: ControllersProduct.NewController(repository),
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

// @Tags Product
// @Summary List product by id
// @ID get-one-product
// @Accept  json
// @Produce  json
// @Param ID path string true "ID of the product"
// @Success 200 {object} product.ResponseListOneProduct
// @Failure 400 {object} http.ResponseError
// @Failure 404 {object} http.ResponseError
// @Failure 500 {object} http.ResponseError
// @Router /product/{ID} [get]
func (h *Handler) getOne(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil || ID == uuid.Nil {
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

// @Tags Product
// @Summary List all products
// @ID get-all-products
// @Accept  json
// @Produce  json
// @Success 200 {object} product.ResponseListAllProduct
// @Failure 500 {object} http.ResponseError
// @Router /product [get]
func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	response, err := h.Controller.ListAll()
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, response)
}

// @Tags Product
// @Summary Create an product
// @ID post-product
// @Accept json
// @Produce json
// @Param product body product.RequestBodyToCreateOrUpdateProduct true "Body of add product"
// @Success 200 {object} product.ResponseCreateProduct
// @Failure 500 {object} http.ResponseError
// @Failure 400 {object} http.ResponseError
// @Router /product [post]
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	productBody, err := h.Rules.ConvertIoReaderToProduct(r.Body, uuid.Nil)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}
	ID, err := h.Controller.Create(productBody)
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, errors.New("error when create"))
		return
	}

	HttpStatus.StatusOK(w, r, map[string]interface{}{"id": ID.String()})
}

// @Tags Product
// @Summary Update an product
// @ID put-product
// @Accept  json
// @Produce  json
// @Param ID path string true "ID of the product"
// @Param product body product.RequestBodyToCreateOrUpdateProduct true "Body of update product"
// @Success 204
// @Failure 400 {object} http.ResponseError
// @Failure 404 {object} http.ResponseError
// @Failure 500 {object} http.ResponseError
// @Router /product/{ID} [put]
func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil || ID == uuid.Nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}

	productBody, err := h.Rules.ConvertIoReaderToProduct(r.Body, ID)
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

// @Tags Product
// @Summary Delete an product
// @ID delete-product
// @Accept  json
// @Produce  json
// @Param ID path string true "ID of the product"
// @Success 204
// @Failure 400 {object} http.ResponseError
// @Failure 404 {object} http.ResponseError
// @Failure 500 {object} http.ResponseError
// @Router /product/{ID} [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil || ID == uuid.Nil {
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
