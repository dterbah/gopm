package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Long:  "Initialiaze a new project with basic information",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("INit a new project")
	},
}

func init() {
	rootCmd.AddCommand(initCommand)
}
