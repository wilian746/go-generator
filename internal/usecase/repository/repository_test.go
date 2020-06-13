package repository

import (
	"github.com/stretchr/testify/assert"
	enumsRepository "github.com/wilian746/go-generator/internal/enums/repository"
	enumsCommands "github.com/wilian746/go-generator/internal/enums/repository/commands"
	"testing"
)

func TestIsValidRepositoryAndCommand(t *testing.T) {
	t.Run("Should pass gorm and app to validate and return true", func(t *testing.T) {
		assert.True(t, IsValidRepositoryAndCommand("gorm", "app"))
	})
	t.Run("Should pass mongo and app to validate and return false", func(t *testing.T) {
		assert.False(t, IsValidRepositoryAndCommand("mongo", "app"))
	})
	t.Run("Should pass gorm and controller to validate and return false", func(t *testing.T) {
		assert.False(t, IsValidRepositoryAndCommand("gorm", "controller"))
	})
	t.Run("Should pass mongo and controller to validate and return false", func(t *testing.T) {
		assert.False(t, IsValidRepositoryAndCommand("mongo", "controller"))
	})
}

func TestGetCommandsValidByRepository(t *testing.T) {
	t.Run("Should return one command whith app, when get all commands from repository gorm", func(t *testing.T) {
		commands := GetCommandsValidByRepository(enumsRepository.Gorm.String())
		assert.Len(t, commands, 1)
		assert.Equal(t, commands[0], enumsCommands.App.String())
	})
	t.Run("Should return commands empty, when get all commands from repository unknown", func(t *testing.T) {
		commands := GetCommandsValidByRepository(enumsRepository.Unknown.String())
		assert.Len(t, commands, 0)
	})
}

func TestGetAvailableCommands(t *testing.T) {
	t.Run("Check if exists all commands available", func(t *testing.T) {
		examples := GetAvailableCommands()
		assert.Contains(t, examples, "go-generator init gorm app")
	})
}
