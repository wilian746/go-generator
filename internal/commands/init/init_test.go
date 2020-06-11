package init

import (
	"github.com/stretchr/testify/assert"
	"github.com/wilian746/go-generator/internal/utils/prompt"
	"testing"
)

func TestNewInitCommand(t *testing.T) {
	t.Run("Should create new command without error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewInitCommand(prompt.NewPrompt())
		})
	})
}
func TestCommand_Execute(t *testing.T) {
	path := "./tmp"
	t.Run("Should execute command exec without error", func(t *testing.T) {
		promptMock := &prompt.Mock{}
		promptMock.On("Ask").Return(path, nil)
		cobraCmd := NewInitCommand(promptMock)
		assert.NoError(t, cobraCmd.Execute(cobraCmd.Cmd(), []string{"gorm", "app"}))
	})
	t.Run("Should execute command exec with error", func(t *testing.T) {
		promptMock := &prompt.Mock{}
		promptMock.On("Ask").Return(path, nil)
		cobraCmd := NewInitCommand(promptMock)
		assert.Error(t, cobraCmd.Execute(cobraCmd.Cmd(), []string{"other-db", "app"}))
	})
	t.Run("Should execute command exec with error", func(t *testing.T) {
		promptMock := &prompt.Mock{}
		promptMock.On("Ask").Return(path, nil)
		cobraCmd := NewInitCommand(promptMock)
		assert.Error(t, cobraCmd.Execute(cobraCmd.Cmd(), []string{"gorm", "other-cmd"}))
	})
}
