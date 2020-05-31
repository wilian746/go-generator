package prompt

import (
	"testing"
)

func TestPrompt_Ask(t *testing.T) {
	t.Run("Should run command without panics", func(t *testing.T) {
		NewPrompt()
	})
}
