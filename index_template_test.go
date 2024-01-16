package opensearchhandler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/disaster37/opensearch/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var urlIndexTemplate = fmt.Sprintf("%s/_index_template/test", baseURL)

func (t *OpensearchHandlerTestSuite) TestIndexTemplateGet() {

	result := &opensearch.IndicesGetIndexTemplateResponse{}
	template := &opensearch.IndicesGetIndexTemplate{
		IndexPatterns: []string{"test-index-template"},
		Priority:      2,
		Template: &opensearch.IndicesGetIndexTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
		},
	}
	result.IndexTemplates = opensearch.IndicesGetIndexTemplatesSlice{opensearch.IndicesGetIndexTemplates{IndexTemplate: template}}

	httpmock.RegisterResponder("GET", urlIndexTemplate, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, result)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.IndexTemplateGet("test")
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), template, resp)

	// When error
	httpmock.RegisterResponder("GET", urlIndexTemplate, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.IndexTemplateGet("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestIndexTemplateDelete() {

	httpmock.RegisterResponder("DELETE", urlIndexTemplate, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.IndexTemplateDelete("test")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("DELETE", urlIndexTemplate, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.IndexTemplateDelete("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestIndexTemplateUpdate() {
	template := &opensearch.IndicesGetIndexTemplate{
		IndexPatterns: []string{"test-index-template"},
		Priority:      2,
		Template: &opensearch.IndicesGetIndexTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
		},
	}

	httpmock.RegisterResponder("PUT", urlIndexTemplate, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.IndexTemplateUpdate("test", template)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlIndexTemplate, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.IndexTemplateUpdate("test", template)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestIndexTemplateDiff() {
	var actual, expected, original *opensearch.IndicesGetIndexTemplate

	expected = &opensearch.IndicesGetIndexTemplate{
		IndexPatterns: []string{"test-index-template"},
		Priority:      2,
		Template: &opensearch.IndicesGetIndexTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
		},
	}

	// When template not exist yet
	actual = nil
	diff, err := t.opensearchHandler.IndexTemplateDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When template is the same
	actual = &opensearch.IndicesGetIndexTemplate{
		IndexPatterns: []string{"test-index-template"},
		Priority:      2,
		Template: &opensearch.IndicesGetIndexTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
		},
	}
	diff, err = t.opensearchHandler.IndexTemplateDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When template is not the same
	expected.Template = &opensearch.IndicesGetIndexTemplateData{
		Mappings: map[string]any{
			"_source.enabled":           false,
			"properties.host_name.type": "keyword",
		},
	}
	diff, err = t.opensearchHandler.IndexTemplateDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When Elastic add default value
	actual = &opensearch.IndicesGetIndexTemplate{
		IndexPatterns: []string{"test-index-template"},
		Priority:      2,
		Template: &opensearch.IndicesGetIndexTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
		},
		Meta: map[string]interface{}{
			"default": "test",
		},
	}

	expected = &opensearch.IndicesGetIndexTemplate{
		IndexPatterns: []string{"test-index-template"},
		Priority:      2,
		Template: &opensearch.IndicesGetIndexTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
		},
	}

	original = &opensearch.IndicesGetIndexTemplate{
		IndexPatterns: []string{"test-index-template"},
		Priority:      2,
		Template: &opensearch.IndicesGetIndexTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
		},
	}

	diff, err = t.opensearchHandler.IndexTemplateDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), actual, diff.Patched)

}
