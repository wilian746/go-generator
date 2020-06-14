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
)

const ModuleRestart = "restart"

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
	hostname, username, projectName, err := c.getHostnameUsernameProjectName()
	if err != nil {
		return err
	}
	moduleName, err := c.askModuleCorrectly(hostname, username, projectName)
	if err != nil {
		return err
	}
	if moduleName == ModuleRestart {
		return c.initApp(db)
	}
	return c.startCreateFoldersAndFiles(projectName, moduleName, db)
}

func (c *Command) startCreateFoldersAndFiles(projectName, moduleName string, db EnumsRepository.Repository) error {
	pathDestiny, err := c.getDirectoryByProjectNameBySelect(projectName)
	if err != nil {
		return err
	}
	logger.INFO(fmt.Sprintf(`The template will generate with:
	* Path of destiny = %s
	* Module name of the project = %s
`, pathDestiny, moduleName), nil)
	return app.NewApp().CreateFoldersAndFiles(pathDestiny, moduleName, db)
}

func (c *Command) getHostnameUsernameProjectName() (string, string, string, error) {
	hostname, err := c.selectHostName()
	if err != nil {
		return "", "", "", err
	}
	username, err := c.askUsername()
	if err != nil {
		return "", "", "", err
	}
	projectName, err := c.askProjectName()
	if err != nil {
		return "", "", "", err
	}
	return hostname, username, projectName, nil
}

func (c *Command) getDirectoryByProjectNameBySelect(projectName string) (string, error) {
	actualDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := fmt.Sprintf("%s/%s", actualDirectory, projectName)
	response, err := c.prompt.Select(
		fmt.Sprintf("The project will generated in path: `%s` is correct?", path), []string{"Y", "N"})
	if err != nil {
		return "", errors.ErrDirectoryPathInvalid
	}
	if response == "N" {
		return c.getDirectoryByProjectNameByAsk(projectName, path)
	}
	return path, nil
}

func (c *Command) getDirectoryByProjectNameByAsk(projectName, path string) (string, error) {
	hostName, err := c.prompt.Ask("Enter with path of destiny of the your project", path)
	if err != nil {
		return "", errors.ErrDirectoryPathInvalid
	}
	if hostName == "" {
		logger.ERROR("Invalid response!", errors.ErrDirectoryPathInvalid)
		return c.getDirectoryByProjectNameBySelect(projectName)
	}
	return hostName, nil
}

func (c *Command) selectHostName() (string, error) {
	logger.INFO("Create your module: STEP (1/3)", nil)
	response, err := c.prompt.Select("Select your host of the repository",
		[]string{"github.com", "gitlab.com", "bitbucket.org", "other"})
	if err != nil {
		return "", errors.ErrDirectoryPathInvalid
	}
	if response == "other" {
		return c.askHostName()
	}
	return response, nil
}

func (c *Command) askHostName() (string, error) {
	hostName, err := c.prompt.Ask("Enter with your hostname", "")
	if err != nil {
		return "", errors.ErrHostNameInvalid
	}
	if hostName == "" {
		logger.ERROR("Invalid response!", errors.ErrHostNameInvalid)
		return c.selectHostName()
	}
	return hostName, nil
}

func (c *Command) askUsername() (string, error) {
	logger.INFO("Create your module: STEP (2/3)", nil)
	username, err := c.prompt.Ask("Enter with the your username of the repository", "")
	if err != nil {
		return "", errors.ErrUsernameInvalid
	}
	if username == "" {
		logger.ERROR("Invalid response!", errors.ErrUsernameInvalid)
		return c.selectHostName()
	}
	return username, nil
}

func (c *Command) askProjectName() (string, error) {
	logger.INFO("Create your module: STEP (3/3)", nil)
	projectName, err := c.prompt.Ask("Enter with the your project name", "")
	if err != nil {
		return "", errors.ErrProjectNameInvalid
	}
	if projectName == "" {
		logger.ERROR("Invalid response!", errors.ErrProjectNameInvalid)
		return c.selectHostName()
	}
	return projectName, nil
}

func (c *Command) askModuleCorrectly(hostName, username, projectName string) (string, error) {
	moduleName := fmt.Sprintf("%s/%s/%s", hostName, username, projectName)
	label := fmt.Sprintf("The module created is: `%s` is correct?", moduleName)
	response, err := c.prompt.Select(label, []string{"Y", "N"})
	if err != nil {
		return "", err
	}
	if response == "N" {
		return ModuleRestart, nil
	}
	return moduleName, nil
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
