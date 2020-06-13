package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wilian746/go-generator/internal/commands/help"
	cmdInit "github.com/wilian746/go-generator/internal/commands/init"
	"github.com/wilian746/go-generator/internal/commands/version"
	"github.com/wilian746/go-generator/internal/utils/prompt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-generator",
	Short: "GO Generator",
	Long:  "GO Generator is an command line interface to create your API using some databases more facility",
}

// nolint
func init() {
	rootCmd.AddCommand(cmdInit.NewInitCommand(prompt.NewPrompt()).Cmd())
	rootCmd.AddCommand(version.NewVersionCommand(rootCmd).Cmd())
	rootCmd.SetHelpCommand(help.NewHelpCommand(rootCmd).Cmd())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
