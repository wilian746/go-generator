package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPANIC(t *testing.T) {
	t.Run("should not return panic", func(t *testing.T) {
		assert.NotPanics(t, func() { INFO("Example error", nil) }, "The code did not panic")
	})
}
