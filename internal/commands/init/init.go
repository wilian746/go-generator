package init

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wilian746/go-generator/internal/controllers/generate/app"
	"github.com/wilian746/go-generator/internal/enums/errors"
	EnumsRepository "github.com/wilian746/go-generator/internal/enums/repository"
	EnumsRepositoryCommands "github.com/wilian746/go-generator/internal/enums/repository/commands"
	UseCaseRepository "github.com/wilian746/go-generator/internal/usecase/repository"
	"github.com/wilian746/go-generator/internal/utils/logger"
	"github.com/wilian746/go-generator/internal/utils/prompt"
	"os"
	"strings"
)

type ICommand interface {
	Cmd() *cobra.Command
	Execute(_ *cobra.Command, args []string) error
}

type Command struct {
	cmd    *cobra.Command
	prompt prompt.Interface
}

func NewInitCommand(p prompt.Interface) ICommand {
	cmd := &Command{
		prompt: p,
	}
	cmd.Init()
	return cmd
}

func (c *Command) Cmd() *cobra.Command {
	return c.cmd
}

func (c *Command) Execute(_ *cobra.Command, args []string) error {
	switch EnumsRepository.ValueOf(args[0]) {
	case EnumsRepository.Gorm:
		return c.gormInit(EnumsRepositoryCommands.ValueOf(args[1]))
	default:
		return errors.ErrInitTypeInvalid
	}
}

func (c *Command) gormInit(value EnumsRepositoryCommands.Command) error {
	switch value {
	case EnumsRepositoryCommands.App:
		return c.initApp(EnumsRepository.Gorm)
	default:
		return errors.ErrInitTypeInvalid
	}
}

func (c *Command) Init() {
	c.cmd = &cobra.Command{
		Use:     "init",
		Short:   "Initialize template application using selected repository",
		Example: "go-generator init gorm app",
		Args:    c.validateArgs,
		RunE:    c.Execute,
	}
	c.setUsageCommand()
}

func (c *Command) initApp(db EnumsRepository.Repository) error {
	pathDestiny, err := c.askPathDestiny()
	if err != nil {
		return err
	}
	moduleName, err := c.prompt.Ask("Enter module of golang project", "github.com/wilian746/go-generator/tmp")
	if err != nil || moduleName == "" {
		return errors.ErrModuleNameInvalid
	}
	return app.NewApp().CreateFoldersAndFiles(pathDestiny, moduleName, db)
}

func (c *Command) askPathDestiny() (string, error) {
	actualDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	pathDestiny, err := c.prompt.Ask("Enter the full path of the directory destiny!", actualDirectory)
	if err != nil {
		return "", errors.ErrDirectoryPathInvalid
	}
	lastChar := pathDestiny[len(pathDestiny)-1:]
	if lastChar == "/" {
		pathDestiny = strings.TrimSuffix(pathDestiny, lastChar)
	}
	return pathDestiny, nil
}

func (c *Command) setUsageCommand() {
	c.cmd.SetUsageFunc(func(command *cobra.Command) error {
		logger.PRINT("Get base of the project, handlers, controllers, repository using selected database")
		logger.PRINT(fmt.Sprintf(`
Usage:
	go-generator init [REPOSITORY] [GENERATE_TYPE]

Examples:
	%s
`, UseCaseRepository.GetAvailableCommands()))
		return nil
	})
}

func (c *Command) validateArgs(_ *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.ErrInitArgsInvalid
	}
	if !UseCaseRepository.IsValidRepositoryAndCommand(args[0], args[1]) {
		return errors.ErrArgsRepositoryOrCommandInvalid
	}
	return nil
}
