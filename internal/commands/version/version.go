package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wilian746/go-generator/internal/utils/logger"
)

type Interface interface {
	Cmd() *cobra.Command
	Execute(_ *cobra.Command, _ []string) error
}

type Version struct {
	cmd     *cobra.Command
	rootCmd *cobra.Command
}

func NewVersionCommand(rootCmd *cobra.Command) Interface {
	version := &Version{
		rootCmd: rootCmd,
	}
	version.Version()
	return version
}

func (v *Version) Execute(_ *cobra.Command, _ []string) error {
	logger.PRINT(fmt.Sprintf("Actual version installed of the Go-Generator is %s", v.rootCmd.Version))
	return nil
}

func (v *Version) Cmd() *cobra.Command {
	return v.cmd
}

func (v *Version) Version() {
	v.cmd = &cobra.Command{
		Use:     "version",
		Short:   "Get actual version of the go-generator",
		Example: "go-generator version",
		RunE:    v.Execute,
	}
}
