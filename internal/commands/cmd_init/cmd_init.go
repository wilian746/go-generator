package cmdinit

import (
	"github.com/spf13/cobra"
	"github.com/wilian746/gorm-crud-generator/internal/commands/generate/server"
	"github.com/wilian746/gorm-crud-generator/internal/enums/errors"
	"github.com/wilian746/gorm-crud-generator/internal/utils/prompt"
	"os"
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
	if len(args) == 0 {
		return errors.ErrInitTypeEmpty
	}

	if args[0] == "server" {
		return c.initServer()
	}
	return errors.ErrInitTypeInvalid
}

func (c *Command) Init() {
	c.cmd = &cobra.Command{
		Use:       "init",
		Short:     "Initialize gorm standart",
		Long:      "Get base of project, api, controller using gorm-crud standart",
		Example:   "gorm-crud init server",
		ValidArgs: []string{"server"},
		Args:      cobra.ExactValidArgs(1),
		RunE:      c.Execute,
	}
}

func (c *Command) initServer() error {
	actualDirectory, err := os.Getwd()
	if err != nil {
		return err
	}
	pathDestiny, err := c.prompt.Ask("Enter the full path of the directory destiny!", actualDirectory)
	if err != nil {
		return errors.ErrDirectoryPathInvalid
	}
	moduleName, err := c.prompt.Ask("Enter module of golang project", "")
	if err != nil || moduleName == "" {
		return errors.ErrModuleNameInvalid
	}
	return server.NewServer().CreateFoldersAndFiles(pathDestiny, moduleName)
}
