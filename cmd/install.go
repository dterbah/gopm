package cmd

import (
	"github.com/dterbah/gopm/core/engine"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [packages]",
	Short: "Install packages",
	Long:  "Insstall specific packages for your Go project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := engine.AddDependencies(args)
		if err != nil {
			logrus.Errorf("Error when installing dependencies [%s]", err)
		} else {
			logrus.Infof("⌛️ %d dependecies added to your project", len(args))
		}
	},
}

var iCmd = &cobra.Command{
	Use:   "i [packages]",
	Short: "Install packages",
	Long:  "Insstall specific packages for your Go project",
	Args:  cobra.MinimumNArgs(1),
	Run:   installCmd.Run,
}

func init() {
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(iCmd)
}
