package repository

import (
	"github.com/stretchr/testify/assert"
	enumsRepository "github.com/wilian746/go-generator/internal/enums/repository"
	enumsCommands "github.com/wilian746/go-generator/internal/enums/repository/commands"
	"testing"
)

func TestRepository_GetListCommands(t *testing.T) {
	t.Run("Should pass gorm and app to validate and return true", func(t *testing.T) {
		r := Repository{
			Name: enumsRepository.Gorm,
			Commands: []enumsCommands.Command{
				enumsCommands.App,
			},
		}
		assert.Equal(t, r.GetListCommands(), []string{"app"})
	})
}
