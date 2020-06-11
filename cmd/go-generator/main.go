package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wilian746/go-generator/internal/commands/help"
	cmdInit "github.com/wilian746/go-generator/internal/commands/init"
	"github.com/wilian746/go-generator/internal/commands/version"
	"github.com/wilian746/go-generator/internal/entities/semver"
	"github.com/wilian746/go-generator/internal/utils/prompt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-generator",
	Short: "GO Generator",
	Long:  "GO Generator is an command line interface to create your API using some databases more facility",
}

// nolint
func init() {
	setVersion()
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

func setVersion() {
	yamlFile, err := ioutil.ReadFile(".semver.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	semverVersion := semver.Semver{}
	err = yaml.Unmarshal(yamlFile, &semverVersion)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	rootCmd.Version = semverVersion.Release
}
