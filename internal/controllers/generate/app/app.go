package app

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/progress"
	printtree "github.com/wilian746/go-generator/internal/controllers/generate/app/print_tree"
	"github.com/wilian746/go-generator/internal/enums/files"
	"github.com/wilian746/go-generator/internal/enums/folders"
	EnumsRepository "github.com/wilian746/go-generator/internal/enums/repository"
	"github.com/wilian746/go-generator/internal/utils/environment"
	"github.com/wilian746/go-generator/internal/utils/github"
	"github.com/wilian746/go-generator/internal/utils/logger"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const ImportModuleName = "github.com/wilian746/go-generator"

type Interface interface {
	CreateFoldersAndFiles(pathDestiny, moduleName string, db EnumsRepository.Repository) error
}

type App struct {
	db                 EnumsRepository.Repository
	pathOfFilesCreated []string
	moduleName         string
}

func NewApp() Interface {
	return &App{}
}

func (a *App) CreateFoldersAndFiles(pathDestiny, moduleName string, db EnumsRepository.Repository) error {
	a.db = db
	a.moduleName = moduleName
	if err := a.createFolders(pathDestiny); err != nil {
		return err
	}
	logger.INFO("", nil)
	if err := a.factoryCopyContent(pathDestiny); err != nil {
		return err
	}
	if err := a.copyDefaultFiles(pathDestiny); err != nil {
		return err
	}
	return a.printAllFilesGenerated(pathDestiny)
}

func (a *App) printAllFilesGenerated(pathDestiny string) error {
	logger.INFO("============ All files was generated with success! ============", nil)
	paths := []string{}
	for _, value := range a.pathOfFilesCreated {
		paths = append(paths, strings.ReplaceAll(value, pathDestiny, ""))
	}
	printtree.NewTree(paths, pathDestiny).Print()
	return nil
}

func (a *App) factoryCopyContent(destiny string) error {
	switch a.db {
	case EnumsRepository.Gorm:
		return a.createFiles(destiny, "standart-gorm")
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

func (a *App) createFiles(pathDestiny, databaseFolderName string) error {
	totalFiles := a.getFilesSliceToCreateByDatabase()
	pw := a.getProgressInstance()
	go pw.Render()
	logger.INFO(fmt.Sprintf("Total of files of %s to generate: %v", databaseFolderName, len(totalFiles)), nil)
	if err := a.loopToCreateFiles(totalFiles, databaseFolderName, pathDestiny, pw); err != nil {
		return err
	}
	a.appendProgressPercentage(len(totalFiles)+1, len(totalFiles), pw)
	time.Sleep(time.Second * 1)
	pw.Stop()
	return nil
}

func (a *App) loopToCreateFiles(
	totalFiles []files.Files, databaseFolderName, pathDestiny string, pw progress.Writer) error {
	for index, dir := range totalFiles {
		fileContent, err := a.getFileStringFromRepository("pkg/"+databaseFolderName, string(dir))
		if err != nil {
			return err
		}
		if dir != files.Readme {
			fileContent = a.replaceImportsToModuleName(fileContent)
		}
		if err := a.writeContent(pathDestiny, string(dir), fileContent); err != nil {
			return err
		}
		a.appendProgressPercentage(index, len(totalFiles), pw)
	}
	return nil
}
func (a *App) writeContent(pathDestiny, dir string, fileContent []byte) error {
	absPath := fmt.Sprintf("%s/%s", pathDestiny, dir)
	err := ioutil.WriteFile(absPath, fileContent, os.ModePerm)
	if err != nil {
		return err
	}
	a.pathOfFilesCreated = append(a.pathOfFilesCreated, absPath)
	return nil
}

func (a *App) replaceImportsToModuleName(fileContent []byte) []byte {
	importModule := ImportModuleName + "/pkg/standart-gorm"

	fileContentReplaced := strings.ReplaceAll(string(fileContent), importModule, a.moduleName)

	return []byte(fileContentReplaced)
}

func (a *App) replaceModuleToModuleName(fileContent []byte) []byte {
	fileContentReplaced := strings.ReplaceAll(string(fileContent), ImportModuleName, a.moduleName)

	return []byte(fileContentReplaced)
}

func (a *App) copyDefaultFiles(pathDestiny string) error {
	totalFiles := files.ValuesNoGO()
	pw := a.getProgressInstance()
	go pw.Render()
	logger.INFO(fmt.Sprintf("Total of files default to generate: %v", len(totalFiles)), nil)
	if err := a.loopToCreateFilesNoGo(pathDestiny, totalFiles, pw); err != nil {
		return err
	}
	a.appendProgressPercentage(len(totalFiles)+1, len(totalFiles), pw)
	time.Sleep(time.Second * 1)
	pw.Stop()
	return nil
}

func (a *App) loopToCreateFilesNoGo(pathDestiny string, totalFiles []files.NoGo, pw progress.Writer) error {
	for index, dir := range files.ValuesNoGO() {
		fileContent, err := a.getFileStringFromRepository("", string(dir))
		if err != nil {
			return err
		}
		if dir == files.GoMod {
			fileContent = a.replaceModuleToModuleName(fileContent)
		}
		if err := a.writeContent(pathDestiny, string(dir), fileContent); err != nil {
			return err
		}
		a.appendProgressPercentage(index, len(totalFiles), pw)
	}
	return nil
}

func (a *App) getFileStringFromRepository(databaseFolderName, dir string) ([]byte, error) {
	latestVersionStable := environment.GetEnvString("GO_GENERATOR_TAG_NAME", "master")
	routerGithub := fmt.Sprintf("%s/%s/%s", latestVersionStable, databaseFolderName, dir)
	return github.GetFileFromGithub(routerGithub)
}

func (a *App) getProgressInstance() progress.Writer {
	pw := progress.NewWriter()
	pw.ShowOverallTracker(false)
	pw.ShowTime(false)
	pw.ShowTracker(true)
	pw.ShowValue(false)
	pw.ShowPercentage(true)
	pw.Style().Options.PercentFormat = "%5.2f%%"
	pw.Style().Options.DoneString = ""
	pw.Style().Options.Separator = ""
	return pw
}

func (a *App) appendProgressPercentage(index, total int, pw progress.Writer) {
	tracker := &progress.Tracker{
		Units: progress.UnitsDefault,
		Total: int64(total),
	}
	tracker.SetValue(int64(index))
	tracker.IsDone()
	pw.AppendTracker(tracker)
}
