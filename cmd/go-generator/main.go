package main

import (
	"fmt"
	"github.com/spf13/cobra"
	cmdHelp "github.com/wilian746/go-generator/internal/commands/help"
	cmdInit "github.com/wilian746/go-generator/internal/commands/init"
	cmdVersion "github.com/wilian746/go-generator/internal/commands/version"
	"github.com/wilian746/go-generator/internal/utils/logger"
	"github.com/wilian746/go-generator/internal/utils/prompt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-generator",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.PRINT("GO Generator is an command line interface to create your API using some databases more facility.")
		logger.PRINT("")
		logger.PRINT("Welcome to the go-generator. Use the `go-generator help` command to view the available commands.")
		return nil
	},
}

// nolint
func init() {
	rootCmd.SetUsageFunc(func(command *cobra.Command) error {
		return nil
	})
	rootCmd.AddCommand(cmdInit.NewInitCommand(prompt.NewPrompt()).Cmd())
	rootCmd.AddCommand(cmdVersion.NewVersionCommand().CmdVersion())
	rootCmd.AddCommand(cmdHelp.NewHelpCommand(rootCmd).UsageHelp())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
