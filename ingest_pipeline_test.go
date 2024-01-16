package opensearchhandler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/disaster37/opensearch/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var urlIndexIngestPipeline = fmt.Sprintf("%s/_ingest/pipeline/test", baseURL)

func (t *OpensearchHandlerTestSuite) TestIngestPipelineGet() {

	result := opensearch.IngestGetPipelineResponse{}
	pipeline := &opensearch.IngestGetPipeline{
		Description: "test",
	}
	result["test"] = pipeline

	httpmock.RegisterResponder("GET", urlIndexIngestPipeline, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, result)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.IngestPipelineGet("test")
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), pipeline, resp)

	// When error
	httpmock.RegisterResponder("GET", urlIndexIngestPipeline, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.IngestPipelineGet("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestIngestPilelineDelete() {

	httpmock.RegisterResponder("DELETE", urlIndexIngestPipeline, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.IngestPipelineDelete("test")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("DELETE", urlIndexIngestPipeline, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.IngestPipelineDelete("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestIngestPipelineUpdate() {
	pipeline := &opensearch.IngestGetPipeline{
		Description: "test",
	}

	httpmock.RegisterResponder("PUT", urlIndexIngestPipeline, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.IngestPipelineUpdate("test", pipeline)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlIndexIngestPipeline, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.IngestPipelineUpdate("test", pipeline)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestIngestPipelineDiff() {
	var actual, expected, original *opensearch.IngestGetPipeline

	expected = &opensearch.IngestGetPipeline{
		Description: "test",
		Version:     0,
		Processors: []map[string]any{
			{
				"test": "plop",
			},
		},
		OnFailure: []map[string]any{
			{
				"test2": "plop2",
			},
		},
	}

	// When pipeline not exist yet
	actual = nil
	diff, err := t.opensearchHandler.IngestPipelineDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When pipeline is the same
	actual = &opensearch.IngestGetPipeline{
		Description: "test",
		Version:     0,
		Processors: []map[string]any{
			{
				"test": "plop",
			},
		},
		OnFailure: []map[string]any{
			{
				"test2": "plop2",
			},
		},
	}
	diff, err = t.opensearchHandler.IngestPipelineDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When pipeline is not the same
	expected.Processors = []map[string]any{
		{
			"test3": "plop3",
		},
	}
	diff, err = t.opensearchHandler.IngestPipelineDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When Elastic add default value
	actual = &opensearch.IngestGetPipeline{
		Description: "test",
		Version:     10,
		Processors: []map[string]any{
			{
				"test": "plop",
			},
		},
		OnFailure: []map[string]any{
			{
				"test2": "plop2",
			},
		},
	}

	expected = &opensearch.IngestGetPipeline{
		Description: "test",
		Processors: []map[string]any{
			{
				"test": "plop",
			},
		},
		OnFailure: []map[string]any{
			{
				"test2": "plop2",
			},
		},
	}

	original = &opensearch.IngestGetPipeline{
		Description: "test",
		Processors: []map[string]any{
			{
				"test": "plop",
			},
		},
		OnFailure: []map[string]any{
			{
				"test2": "plop2",
			},
		},
	}

	diff, err = t.opensearchHandler.IngestPipelineDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), actual, diff.Patched)

}
