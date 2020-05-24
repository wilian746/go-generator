package server

import (
	"fmt"
	"github.com/wilian746/gorm-crud-generator/internal/enums/files"
	"github.com/wilian746/gorm-crud-generator/internal/enums/folders"
	"os"
)

type Interface interface {
	CreateFoldersAndFiles(pathDestiny, moduleName string) error
}

type Server struct{}

func NewServer() Interface {
	return &Server{}
}

func (s *Server) CreateFoldersAndFiles(pathDestiny, moduleName string) error {
	for _, dir := range folders.Values() {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", pathDestiny, dir), os.ModePerm)
		if err != nil {
			return err
		}
	}
	for _, dir := range files.Values() {
		file, err := os.Create(fmt.Sprintf("%s/%s", pathDestiny, dir))
		if err != nil {
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}
	}
	return nil
}
