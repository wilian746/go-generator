package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wilian746/gorm-crud-generator/internal/commands/cmd_init"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-generator",
	Short: "GO Generator",
	Long:  "GO Generator is an command line interface to create your API using some databases more facility",
}

// nolint
func init() {
	rootCmd.AddCommand(cmdinit.NewInitCommand().Cmd())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
