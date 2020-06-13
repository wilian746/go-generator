package app

import (
	"fmt"
	"github.com/wilian746/go-generator/internal/enums/files"
	"github.com/wilian746/go-generator/internal/enums/folders"
	EnumsRepository "github.com/wilian746/go-generator/internal/enums/repository"
	"github.com/wilian746/go-generator/internal/utils/environment"
	"github.com/wilian746/go-generator/internal/utils/github"
	"github.com/wilian746/go-generator/internal/utils/logger"
	"io/ioutil"
	"os"
	"strings"
)

const ImportModuleName = "github.com/wilian746/go-generator"

type Interface interface {
	CreateFoldersAndFiles(pathDestiny, moduleName string, db EnumsRepository.Database) error
}

type App struct {
	db EnumsRepository.Database
}

func NewApp() Interface {
	return &App{}
}

func (a *App) CreateFoldersAndFiles(pathDestiny, moduleName string, db EnumsRepository.Database) error {
	a.db = db
	if err := a.createFolders(pathDestiny); err != nil {
		return err
	}
	if err := a.factoryCopyContent(pathDestiny, moduleName); err != nil {
		return err
	}
	return a.copyDefaultFiles(pathDestiny, moduleName)
}

func (a *App) factoryCopyContent(destiny, moduleName string) error {
	switch a.db {
	case EnumsRepository.Gorm:
		return a.createFiles(destiny, moduleName, "standart-gorm")
	default:
		return nil
	}
}

func (a *App) createFolders(pathDestiny string) error {
	for _, dir := range a.getFoldersSliceToCreateByDatabase() {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", pathDestiny, dir), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) getFoldersSliceToCreateByDatabase() []folders.Folders {
	switch a.db {
	case EnumsRepository.Gorm:
		list := folders.Values()
		list = append(list, folders.ValuesGorm()...)
		return list
	default:
		return []folders.Folders{}
	}
}

func (a *App) getFilesSliceToCreateByDatabase() []files.Files {
	switch a.db {
	case EnumsRepository.Gorm:
		list := files.Values()
		list = append(list, files.ValuesGorm()...)
		return list
	default:
		return []files.Files{}
	}
}

func (a *App) createFiles(pathDestiny, moduleName, databaseFolderName string) error {
	for _, dir := range a.getFilesSliceToCreateByDatabase() {
		fileContent, err := a.getFileStringFromRepository("pkg/"+databaseFolderName, string(dir))
		if err != nil {
			return err
		}
		if dir != files.Readme {
			fileContent = a.replaceImportsToModuleName(fileContent, moduleName)
		}
		err = a.writeContent(pathDestiny, string(dir), fileContent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) writeContent(pathDestiny, dir string, fileContent []byte) error {
	absPath := fmt.Sprintf("%s/%s", pathDestiny, dir)
	err := ioutil.WriteFile(absPath, fileContent, os.ModePerm)
	if err != nil {
		return err
	}
	logger.PRINT("File generated with success: " + absPath)
	return nil
}

func (a *App) replaceImportsToModuleName(fileContent []byte, moduleName string) []byte {
	importModule := ImportModuleName + "/pkg/standart-gorm"

	fileContentReplaced := strings.ReplaceAll(string(fileContent), importModule, moduleName)

	return []byte(fileContentReplaced)
}

func (a *App) replaceModuleToModuleName(fileContent []byte, moduleName string) []byte {
	fileContentReplaced := strings.ReplaceAll(string(fileContent), ImportModuleName, moduleName)

	return []byte(fileContentReplaced)
}

func (a *App) copyDefaultFiles(pathDestiny, moduleName string) error {
	for _, dir := range files.ValuesNoGO() {
		fileContent, err := a.getFileStringFromRepository("", string(dir))
		if err != nil {
			return err
		}
		if dir == files.GoMod {
			fileContent = a.replaceModuleToModuleName(fileContent, moduleName)
		}
		err = a.writeContent(pathDestiny, string(dir), fileContent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) getFileStringFromRepository(databaseFolderName, dir string) ([]byte, error) {
	latestVersionStable := environment.GetEnvString("GO_GENERATOR_TAG_NAME", "master")
	routerGithub := fmt.Sprintf("%s/%s/%s", latestVersionStable, databaseFolderName, dir)
	return github.GetFileFromGithub(routerGithub)
}
