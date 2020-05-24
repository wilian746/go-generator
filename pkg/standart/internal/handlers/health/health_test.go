package health

import (
	"github.com/stretchr/testify/assert"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/repository/adapter"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/repository/database"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHandler(t *testing.T) {
	t.Run("Should return instance of handler", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		repository := adapter.NewAdapter(conn)
		assert.NotEmpty(t, NewHandler(repository))
	})
}

func TestHandler_Options(t *testing.T) {
	t.Run("Should return no content when call options", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		h := NewHandler(adapter.NewAdapter(conn))
		r, _ := http.NewRequest(http.MethodOptions, "/health", nil)
		w := httptest.NewRecorder()
		h.Options(w, r)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}

func TestHandler_Post(t *testing.T) {
	t.Run("Should return no content when call post", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		h := NewHandler(adapter.NewAdapter(conn))
		r, _ := http.NewRequest(http.MethodPost, "/health", nil)
		w := httptest.NewRecorder()
		h.Post(w, r)
		assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	})
}

func TestHandler_Put(t *testing.T) {
	t.Run("Should return no content when call put", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		h := NewHandler(adapter.NewAdapter(conn))
		r, _ := http.NewRequest(http.MethodPut, "/health", nil)
		w := httptest.NewRecorder()
		h.Put(w, r)
		assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	})
}

func TestHandler_Delete(t *testing.T) {
	t.Run("Should return no content when call delete", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		h := NewHandler(adapter.NewAdapter(conn))
		r, _ := http.NewRequest(http.MethodDelete, "/health", nil)
		w := httptest.NewRecorder()
		h.Delete(w, r)
		assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	})
}

func TestHandler_Get(t *testing.T) {
	t.Run("Should return OK when call get with mock", func(t *testing.T) {
		mock := &adapter.Mock{}
		mock.On("Health").Return(true)
		h := NewHandler(mock)
		r, _ := http.NewRequest(http.MethodGet, "/health", nil)
		w := httptest.NewRecorder()
		h.Get(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Should return internal_error when call get with mock", func(t *testing.T) {
		mock := &adapter.Mock{}
		mock.On("Health").Return(false)
		h := NewHandler(mock)
		r, _ := http.NewRequest(http.MethodGet, "/health", nil)
		w := httptest.NewRecorder()
		h.Get(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
	t.Run("Should return OK when call get with memory database", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		h := NewHandler(adapter.NewAdapter(conn))
		r, _ := http.NewRequest(http.MethodGet, "/health", nil)
		w := httptest.NewRecorder()
		h.Get(w, r)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
