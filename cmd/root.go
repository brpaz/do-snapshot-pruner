package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "do-snapshot-pruner",
	Short: "Prunes old DigitalOcean volume snapshots",
}

// Execute Main command entrypont
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(pruneCmd)
	rootCmd.AddCommand(versionCmd)
}
