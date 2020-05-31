package server

import (
	"fmt"
	"github.com/wilian746/go-generator/internal/enums/database"
	"github.com/wilian746/go-generator/internal/enums/files"
	"github.com/wilian746/go-generator/internal/enums/folders"
	"github.com/wilian746/go-generator/internal/utils/environment"
	"github.com/wilian746/go-generator/internal/utils/logger"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const ImportModuleName = "github.com/wilian746/go-generator"

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
		fileContent, err := s.getFileStringFromRepository("pkg/"+databaseFolderName, string(dir))
		if err != nil {
			return err
		}
		if dir != files.Readme {
			fileContent = s.replaceImportsToModuleName(fileContent, moduleName)
		}
		err = s.writeContent(pathDestiny, string(dir), fileContent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) writeContent(pathDestiny, dir string, fileContent []byte) error {
	absPath := fmt.Sprintf("%s/%s", pathDestiny, dir)
	err := ioutil.WriteFile(absPath, fileContent, os.ModePerm)
	if err != nil {
		return err
	}
	logger.INFO("File generated with success: "+absPath, nil)
	return nil
}

func (s *Server) replaceImportsToModuleName(fileContent []byte, moduleName string) []byte {
	importModule := ImportModuleName + "/pkg/standart-gorm"

	fileContentReplaced := strings.ReplaceAll(string(fileContent), importModule, moduleName)

	return []byte(fileContentReplaced)
}

func (s *Server) replaceModuleToModuleName(fileContent []byte, moduleName string) []byte {
	fileContentReplaced := strings.ReplaceAll(string(fileContent), ImportModuleName, moduleName)

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
	latestVersionStable := environment.GetEnvString("GO_GENERATOR_TAG_NAME", "master")
	urlBase := "https://raw.githubusercontent.com/wilian746/go-generator/" + latestVersionStable
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
