package routes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_NewCors(t *testing.T) {
	t.Run("Should not nil response", func(t *testing.T) {
		assert.NotNil(t, NewConfig())
	})
}

func TestConfig_SetTimeout(t *testing.T) {
	t.Run("Should not nil response", func(t *testing.T) {
		c := NewConfig()
		assert.NotNil(t, c.SetTimeout(1))
	})
}

func TestConfig_GetTimeout(t *testing.T) {
	t.Run("Should not nil response", func(t *testing.T) {
		c := NewConfig()
		assert.NotNil(t, c.GetTimeout())
	})
}

func TestConfig_Cors(t *testing.T) {
	t.Run("Should not nil response", func(t *testing.T) {
		c := NewConfig()
		assert.NotNil(t, c.Cors(nil))
	})
}
