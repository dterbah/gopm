package cmd

import (
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [script]",
	Short: "Run a script",
	Long:  "Run a script present in your gopm.json",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//script := args[0]

	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
