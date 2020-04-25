package cmd

import (
	"log"
	"os"

	"github.com/brpaz/do-snapshot-pruner/internal/pruner"
	"github.com/spf13/cobra"
)

const (
	// DefaultDaysToDelete Defines the default number of days to delete older snapshots
	DefaultDaysToDelete = 7
	// ResourceTypeDroplet = "droplet"
	// ResourceTypeVolume  = "volume"
	// ResourceTypeAll     = "all"
)

var daysToDelete int

var doToken string

var resourceType string

var rootCmd = &cobra.Command{
	Use:   "do-snapshot-pruner",
	Short: "Prunes old DigitalOcean volume snapshots",
	Run: func(cmd *cobra.Command, args []string) {
		err := pruner.Prune(doToken, pruner.Options{
			DaysToDelete: daysToDelete,
			ResourceType: resourceType,
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}

// Execute Main command entrypont
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {

	flags := rootCmd.Flags()
	flags.IntVarP(&daysToDelete, "num-days", "n", DefaultDaysToDelete, "number of days old to delete")
	//flags.StringVarP(&resourceType, "resource-type", "r", ResourceTypeAll, "the resource type of snapshots to process (all|volume|droplet)")
	flags.StringVarP(&doToken, "token", "t", os.Getenv("DO_TOKEN"), "DigitalOcean API Token")

	_ = rootCmd.MarkPersistentFlagRequired("token")

	rootCmd.AddCommand(versionCmd)
}
