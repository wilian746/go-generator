package prompt

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPrompt(t *testing.T) {
	t.Run("Should run command without panics", func(t *testing.T) {
		assert.NotNil(t, NewPrompt())
	})
	t.Run("Should mock return error", func(t *testing.T) {
		mock := &Mock{}
		mock.On("Ask").Return("", errors.New("some error"))
		res, err := mock.Ask("", "")
		assert.Empty(t, res)
		assert.Error(t, err)
	})
	t.Run("Should mock return nil err", func(t *testing.T) {
		mock := &Mock{}
		mock.On("Ask").Return("", nil)
		res, err := mock.Ask("", "")
		assert.Empty(t, res)
		assert.NoError(t, err)
	})
	t.Run("Should mock return nil err and response not empty", func(t *testing.T) {
		mock := &Mock{}
		mock.On("Ask").Return("some text", nil)
		res, err := mock.Ask("", "")
		assert.NotEmpty(t, res)
		assert.NoError(t, err)
	})
}
