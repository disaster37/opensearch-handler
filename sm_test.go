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

var urlSm = fmt.Sprintf("%s/_plugins/_sm/policies/test", baseURL)

func (t *OpensearchHandlerTestSuite) TestSmGet() {

	smResp := &opensearch.SmGetPolicyResponse{
		Id:             "qsqds",
		Version:        1,
		SequenceNumber: 0,
		PrimaryTerm:    1,
		Policy: opensearch.SmPolicy{
			SchemaVersion: ptr.To[int64](2),
			SmPutPolicy: opensearch.SmPutPolicy{
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
			},
		},
	}

	httpmock.RegisterResponder("GET", urlSm, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, smResp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.SmGet("test")
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), smResp, resp)

	// When error
	httpmock.RegisterResponder("GET", urlSm, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.SmGet("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestSmDelete() {

	httpmock.RegisterResponder("DELETE", urlSm, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.SmDelete("test")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("DELETE", urlSm, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.SmDelete("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestSmCreate() {
	smPolicy := &opensearch.SmPutPolicy{
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

	httpmock.RegisterResponder("POST", urlSm, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.SmCreate("test", smPolicy)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("POST", urlSm, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.SmCreate("test", smPolicy)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestSmUpdate() {

	urlSmPut := fmt.Sprintf("%s?if_seq_no=7&if_primary_term=1", urlSm)

	smPolicy := &opensearch.SmPutPolicy{
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

	httpmock.RegisterResponder("PUT", urlSmPut, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.SmUpdate("test", 7, 1, smPolicy)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlSmPut, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.SmUpdate("test", 7, 1, smPolicy)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestSmDiff() {
	var actual, expected, original *opensearch.SmPutPolicy

	expected = &opensearch.SmPutPolicy{
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

	// When SM not exist yet
	actual = nil
	diff, err := t.opensearchHandler.SmDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When SM is the same
	actual = &opensearch.SmPutPolicy{
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
	diff, err = t.opensearchHandler.SmDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When SM is not the same
	expected.Description = ptr.To[string]("test2")
	diff, err = t.opensearchHandler.SmDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When opensearch add default values
	expected = &opensearch.SmPutPolicy{
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

	original = &opensearch.SmPutPolicy{
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

	actual = &opensearch.SmPutPolicy{
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

		Enabled: ptr.To[bool](true),
	}

	diff, err = t.opensearchHandler.SmDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), actual, diff.Patched)

}
