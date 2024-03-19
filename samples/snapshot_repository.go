package samples

import (
	"log"

	"github.com/disaster37/opensearch/v2"
)

func ManageSnapshotRepository() {

	var (
		snapshotRepository         *opensearch.SnapshotRepositoryMetaData
		expectedSnapshotRepository *opensearch.SnapshotRepositoryMetaData
		originalSnapshotRepository *opensearch.SnapshotRepositoryMetaData
		err                        error
	)

	client := GetClient()

	// Create snapshot repository
	snapshotRepository = &opensearch.SnapshotRepositoryMetaData{
		Type: "fs",
		Settings: map[string]interface{}{
			"location": "/snapshot",
		},
	}

	if err = client.SnapshotRepositoryUpdate("test", snapshotRepository); err != nil {
		log.Fatalf("Error when create snapshot repository: %s", err.Error())
	}

	// Get snapshot repository
	snapshotRepository, err = client.SnapshotRepositoryGet("test")
	if err != nil {
		log.Fatalf("Error when get snapshot repository: %s", err.Error())
	}
	log.Printf("Get snapshot repository of type %s\n", snapshotRepository.Type)

	// Diff snapshot repository on 3 way merge pattern
	// You need to track somewhere the original snapshot repository.
	// You need to store them after create or update it
	originalSnapshotRepository = &opensearch.SnapshotRepositoryMetaData{
		Type: "fs",
		Settings: map[string]interface{}{
			"location": "/snapshot",
		},
	}

	expectedSnapshotRepository = &opensearch.SnapshotRepositoryMetaData{
		Type: "fs",
		Settings: map[string]interface{}{
			"location": "/snapshot_new",
		},
	}

	diff, err := client.SnapshotRepositoryDiff(snapshotRepository, expectedSnapshotRepository, originalSnapshotRepository)
	if err != nil {
		log.Fatalf("Error when diff snapshot repository: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())
		// Update Snapshot repository from diff
		if err = client.SnapshotRepositoryUpdate("test", diff.Patched.(*opensearch.SnapshotRepositoryMetaData)); err != nil {
			log.Fatalf("Error when update snapshot repository: %s", err.Error())
		}
	}

	// Delete snapshot repository
	if err = client.SnapshotRepositoryDelete("test"); err != nil {
		log.Fatalf("Error when delete snapshot repository: %s", err.Error())
	}

}
