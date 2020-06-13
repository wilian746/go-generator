package version

import (
	"github.com/spf13/cobra"
	"github.com/wilian746/go-generator/internal/utils/logger"
)

type IVersion interface {
	CmdVersion() *cobra.Command
}

type Version struct {
}

func NewVersionCommand() IVersion {
	return &Version{}
}

func (v *Version) CmdVersion() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Actual version installed of the Go-Generator",
		Example: "go-generator version",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PRINT(cmd.Short + " is: v0.1.18")
			return nil
		},
	}
}
