package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfig(t *testing.T) {
	t.Run("Should return configs", func(t *testing.T) {
		c := GetConfig()
		assert.NotEmpty(t, c.DatabaseURI)
		assert.NotEmpty(t, c.Dialect)
		assert.NotEqual(t, c.Port, 0)
		assert.NotEqual(t, c.Timeout, 0)
	})
}
