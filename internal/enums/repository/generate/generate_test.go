package generate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnum(t *testing.T) {
	t.Run("Should return valid repository", func(t *testing.T) {
		v := Values()
		assert.Equal(t, v, []Command{App})
	})
	t.Run("Should return app repository", func(t *testing.T) {
		assert.Equal(t, ValueOf("app"), App)
	})
	t.Run("Should return unknown repository", func(t *testing.T) {
		assert.Equal(t, ValueOf("other"), Unknown)
	})
	t.Run("Should return invalid enum", func(t *testing.T) {
		assert.False(t, Valid("other"))
	})
	t.Run("Should return valid enum", func(t *testing.T) {
		assert.True(t, Valid("app"))
	})
}
