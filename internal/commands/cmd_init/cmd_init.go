package cmdinit

import (
	"github.com/spf13/cobra"
	"github.com/wilian746/go-generator/internal/commands/generate/server"
	"github.com/wilian746/go-generator/internal/enums/database"
	"github.com/wilian746/go-generator/internal/enums/errors"
	"github.com/wilian746/go-generator/internal/utils/prompt"
	"os"
	"strings"
)

const (
	App = "app"
)

type Interface interface {
	Execute(cmd *cobra.Command, args []string) error
	Cmd() *cobra.Command
	Init()
}

type Command struct {
	cmd    *cobra.Command
	prompt prompt.Interface
}

func NewInitCommand() Interface {
	cmd := &Command{
		prompt: prompt.NewPrompt(),
	}
	cmd.Init()
	return cmd
}

func (c *Command) Cmd() *cobra.Command {
	return c.cmd
}

func (c *Command) Execute(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.ErrInitTypeInvalid
	}
	if err := c.validateArgs(args); err != nil {
		return err
	}

	return c.factoryDatabase(args)
}

func (c *Command) validateArgs(args []string) error {
	for _, db := range database.Values() {
		if args[0] == string(db) {
			dbGenerators, err := c.getValidDBGenerator(db)
			if err != nil {
				return err
			}
			for _, dbGenerator := range dbGenerators {
				if args[1] == dbGenerator {
					return nil
				}
			}
			return errors.ErrInitTypeInvalid
		}
	}
	return errors.ErrInitTypeInvalid
}

func (c *Command) getValidDBGenerator(db database.Database) ([]string, error) {
	if database.Gorm == db {
		return []string{App}, nil
	}
	return []string{}, errors.ErrInitDBCmdInvalid
}

func (c *Command) factoryDatabase(args []string) error {
	switch database.ValueOf(args[0]) {
	case database.Gorm:
		return c.gormInit(args[1])
	default:
		return errors.ErrInitTypeInvalid
	}
}

func (c *Command) gormInit(value string) error {
	switch value {
	case App:
		return c.initServer(database.Gorm)
	default:
		return errors.ErrInitTypeInvalid
	}
}

func (c *Command) Init() {
	c.cmd = &cobra.Command{
		Use:     "init",
		Short:   "Initialize gorm standart-gorm",
		Long:    "Get base of project, api, controller using gorm-crud standart-gorm",
		Example: "go-generator init gorm server",
		Args:    cobra.ExactValidArgs(2),
		RunE:    c.Execute,
	}
}

func (c *Command) initServer(db database.Database) error {
	actualDirectory, err := os.Getwd()
	if err != nil {
		return err
	}
	pathDestiny, err := c.prompt.Ask("Enter the full path of the directory destiny!", actualDirectory)
	if err != nil {
		return errors.ErrDirectoryPathInvalid
	}
	lastChar := pathDestiny[len(pathDestiny)-1:]
	if lastChar == "/" {
		pathDestiny = strings.TrimSuffix(pathDestiny, lastChar)
	}
	moduleName, err := c.prompt.Ask("Enter module of golang project", "github.com/wilian746/go-generator/tmp")
	if err != nil || moduleName == "" {
		return errors.ErrModuleNameInvalid
	}
	return server.NewServer().CreateFoldersAndFiles(pathDestiny, moduleName, db)
}
