package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionCommand_Execute(t *testing.T) {
	t.Run("Should execute command exec without error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			cobraCmd := NewVersionCommand()
			cobraCmd.CmdVersion()
		})
	})
}
