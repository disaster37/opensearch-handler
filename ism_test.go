package opensearchhandler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/disaster37/opensearch/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"k8s.io/utils/ptr"
)

var urlIsm = fmt.Sprintf("%s/_plugins/_ism/policies/test", baseURL)

func (t *OpensearchHandlerTestSuite) TestIsmGet() {

	ismResp := &opensearch.IsmGetPolicyResponse{
		Id:             "qsqds",
		Version:        1,
		SequenceNumber: 0,
		PrimaryTerm:    1,
		Policy: opensearch.IsmGetPolicy{
			SchemaVersion: ptr.To[int64](2),
			IsmPolicyBase: opensearch.IsmPolicyBase{
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
		},
	}

	httpmock.RegisterResponder("GET", urlIsm, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, ismResp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.IsmGet("test")
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), ismResp, resp)

	// When error
	httpmock.RegisterResponder("GET", urlIsm, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.IsmGet("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestIsmDelete() {

	httpmock.RegisterResponder("DELETE", urlIsm, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.IsmDelete("test")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("DELETE", urlIsm, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.IsmDelete("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestIsmCreate() {
	ismPolicy := &opensearch.IsmPutPolicy{
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

	httpmock.RegisterResponder("PUT", urlIsm, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.IsmCreate("test", ismPolicy)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlIsm, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.IsmCreate("test", ismPolicy)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestIsmUpdate() {

	urlIsmPut := fmt.Sprintf("%s?if_seq_no=7&if_primary_term=1", urlIsm)

	ismPolicy := &opensearch.IsmPutPolicy{
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

	httpmock.RegisterResponder("PUT", urlIsmPut, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.IsmUpdate("test", 7, 1, ismPolicy)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlIsmPut, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.IsmUpdate("test", 7, 1, ismPolicy)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestIsmDiff() {
	var actual, expected, original *opensearch.IsmPutPolicy

	expected = &opensearch.IsmPutPolicy{
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

	// When ISM not exist yet
	actual = nil
	diff, err := t.opensearchHandler.IsmDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When ISM is the same
	actual = &opensearch.IsmPutPolicy{
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
	diff, err = t.opensearchHandler.IsmDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When ISM is not the same
	expected.Policy.Description = ptr.To[string]("test2")
	diff, err = t.opensearchHandler.IsmDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When opensearch add default values
	expected = &opensearch.IsmPutPolicy{
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

	original = &opensearch.IsmPutPolicy{
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

	actual = &opensearch.IsmPutPolicy{
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
			ErrorNotification: &opensearch.IsmErrorNotification{
				Channel: &opensearch.IsmErrorNotificationChannel{
					ID: "test",
				},
			},
		},
	}

	diff, err = t.opensearchHandler.IsmDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), actual, diff.Patched)

	// Check fix real issue on opensearch operator
	expected = &opensearch.IsmPutPolicy{
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

	original = &opensearch.IsmPutPolicy{
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

	actual = &opensearch.IsmPutPolicy{
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

	diff, err = t.opensearchHandler.IsmDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), `{}`, string(diff.Patch))

}
