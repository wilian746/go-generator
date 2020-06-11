package app

import (
	"github.com/stretchr/testify/assert"
	"github.com/wilian746/go-generator/internal/enums/repository"
	"testing"
)

func TestServer_CreateFoldersAndFiles(t *testing.T) {
	t.Run("Create default folders without error", func(t *testing.T) {
		path := "../../../../../tmp"
		module := "github.com/wilian746/tmp"
		err := NewApp().CreateFoldersAndFiles(path, module, repository.Gorm)
		assert.NoError(t, err)
	})
}
