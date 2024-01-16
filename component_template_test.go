package opensearchhandler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/disaster37/opensearch/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var urlComponentTemplate = fmt.Sprintf("%s/_component_template/test", baseURL)

func (t *OpensearchHandlerTestSuite) TestComponentTemplateGet() {

	result := &opensearch.IndicesGetComponentTemplateResponse{}
	component := &opensearch.IndicesGetComponentTemplate{
		Template: &opensearch.IndicesGetComponentTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
			Mappings: map[string]any{
				"_source.enabled":           true,
				"properties.host_name.type": "keyword",
			},
		},
	}

	result.ComponentTemplates = []opensearch.IndicesGetComponentTemplates{{ComponentTemplate: component}}

	httpmock.RegisterResponder("GET", urlComponentTemplate, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, result)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.ComponentTemplateGet("test")
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), component, resp)

	// When error
	httpmock.RegisterResponder("GET", urlComponentTemplate, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.ComponentTemplateGet("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestComponentTemplateDelete() {

	httpmock.RegisterResponder("DELETE", urlComponentTemplate, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.ComponentTemplateDelete("test")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("DELETE", urlComponentTemplate, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.ComponentTemplateDelete("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestComponentTemplateUpdate() {
	component := &opensearch.IndicesGetComponentTemplate{
		Template: &opensearch.IndicesGetComponentTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
			Mappings: map[string]any{
				"_source.enabled":           true,
				"properties.host_name.type": "keyword",
			},
		},
	}

	httpmock.RegisterResponder("PUT", urlComponentTemplate, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.ComponentTemplateUpdate("test", component)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlComponentTemplate, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.ComponentTemplateUpdate("test", component)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestComponentTemplateDiff() {
	var actual, expected, original *opensearch.IndicesGetComponentTemplate

	expected = &opensearch.IndicesGetComponentTemplate{
		Template: &opensearch.IndicesGetComponentTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
			Mappings: map[string]any{
				"_source.enabled":           true,
				"properties.host_name.type": "keyword",
			},
		},
	}

	// When component not exist yet
	actual = nil
	diff, err := t.opensearchHandler.ComponentTemplateDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When component is the same
	actual = &opensearch.IndicesGetComponentTemplate{
		Template: &opensearch.IndicesGetComponentTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
			Mappings: map[string]any{
				"_source.enabled":           true,
				"properties.host_name.type": "keyword",
			},
		},
	}
	diff, err = t.opensearchHandler.ComponentTemplateDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When component is not the same
	expected.Template.Mappings = map[string]any{
		"_source.enabled":           false,
		"properties.host_name.type": "keyword",
	}
	diff, err = t.opensearchHandler.ComponentTemplateDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When elastic add default value
	actual = &opensearch.IndicesGetComponentTemplate{
		Template: &opensearch.IndicesGetComponentTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
			Mappings: map[string]any{
				"_source.enabled":           true,
				"properties.host_name.type": "keyword",
				"default":                   "test",
			},
		},
	}

	expected = &opensearch.IndicesGetComponentTemplate{
		Template: &opensearch.IndicesGetComponentTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
			Mappings: map[string]any{
				"_source.enabled":           true,
				"properties.host_name.type": "keyword",
			},
		},
	}

	original = &opensearch.IndicesGetComponentTemplate{
		Template: &opensearch.IndicesGetComponentTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
			Mappings: map[string]any{
				"_source.enabled":           true,
				"properties.host_name.type": "keyword",
			},
		},
	}

	diff, err = t.opensearchHandler.ComponentTemplateDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), actual, diff.Patched)

}
