package help

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/wilian746/go-generator/internal/enums/globals"
	"github.com/wilian746/go-generator/internal/utils/logger"
	"os"
	"strings"
)

type IHelp interface {
	CmdHelp() *cobra.Command
}

type Help struct {
	rootCmd *cobra.Command
	infos   []string
}

func NewHelpCommand(rootCmd *cobra.Command, infos ...string) IHelp {
	version := &Help{
		rootCmd: rootCmd,
		infos:   infos,
	}
	return version
}

func (h *Help) CmdHelp() *cobra.Command {
	return &cobra.Command{
		Use:     "help",
		Short:   "Help about any command",
		Example: "go-generator help",
		RunE: func(cmd *cobra.Command, args []string) error {
			h.printHeader()
			h.printAvailableCommands()
			h.printAvailableFlags()
			h.printAdditionalInformation()
			return nil
		},
	}
}

func (h *Help) printHeader() {
	logHeader := fmt.Sprintf(`
%s
Usage:
	go-generator init [REPOSITORY] [GENERATE_TYPE]

Examples:
	go-generator init gorm app
`, globals.GoGeneratorHeader)

	logger.PRINT(logHeader)
}

func (h *Help) printAvailableCommands() {
	logTable := table.NewWriter()
	logTable.SetOutputMirror(os.Stdout)
	logTable = h.addTableLogToCommands(logTable)
	logger.PRINT("Available Commands:")
	logTable.AppendSeparator()
	logTable.Render()
}

func (h *Help) addTableLogToCommands(logTable table.Writer) table.Writer {
	logTable.AppendHeader(table.Row{"Command", "Example", "Description"})
	availableCommands := []string{}
	for _, command := range h.rootCmd.Commands() {
		availableCommands = append(availableCommands, command.Name())
		if len(availableCommands) == 0 || !h.checkIfExistCommandAvailableInList(availableCommands, command) {
			logTable.AppendRow(table.Row{command.Name(), command.Example, command.Short})
			logTable.AppendSeparator()
		}
	}
	return logTable
}

func (h *Help) checkIfExistCommandAvailableInList(availableCommands []string, command *cobra.Command) bool {
	for _, existingCmd := range availableCommands {
		if existingCmd == command.Name() {
			return true
		}
	}
	return false
}

func (h *Help) printAvailableFlags() {
	flags := h.rootCmd.Flags().FlagUsages()
	allFlags := strings.Split(flags, "\n")
	if len(allFlags) == 1 && strings.TrimSpace(allFlags[0]) == "" {
		return
	}
	logTable := table.NewWriter()
	logTable = h.addTableLogToFlags(logTable, allFlags)
	logTable.SetOutputMirror(os.Stdout)

	logger.PRINT("Available Flags:")
	logTable.AppendSeparator()
	logTable.Render()
}

func (h *Help) addTableLogToFlags(logTable table.Writer, allFlags []string) table.Writer {
	logTable.AppendHeader(table.Row{"Flag"})
	for _, flag := range allFlags {
		flagTrim := strings.TrimSpace(flag)
		if flagTrim != "" {
			logTable.AppendRow(table.Row{flagTrim})
			logTable.AppendSeparator()
		}
	}
	return logTable
}

func (h *Help) printAdditionalInformation() {
	docsURL := "https://github.com/wilian746/go-generator#go-generator"
	docsPrint := fmt.Sprintf(`
Available Docs:
You can access we docs in %s to more information!
`, docsURL)
	for _, info := range h.infos {
		docsPrint += info
	}
	logger.PRINT(docsPrint)
}
