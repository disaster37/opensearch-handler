package opensearchhandler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/disaster37/opensearch/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var urlSnapshotRepository = fmt.Sprintf("%s/_snapshot/test", baseURL)

func (t *OpensearchHandlerTestSuite) TestSnapshotRespositoryGet() {

	snapshotRepository := make(opensearch.SnapshotGetRepositoryResponse)
	snapshotRepository["test"] = &opensearch.SnapshotRepositoryMetaData{
		Type: "fs",
		Settings: map[string]interface{}{
			"location": "/snapshot",
		},
	}

	httpmock.RegisterResponder("GET", urlSnapshotRepository, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, snapshotRepository)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	repo, err := t.opensearchHandler.SnapshotRepositoryGet("test")
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), snapshotRepository["test"], repo)

	// When error
	httpmock.RegisterResponder("GET", urlSnapshotRepository, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.SnapshotRepositoryGet("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestSnapshotRepositoryDelete() {

	httpmock.RegisterResponder("DELETE", urlSnapshotRepository, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.SnapshotRepositoryDelete("test")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("DELETE", urlSnapshotRepository, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.SnapshotRepositoryDelete("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestSnapshotRepositoryUpdate() {

	snapshotRepository := &opensearch.SnapshotRepositoryMetaData{
		Type: "fs",
		Settings: map[string]interface{}{
			"location": "/snapshot",
		},
	}

	httpmock.RegisterResponder("PUT", urlSnapshotRepository, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.SnapshotRepositoryUpdate("test", snapshotRepository)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlSnapshotRepository, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.SnapshotRepositoryUpdate("test", snapshotRepository)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestSnapshotRepositoryDiff() {
	var actual, expected, original *opensearch.SnapshotRepositoryMetaData

	expected = &opensearch.SnapshotRepositoryMetaData{
		Type: "fs",
		Settings: map[string]interface{}{
			"location": "/snapshot",
		},
	}

	// When SLM not exist yet
	actual = nil
	diff, err := t.opensearchHandler.SnapshotRepositoryDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When policy is the same
	actual = &opensearch.SnapshotRepositoryMetaData{
		Type: "fs",
		Settings: map[string]interface{}{
			"location": "/snapshot",
		},
	}
	diff, err = t.opensearchHandler.SnapshotRepositoryDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When policy is not the same
	expected.Type = "s3"
	diff, err = t.opensearchHandler.SnapshotRepositoryDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When elastic set default values
	actual = &opensearch.SnapshotRepositoryMetaData{
		Type: "fs",
		Settings: map[string]interface{}{
			"location": "/snapshot",
			"default":  "plop",
		},
	}
	expected = &opensearch.SnapshotRepositoryMetaData{
		Type: "fs",
		Settings: map[string]interface{}{
			"location": "/snapshot",
		},
	}
	original = &opensearch.SnapshotRepositoryMetaData{
		Type: "fs",
		Settings: map[string]interface{}{
			"location": "/snapshot",
		},
	}

	diff, err = t.opensearchHandler.SnapshotRepositoryDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), actual, diff.Patched)

}
