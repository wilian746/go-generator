package server

import (
	"fmt"
	"github.com/wilian746/go-generator/internal/enums/database"
	"github.com/wilian746/go-generator/internal/enums/files"
	"github.com/wilian746/go-generator/internal/enums/folders"
	"os"
)

type Interface interface {
	CreateFoldersAndFiles(pathDestiny, moduleName string, db database.Database) error
}

type Server struct{}

func NewServer() Interface {
	return &Server{}
}

func (s *Server) CreateFoldersAndFiles(pathDestiny, moduleName string, db database.Database) error {
	if err := s.createFolders(pathDestiny); err != nil {
		return err
	}
	if err := s.createFiles(pathDestiny); err != nil {
		return err
	}
	if err := s.factoryCopyContent(pathDestiny, db); err != nil {
		return err
	}
	return nil
}

func (s *Server) createFolders(pathDestiny string) error {
	for _, dir := range folders.Values() {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", pathDestiny, dir), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) createFiles(pathDestiny string) error {
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

func (s *Server) factoryCopyContent(destiny string, db database.Database) error {
	switch db {
	case database.Gorm:
		return s.copyContent(destiny, "standart-gorm")
	default:
		return nil
	}
}

func (s *Server) copyContent(destiny string, databaseFolderName string) error {
	return nil
}