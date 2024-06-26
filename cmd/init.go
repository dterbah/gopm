package cmd

import (
	"fmt"
	"os"

	"github.com/dterbah/gopm/core"
	"github.com/dterbah/gopm/core/engine"
	logger "github.com/dterbah/gopm/log"
	"github.com/dterbah/gopm/utils/input"
	"github.com/dterbah/gopm/utils/user"
	"github.com/spf13/cobra"
)

const DEFAULT_PROJECT_NAME = "go-project"
const DEFAULT_PROJECT_VERSION = "1.0.0"
const DEFAULT_PROJECT_DESCRIPTION = "Description of the project"
const DEFAULT_PROJECT_ENTRY_POINT = "main.go"
const DEFAULT_PROJECT_LICENCE = "MIT"

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Long:  "Initialize a new project with basic information",
	Run: func(cmd *cobra.Command, args []string) {
		var config *core.GoPMConfig = core.NewGoPMConfig()
		config.ProjectName = input.ReadUserInput("Name of your Go project :", DEFAULT_PROJECT_NAME)
		// verify if there is a directory with the project name already existing
		if _, err := os.Stat(config.ProjectName); err == nil {
			logger.Error("Project %s already exists !", config.ProjectName)
			return
		}

		config.Version = input.ReadUserInput("version :", DEFAULT_PROJECT_VERSION)
		config.Description = input.ReadUserInput("description :", DEFAULT_PROJECT_DESCRIPTION)
		config.EntryPoint = input.ReadUserInput("entry point :", DEFAULT_PROJECT_ENTRY_POINT)
		config.Author = input.ReadUserInput("author :", user.GetUsername())
		config.License = input.ReadUserInput("license :", DEFAULT_PROJECT_LICENCE)
		config.Git = input.ReadUserInput("git :", fmt.Sprintf("github.com/%s/%s", config.Author, config.ProjectName))

		err := engine.InitProject(*config)

		if err != nil {
			logger.Error("Error when creating the project %s --> %s", config.ProjectName, err)
		} else {
			logger.Info("Project '%s' created with success !", config.ProjectName)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCommand)
}
