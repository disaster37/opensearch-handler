package samples

import (
	"log"

	"github.com/disaster37/opensearch/v2"
	"k8s.io/utils/ptr"
)

func ManageIndexStateManagement() {

	var (
		ism         *opensearch.IsmGetPolicyResponse
		expectedIsm *opensearch.IsmPutPolicy
		originalIsm *opensearch.IsmPutPolicy
		err         error
	)

	client := GetClient()

	// Create ism
	ismCreate := &opensearch.IsmPutPolicy{
		Policy: opensearch.IsmPolicyBase{
			Description:  ptr.To[string]("ingesting logs"),
			DefaultState: ptr.To[string]("ingest"),
			States: []opensearch.IsmPolicyState{
				{
					Name: "ingest",
					Actions: []map[string]any{
						{
							"rollover": map[string]any{
								"min_doc_count": float64(5),
							},
						},
					},
					Transitions: []opensearch.IsmPolicyStateTransition{
						{
							StateName: "search",
						},
					},
				},
				{
					Name: "search",
					Transitions: []opensearch.IsmPolicyStateTransition{
						{
							StateName: "delete",
							Conditions: map[string]any{
								"min_index_age": "5m",
							},
						},
					},
				},
				{
					Name: "delete",
					Actions: []map[string]any{
						{
							"delete": map[string]any{},
						},
					},
				},
			},
		},
	}

	if err = client.IsmCreate("test", ismCreate); err != nil {
		log.Fatalf("Error when create ism: %s", err.Error())
	}

	// Get ism
	ism, err = client.IsmGet("test")
	if err != nil {
		log.Fatalf("Error when get ism: %s", err.Error())
	}
	log.Print("Get ism successfully\n")

	// Diff ism on 3 way merge pattern
	// You need to track somewhere the original ism.
	// You need to store them after create or update it
	originalIsm = &opensearch.IsmPutPolicy{
		Policy: opensearch.IsmPolicyBase{
			Description:  ptr.To[string]("ingesting logs"),
			DefaultState: ptr.To[string]("ingest"),
			States: []opensearch.IsmPolicyState{
				{
					Name: "ingest",
					Actions: []map[string]any{
						{
							"rollover": map[string]any{
								"min_doc_count": float64(5),
							},
						},
					},
					Transitions: []opensearch.IsmPolicyStateTransition{
						{
							StateName: "search",
						},
					},
				},
				{
					Name: "search",
					Transitions: []opensearch.IsmPolicyStateTransition{
						{
							StateName: "delete",
							Conditions: map[string]any{
								"min_index_age": "5m",
							},
						},
					},
				},
				{
					Name: "delete",
					Actions: []map[string]any{
						{
							"delete": map[string]any{},
						},
					},
				},
			},
		},
	}

	expectedIsm = &opensearch.IsmPutPolicy{
		Policy: opensearch.IsmPolicyBase{
			Description:  ptr.To[string]("ingesting logs"),
			DefaultState: ptr.To[string]("ingest"),
			States: []opensearch.IsmPolicyState{
				{
					Name: "ingest",
					Actions: []map[string]any{
						{
							"rollover": map[string]any{
								"min_doc_count": float64(5),
							},
						},
					},
					Transitions: []opensearch.IsmPolicyStateTransition{
						{
							StateName: "search",
						},
					},
				},
				{
					Name: "search",
					Transitions: []opensearch.IsmPolicyStateTransition{
						{
							StateName: "delete",
							Conditions: map[string]any{
								"min_index_age": "10d",
							},
						},
					},
				},
				{
					Name: "delete",
					Actions: []map[string]any{
						{
							"delete": map[string]any{},
						},
					},
				},
			},
		},
	}

	diff, err := client.IsmDiff(ismCreate, expectedIsm, originalIsm)
	if err != nil {
		log.Fatalf("Error when diff ism: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())
		// Update ism from diff
		if err = client.IsmUpdate("test", ism.SequenceNumber, ism.PrimaryTerm, diff.Patched.(*opensearch.IsmPutPolicy)); err != nil {
			log.Fatalf("Error when update ism: %s", err.Error())
		}
	}

	// Delete ism
	if err = client.IsmDelete("test"); err != nil {
		log.Fatalf("Error when delete ism: %s", err.Error())
	}

}
