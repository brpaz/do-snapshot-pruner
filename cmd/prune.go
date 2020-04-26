package cmd

import (
	"log"
	"os"

	"github.com/brpaz/do-snapshot-pruner/internal/pruner"
	"github.com/spf13/cobra"
)

// DefaultDaysToDelete Defines the default number of days to delete older snapshot
const DefaultDaysToDelete = 7

var daysToDelete int

var doToken string

var resourceType string

var pruneCmd = &cobra.Command{
	Use:     "prune",
	Short:   "Prunes old DigitalOcean volume snapshots",
	Example: "do-snapshot-pruner -d 5 -t <do_token> -t droplet",
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

func init() {

	flags := pruneCmd.Flags()
	flags.IntVarP(&daysToDelete, "days", "d", DefaultDaysToDelete, "number of days old to consider snapshot for deletion")
	flags.StringVarP(&resourceType, "type", "r", pruner.ResourceTypeAll, "The resource type of snapshots to process (all|volume|droplet)")
	flags.StringVarP(&doToken, "token", "t", os.Getenv("DO_TOKEN"), "DigitalOcean API Token")

	_ = pruneCmd.MarkFlagRequired("token")
}
