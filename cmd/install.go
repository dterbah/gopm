package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [package]",
	Short: "Install a package",
	Long:  "Insstall a specific package for your Go project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, packageName := range args {
			logrus.Info(packageName)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
