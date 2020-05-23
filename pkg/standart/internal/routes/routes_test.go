package routes

import (
	"github.com/stretchr/testify/assert"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/adapter"
	"testing"
)

func TestNewRouter(t *testing.T) {
	t.Run("Should not return empty instance", func(t *testing.T) {
		assert.NotNil(t, NewRouter())
	})
}

func TestRouter_SetRouters(t *testing.T) {
	t.Run("Should set routes and not return panics", func(t *testing.T) {
		r := NewRouter()
		assert.NotPanics(t, func() {
			mock := &adapter.Mock{}
			r.SetRouters(mock)
		})
	})
}