package cmd

import (
	"github.com/dterbah/gopm/core/engine"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [script]",
	Short: "Run a script",
	Long:  "Run a script present in your gopm.json",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		script := args[0]
		err := engine.RunScript(script)

		if err != nil {
			logrus.Errorf("Error when launching the command %s --> %s", script, err)
		} else {
			logrus.Info("✅ Command executed with success")
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
