package server

import (
	"fmt"
	"github.com/wilian746/go-generator/internal/enums/database"
	"github.com/wilian746/go-generator/internal/enums/files"
	"github.com/wilian746/go-generator/internal/enums/folders"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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
	if err := s.factoryCopyContent(pathDestiny, moduleName, db); err != nil {
		return err
	}
	return s.copyDefaultFiles(pathDestiny, moduleName)
}

func (s *Server) factoryCopyContent(destiny, moduleName string, db database.Database) error {
	switch db {
	case database.Gorm:
		return s.createFiles(destiny, moduleName, "standart-gorm")
	default:
		return nil
	}
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

func (s *Server) createFiles(pathDestiny, moduleName, databaseFolderName string) error {
	for _, dir := range files.Values() {
		fileContent, err := s.getFileStringFromRepository("pkg/" + databaseFolderName, string(dir))
		if err != nil {
			return err
		}
		fileContent = s.replaceImportsToModuleName(fileContent, moduleName)
		err = s.writeContent(pathDestiny, string(dir), fileContent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) readContent(absPath string) ([]byte, error) {
	actualFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		return []byte{}, err
	}
	return actualFile, nil
}

func (s *Server) writeContent(pathDestiny, dir string, fileContent []byte) error {
	absPath := fmt.Sprintf("%s/%s", pathDestiny, dir)
	err := ioutil.WriteFile(absPath, fileContent, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) replaceImportsToModuleName(fileContent []byte, moduleName string) []byte {
	importModule := "github.com/wilian746/go-generator/pkg/standart-gorm"

	fileContentReplaced := strings.ReplaceAll(string(fileContent), importModule, moduleName)

	return []byte(fileContentReplaced)
}

func (s *Server) replaceModuleToModuleName(fileContent []byte, moduleName string) []byte {
	importModule := "github.com/wilian746/go-generator"

	fileContentReplaced := strings.ReplaceAll(string(fileContent), importModule, moduleName)

	return []byte(fileContentReplaced)
}

func (s *Server) copyDefaultFiles(pathDestiny, moduleName string) error {
	for _, dir := range files.ValuesNoGO() {
		fileContent, err := s.getFileStringFromRepository("", string(dir))
		if err != nil {
			return err
		}
		if dir == files.GoMod {
			fileContent = s.replaceModuleToModuleName(fileContent, moduleName)
		}
		err = s.writeContent(pathDestiny, string(dir), fileContent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) getFileStringFromRepository(databaseFolderName, dir string) ([]byte, error) {
	urlBase := "https://raw.githubusercontent.com/wilian746/go-generator/master"
	url := fmt.Sprintf("%s/%s/%s", urlBase, databaseFolderName, dir)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, err
	}
	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}