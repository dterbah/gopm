package cmd

import (
	"github.com/dterbah/gopm/core/engine"
	logger "github.com/dterbah/gopm/log"
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
			logger.Error("%s", err)
		} else {
			logger.Info("✅ Command executed with success")
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
