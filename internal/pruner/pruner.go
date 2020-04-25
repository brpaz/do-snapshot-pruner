package pruner

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/digitalocean/godo"
)

// Options A struct that encapsulates the configuration options for the prune job.
type Options struct {
	DaysToDelete int
	ResourceType string
}

// Prune Prune the old DigitalOcean snapshots based on the specified options
func Prune(doToken string, opts Options) error {
	client := godo.NewFromToken(doToken)

	snapshotsToDelete, err := findSnapshotsToDelete(client, opts)

	if err != nil {
		return err
	}

	if len(snapshotsToDelete) == 0 {
		log.Println("No snapshots found for deletion")
		return nil
	}

	for _, snapshotID := range snapshotsToDelete {

		err := deleteSnapshot(client, snapshotID)

		if err != nil {
			return fmt.Errorf("Failed to delete snapshot with identifier %s: %s", snapshotID, err.Error())
		}
	}

	return nil
}

// Calls DigitalOcean API and filters the snapshots by the creation date, returing a list of snapshot ids
// that match the specified date interval.
func findSnapshotsToDelete(client *godo.Client, opts Options) ([]string, error) {

	snapshotsToDelete := make([]string, 0)

	snapshots, _, err := client.Snapshots.List(context.TODO(), &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	})

	if err != nil {
		return snapshotsToDelete, err
	}

	maxTime := time.Now().UTC().AddDate(0, 0, -opts.DaysToDelete)

	for _, snapshot := range snapshots {
		snapshotTime, err := time.Parse(time.RFC3339Nano, snapshot.Created)
		if err != nil {
			log.Println("ERROR: Cannot parse snapshot creation date", err)
		}

		// Check if the snapshot is older than the specified date
		if snapshotTime.Before(maxTime) {
			snapshotsToDelete = append(snapshotsToDelete, snapshot.ID)
		}
	}

	return snapshotsToDelete, nil
}

// Delete the snapshot with the specified id
func deleteSnapshot(client *godo.Client, snapshotID string) error {
	log.Printf("Deleting snapshot %s", snapshotID)

	_, err := client.Snapshots.Delete(context.TODO(), snapshotID)

	if err != nil {
		return fmt.Errorf("Failed to delete snapshot with identifier %s: %s", snapshotID, err.Error())
	}

	return nil
}
