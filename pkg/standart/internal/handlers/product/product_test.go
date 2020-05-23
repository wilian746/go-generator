package product

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/adapter"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/instance"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/rules/product"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHandler(t *testing.T) {
	t.Run("Should return instance of handler", func(t *testing.T) {
		conn := instance.GetConnection("sqlite3", ":memory:")
		repository := adapter.NewAdapter(conn)
		assert.NotEmpty(t, NewHandler(repository))
	})
}

func TestHandler_Options(t *testing.T) {
	t.Run("Should return no content when call options", func(t *testing.T) {
		conn := instance.GetConnection("sqlite3", ":memory:")
		h := NewHandler(adapter.NewAdapter(conn))
		r, _ := http.NewRequest(http.MethodOptions, "/product", nil)
		w := httptest.NewRecorder()
		h.Options(w, r)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}

func TestHandler_Get(t *testing.T) {
	t.Run("Should return ok when call getAll", func(t *testing.T) {
		conn := instance.GetConnection("sqlite3", ":memory:")
		rules := product.NewRules()
		rules.Migrate(conn, rules.GetMock())
		h := NewHandler(adapter.NewAdapter(conn))
		r, _ := http.NewRequest(http.MethodGet, "/product", nil)
		w := httptest.NewRecorder()
		h.Get(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Should return internal_error without migrate when call getAll", func(t *testing.T) {
		conn := instance.GetConnection("sqlite3", ":memory:")
		h := NewHandler(adapter.NewAdapter(conn))
		r, _ := http.NewRequest(http.MethodGet, "/product", nil)
		w := httptest.NewRecorder()
		h.Get(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
	t.Run("Should return bad request when call getOne", func(t *testing.T) {
		conn := instance.GetConnection("sqlite3", ":memory:")
		rules := product.NewRules()
		productMock := rules.GetMock()
		rules.Migrate(conn, productMock)
		h := NewHandler(adapter.NewAdapter(conn))
		r, _ := http.NewRequest(http.MethodGet, "/product/123", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("ID", "123")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		w := httptest.NewRecorder()
		h.Get(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("Should return notfound when call getOne", func(t *testing.T) {
		conn := instance.GetConnection("sqlite3", ":memory:")
		ID := uuid.New()
		rules := product.NewRules()
		productMock := rules.GetMock()
		productMock.ID = ID
		rules.Migrate(conn, productMock)
		h := NewHandler(adapter.NewAdapter(conn))
		r, _ := http.NewRequest(http.MethodGet, "/product/"+ID.String(), nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("ID", ID.String())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		w := httptest.NewRecorder()
		h.Get(w, r)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("Should return internal_error without migrate when call getOne", func(t *testing.T) {
		conn := instance.GetConnection("sqlite3", ":memory:")
		h := NewHandler(adapter.NewAdapter(conn))
		ID := uuid.New().String()
		r, _ := http.NewRequest(http.MethodGet, "/product/"+ID, nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("ID", ID)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		w := httptest.NewRecorder()
		h.Get(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
