package help

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/wilian746/go-generator/internal/enums/globals"
	"github.com/wilian746/go-generator/internal/utils/logger"
	"os"
)

type Interface interface {
	Cmd() *cobra.Command
	PrintHelp(examples string, additionalInfo string)
	Execute(_ *cobra.Command, _ []string) error
}

type Help struct {
	rootCmd *cobra.Command
	cmd     *cobra.Command
}

func NewHelpCommand(rootCmd *cobra.Command) Interface {
	version := &Help{
		rootCmd: rootCmd,
	}
	version.Help()
	return version
}

func (h *Help) Cmd() *cobra.Command {
	return h.cmd
}

func (h *Help) Help() {
	h.cmd = &cobra.Command{
		Use:     "help",
		Short:   "help about how usage any command",
		Example: "go-generator version",
		RunE:    h.Execute,
	}
}

func (h *Help) Execute(_ *cobra.Command, _ []string) error {
	h.PrintHelp("go-generator init gorm app", "")
	return nil
}

func (h *Help) PrintHelp(examples, additionalInfo string) {
	h.printHeader(examples)
	h.printAvailableCommands()
	h.printAdditionalInformation(additionalInfo)
}

func (h *Help) printHeader(examples string) {
	logHeader := fmt.Sprintf(`
%s
Usage:
	go-generator init [REPOSITORY] [GENERATE_TYPE]

Examples:
	%s

Available Commands:
`, globals.GoGeneratorHeader, examples)

	logger.PRINT(logHeader)
}

func (h *Help) printAvailableCommands() {
	logTable := table.NewWriter()
	logTable.SetOutputMirror(os.Stdout)
	logTable.AppendHeader(table.Row{"Command", "Short Description"})
	for _, command := range h.rootCmd.Commands() {
		logTable.AppendRow(table.Row{command.Name(), command.Short})
		logTable.AppendSeparator()
	}
	logTable.AppendSeparator()
	logTable.Render()
}

func (h *Help) printAdditionalInformation(info string) {
	docsURL := "https://github.com/wilian746/go-generator#go-generator"
	docsPrint := fmt.Sprintf(`
Available Docs:
You can access we docs in %s to more information!
%s
`, docsURL, info)
	logger.PRINT(docsPrint)
}
