package version

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionCommand_Execute(t *testing.T) {
	t.Run("Should execute command exec without error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			cobraCmd := NewVersionCommand(&cobra.Command{})
			assert.NoError(t, cobraCmd.Execute(cobraCmd.Cmd(), []string{}))
		})
	})
}
