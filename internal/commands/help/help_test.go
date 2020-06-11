package help

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionCommand_Execute(t *testing.T) {
	t.Run("Should execute command exec without error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			root := &cobra.Command{}
			root.AddCommand(&cobra.Command{Use: "test"})
			cobraCmd := NewHelpCommand(root)
			assert.NoError(t, cobraCmd.Execute(cobraCmd.Cmd(), []string{}))
		})
	})
}
