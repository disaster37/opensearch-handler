package samples

import (
	"log"

	"github.com/disaster37/opensearch/v2"
	"k8s.io/utils/ptr"
)

func ManageSnapshotManagement() {

	var (
		snapshotManagement         *opensearch.SmGetPolicyResponse
		expectedSnapshotManagement *opensearch.SmPutPolicy
		originalSnapshotManagement *opensearch.SmPutPolicy
		err                        error
	)

	client := GetClient()

	// Create snapshot management
	snapshotManagementCreate := &opensearch.SmPutPolicy{
		Description: ptr.To[string]("Daily snapshot policy"),
		Creation: opensearch.SmPolicyCreation{
			Schedule: map[string]any{
				"cron": map[string]any{
					"expression": "0 8 * * *",
					"timezone":   "UTC",
				},
			},
			TimeLimit: ptr.To[string]("1h"),
		},
		Deletion: &opensearch.SmPolicyDeletion{
			Schedule: map[string]any{
				"cron": map[string]any{
					"expression": "0 1 * * *",
					"timezone":   "America/Los_Angeles",
				},
			},
			Condition: &opensearch.SmPolicyDeleteCondition{
				MaxAge:   ptr.To[string]("7d"),
				MaxCount: ptr.To[int64](21),
				MinCount: ptr.To[int64](7),
			},
			TimeLimit: ptr.To[string]("1h"),
		},
		SnapshotConfig: opensearch.SmPolicySnapshotConfig{
			DateFormat:         ptr.To[string]("yyyy-MM-dd-HH:mm"),
			Timezone:           ptr.To[string]("America/Los_Angeles"),
			Indices:            ptr.To[string]("*"),
			Repository:         "s3-repo",
			IgnoreUnavailable:  ptr.To[bool](true),
			IncludeGlobalState: ptr.To[bool](false),
			Partial:            ptr.To[bool](true),
			Metadata: map[string]any{
				"any_key": "any_value",
			},
		},
		Notification: &opensearch.SmPolicyNotification{
			Channel: opensearch.SmPolicyNotificationChannel{
				ID: "NC3OpoEBzEoHMX183R3f",
			},
			Conditions: &opensearch.SmPolicyNotificationCondition{
				Creation:          ptr.To[bool](true),
				Deletion:          ptr.To[bool](false),
				Failure:           ptr.To[bool](false),
				TimeLimitExceeded: ptr.To[bool](false),
			},
		},
	}

	if err = client.SmCreate("test", snapshotManagementCreate); err != nil {
		log.Fatalf("Error when create snapshot management: %s", err.Error())
	}

	// Get snapshot management
	snapshotManagement, err = client.SmGet("test")
	if err != nil {
		log.Fatalf("Error when get snapshot management: %s", err.Error())
	}
	log.Print("Get snapshot management successfully\n")

	// Diff snapshotManagement on 3 way merge pattern
	// You need to track somewhere the original snapshotManagement.
	// You need to store them after create or update it
	originalSnapshotManagement = &opensearch.SmPutPolicy{
		Description: ptr.To[string]("Daily snapshot policy"),
		Creation: opensearch.SmPolicyCreation{
			Schedule: map[string]any{
				"cron": map[string]any{
					"expression": "0 8 * * *",
					"timezone":   "UTC",
				},
			},
			TimeLimit: ptr.To[string]("1h"),
		},
		Deletion: &opensearch.SmPolicyDeletion{
			Schedule: map[string]any{
				"cron": map[string]any{
					"expression": "0 1 * * *",
					"timezone":   "America/Los_Angeles",
				},
			},
			Condition: &opensearch.SmPolicyDeleteCondition{
				MaxAge:   ptr.To[string]("7d"),
				MaxCount: ptr.To[int64](21),
				MinCount: ptr.To[int64](7),
			},
			TimeLimit: ptr.To[string]("1h"),
		},
		SnapshotConfig: opensearch.SmPolicySnapshotConfig{
			DateFormat:         ptr.To[string]("yyyy-MM-dd-HH:mm"),
			Timezone:           ptr.To[string]("America/Los_Angeles"),
			Indices:            ptr.To[string]("*"),
			Repository:         "s3-repo",
			IgnoreUnavailable:  ptr.To[bool](true),
			IncludeGlobalState: ptr.To[bool](false),
			Partial:            ptr.To[bool](true),
			Metadata: map[string]any{
				"any_key": "any_value",
			},
		},
		Notification: &opensearch.SmPolicyNotification{
			Channel: opensearch.SmPolicyNotificationChannel{
				ID: "NC3OpoEBzEoHMX183R3f",
			},
			Conditions: &opensearch.SmPolicyNotificationCondition{
				Creation:          ptr.To[bool](true),
				Deletion:          ptr.To[bool](false),
				Failure:           ptr.To[bool](false),
				TimeLimitExceeded: ptr.To[bool](false),
			},
		},
	}

	expectedSnapshotManagement = &opensearch.SmPutPolicy{
		Description: ptr.To[string]("Daily snapshot policy"),
		Creation: opensearch.SmPolicyCreation{
			Schedule: map[string]any{
				"cron": map[string]any{
					"expression": "0 8 * * *",
					"timezone":   "UTC",
				},
			},
			TimeLimit: ptr.To[string]("1h"),
		},
		Deletion: &opensearch.SmPolicyDeletion{
			Condition: &opensearch.SmPolicyDeleteCondition{
				MaxAge:   ptr.To[string]("7d"),
				MaxCount: ptr.To[int64](21),
				MinCount: ptr.To[int64](7),
			},
			TimeLimit: ptr.To[string]("1h"),
		},
		SnapshotConfig: opensearch.SmPolicySnapshotConfig{
			DateFormat:         ptr.To[string]("yyyy-MM-dd-HH:mm"),
			Timezone:           ptr.To[string]("America/Los_Angeles"),
			Indices:            ptr.To[string]("*"),
			Repository:         "s3-repo",
			IgnoreUnavailable:  ptr.To[bool](true),
			IncludeGlobalState: ptr.To[bool](false),
			Partial:            ptr.To[bool](true),
			Metadata: map[string]any{
				"any_key": "any_value",
			},
		},
		Notification: &opensearch.SmPolicyNotification{
			Channel: opensearch.SmPolicyNotificationChannel{
				ID: "NC3OpoEBzEoHMX183R3f",
			},
			Conditions: &opensearch.SmPolicyNotificationCondition{
				Creation:          ptr.To[bool](true),
				Deletion:          ptr.To[bool](false),
				Failure:           ptr.To[bool](false),
				TimeLimitExceeded: ptr.To[bool](false),
			},
		},
	}

	diff, err := client.SmDiff(snapshotManagementCreate, expectedSnapshotManagement, originalSnapshotManagement)
	if err != nil {
		log.Fatalf("Error when diff snapshot management: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())
		// Update snapshot management from diff
		if err = client.SmUpdate("test", snapshotManagement.SequenceNumber, snapshotManagement.PrimaryTerm, diff.Patched.(*opensearch.SmPutPolicy)); err != nil {
			log.Fatalf("Error when update snapshot management: %s", err.Error())
		}
	}

	// Delete snapshot management
	if err = client.SmDelete("test"); err != nil {
		log.Fatalf("Error when delete snapshotManagement: %s", err.Error())
	}

}
