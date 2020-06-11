package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPANIC(t *testing.T) {
	t.Run("should not return panic without data", func(t *testing.T) {
		assert.NotPanics(t, func() { INFO("Example error", nil) })
	})
	t.Run("should not return panic with data", func(t *testing.T) {
		assert.NotPanics(t, func() { INFO("Example error", "some text") })
	})
}
