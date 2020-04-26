package pruner_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
	"time"

	"github.com/brpaz/do-snapshot-pruner/internal/pruner"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

const testToken = "some-token"

const apiEndpoint = "https://api.digitalocean.com/v2"

type SnapshotsList struct {
	Snapshots []struct {
		ID           string    `json:"id"`
		Name         string    `json:"name"`
		CreatedAt    time.Time `json:"created_at"`
		ResourceType string    `json:"resource_type"`
	} `json:"snapshots"`
}

func loadMockResponseFile(filename string) (SnapshotsList, error) {
	var data SnapshotsList

	cwd, _ := os.Getwd()
	filePath := path.Join(cwd, "..", "..", "testdata", filename)

	mockFile, err := os.Open(filePath)

	if err != nil {
		return data, err
	}

	defer mockFile.Close()

	byteValue, _ := ioutil.ReadAll(mockFile)

	err = json.Unmarshal([]byte(byteValue), &data)

	if err != nil {
		return data, err
	}

	return data, nil
}

func TestDoNothingIfNoSnapshotsFound(t *testing.T) {

	data, err := loadMockResponseFile("snapshots_list_empty_response.json")

	if err != nil {
		t.Fatal("Cannot open response file snapshots_list_empty_response.json", err)
	}

	defer gock.Off() // Flush pending mocks after test execution
	defer gock.DisableNetworking()

	gock.New(apiEndpoint).
		Get("/snapshots").
		Reply(200).
		JSON(data)

	err = pruner.Prune(testToken, pruner.Options{
		DaysToDelete: 5,
		ResourceType: "all",
	})

	assert.Nil(t, err)
}

func TestWithSnapshotsToDelete(t *testing.T) {

	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)

	data, err := loadMockResponseFile("snapshots_list_response.json")

	data.Snapshots[0].CreatedAt = time.Now().AddDate(0, 0, -5)
	data.Snapshots[1].CreatedAt = time.Now().AddDate(0, 0, -10)
	data.Snapshots[2].CreatedAt = time.Now().AddDate(0, 0, -2)

	if err != nil {
		t.Fatal("Cannot open response mock file", err)
	}

	defer gock.Off()
	defer gock.DisableNetworking()

	gock.New(apiEndpoint).
		Get("/snapshots").
		Reply(200).
		JSON(data)

	gock.New(apiEndpoint).
		Delete("/v2/snapshots/1").
		Reply(204)

	gock.New(apiEndpoint).
		Delete("/v2/snapshots/2").
		Reply(204)

	err = pruner.Prune(testToken, pruner.Options{
		DaysToDelete: 5,
		ResourceType: "all",
	})

	assert.Nil(t, err)

	logStr := logOutput.String()
	assert.Contains(t, logStr, "Deleting snapshot 1")
	assert.Contains(t, logStr, "Deleting snapshot 2")
	assert.NotContains(t, logStr, "Deleting snapshot 3")
}

func TestWithSnapshotsToDeleteFilteredByResourceType(t *testing.T) {

	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)

	data, err := loadMockResponseFile("snapshots_list_response.json")

	data.Snapshots[0].CreatedAt = time.Now().AddDate(0, 0, -5)
	data.Snapshots[1].CreatedAt = time.Now().AddDate(0, 0, -10)
	data.Snapshots[2].CreatedAt = time.Now().AddDate(0, 0, -2)

	if err != nil {
		t.Fatal("Cannot open response mock file", err)
	}

	defer gock.Off()
	defer gock.DisableNetworking()

	gock.New(apiEndpoint).
		Get("/snapshots").
		Reply(200).
		JSON(data)

	gock.New(apiEndpoint).
		Delete("/v2/snapshots/2").
		Reply(204)

	err = pruner.Prune(testToken, pruner.Options{
		DaysToDelete: 5,
		ResourceType: pruner.ResourceTypeVolume,
	})

	assert.Nil(t, err)

	logStr := logOutput.String()
	assert.Contains(t, logStr, "Deleting snapshot 2")
}
