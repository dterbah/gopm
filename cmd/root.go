package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	version     string
	showVersion bool
)

var rootCmd = &cobra.Command{
	Use:   "gopm",
	Short: "Gopm is a package manager for Go",
	Long:  `Gopm is a CLI tool to manage Go project dependencies.`,
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Printf("Version : %s\n", version)
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func init() {
	// init logger
	// Configuration de Logrus
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		ForceColors:            true,
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})
	logrus.SetLevel(logrus.InfoLevel)

	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Show version")
	version = "1.0.1"
}
