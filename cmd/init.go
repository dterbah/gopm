package cmd

import (
	"fmt"
	"os"

	"github.com/dterbah/gopm/core/config"
	"github.com/dterbah/gopm/core/engine"
	"github.com/dterbah/gopm/utils/input"
	"github.com/dterbah/gopm/utils/user"
	"github.com/sirupsen/logrus"
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
		var config *config.GoPMConfig = config.NewGoPMConfig()
		config.ProjectName = input.ReadUserInput("Name of your Go project :", DEFAULT_PROJECT_NAME)
		if _, err := os.Stat(config.ProjectName); err == nil {
			logrus.Errorf("Project %s already exists !", config.ProjectName)
			return
		}

		// verify if there is a directory with the project name already existing
		config.Version = input.ReadUserInput("version :", DEFAULT_PROJECT_VERSION)
		config.Description = input.ReadUserInput("description :", DEFAULT_PROJECT_DESCRIPTION)
		config.EntryPoint = input.ReadUserInput("entry point :", DEFAULT_PROJECT_ENTRY_POINT)
		config.Author = input.ReadUserInput("author :", user.GetUsername())
		config.License = input.ReadUserInput("license :", DEFAULT_PROJECT_LICENCE)
		config.Git = input.ReadUserInput("git :", fmt.Sprintf("github.com/%s/%s", config.Author, config.ProjectName))

		err := engine.InitProject(*config)

		if err != nil {
			logrus.Errorf("Error when creating the project %s --> %s", config.ProjectName, err)
		} else {
			logrus.Infof("Project '%s' created with success !", config.ProjectName)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCommand)
}
