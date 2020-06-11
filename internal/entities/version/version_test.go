package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	t.Run("should set and get version correctly", func(t *testing.T) {
		c := &Version{}
		assert.NoError(t, c.Set("0.0.0"))
		assert.Equal(t, c.String(), "0.0.0")
		assert.Empty(t, c.Type())
	})
}
