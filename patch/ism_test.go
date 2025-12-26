package patch

import (
	"encoding/json"
	"testing"

	"github.com/disaster37/opensearch/v3"
	"github.com/stretchr/testify/assert"
	"k8s.io/utils/ptr"
)

func TestCleanIsmTemplate(t *testing.T) {
	actual := &opensearch.IsmPutPolicy{
		Policy: opensearch.IsmPolicyBase{
			ID:           ptr.To[string]("policy-test"),
			Description:  ptr.To[string]("ingesting logs"),
			DefaultState: ptr.To[string]("ingest"),
			States: []opensearch.IsmPolicyState{
				{
					Name: "ingest",
					Actions: []map[string]any{
						{
							"retry": map[string]any{
								"count":   3,
								"backoff": "exponential",
								"delay":   "1m",
							},
							"rollover": map[string]any{
								"min_doc_count": float64(5),
								"copy_alias":    false,
							},
						},
						{
							"retry": map[string]any{
								"count":   3,
								"backoff": "exponential",
								"delay":   "1m",
							},
							"shrink": map[string]any{
								"max_shard_size": "50gb",
								"switch_aliases": true,
								"target_index_name_template": map[string]any{
									"source": "shrunk_{{ctx.index}}",
									"lang":   "mustache",
								},
							},
						},
						{
							"retry": map[string]any{
								"count":   3,
								"backoff": "exponential",
								"delay":   "1m",
							},
							"allocation": map[string]any{
								"require": map[string]any{
									"temp": "warm",
								},
								"include":  map[string]any{},
								"exclude":  map[string]any{},
								"wait_for": false,
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
							"retry": map[string]any{
								"count":   3,
								"backoff": "exponential",
								"delay":   "1m",
							},
							"delete": map[string]any{},
						},
					},
				},
			},
			IsmTemplate: []opensearch.IsmPolicyTemplate{
				{
					IndexPatterns:   []string{"test-*"},
					Priority:        ptr.To[int64](0),
					LastUpdatedTime: ptr.To[int64](1709829562552),
				},
			},
		},
	}

	expected := &opensearch.IsmPutPolicy{
		Policy: opensearch.IsmPolicyBase{
			ID:           ptr.To[string]("policy-test"),
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
						{
							"shrink": map[string]any{
								"max_shard_size": "50gb",
								"switch_aliases": true,
								"target_index_name_template": map[string]any{
									"source": "shrunk_{{ctx.index}}",
								},
							},
						},
						{
							"allocation": map[string]any{
								"require": map[string]any{
									"temp": "warm",
								},
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
			IsmTemplate: []opensearch.IsmPolicyTemplate{
				{
					IndexPatterns: []string{"test-*"},
					Priority:      ptr.To[int64](0),
				},
			},
		},
	}

	acualByte, err := json.Marshal(actual)
	if err != nil {
		t.Fatal(err)
	}

	expectedByte, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	acualByte, expectedByte, err = CleanIsmTemplate(acualByte, expectedByte)
	assert.NoError(t, err)
	assert.Equal(t, string(expectedByte), string(acualByte))
}
