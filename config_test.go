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

var urlConfig = fmt.Sprintf("%s/_plugins/_security/api/securityconfig", baseURL)

func (t *OpensearchHandlerTestSuite) TestConfigGet() {

	config := &opensearch.SecurityGetConfigResponse{
		Config: opensearch.SecurityConfig{
			Dynamic: opensearch.SecurityConfigDynamic{
				DoNotFailOnForbidden:      ptr.To[bool](true),
				DoNotFailOnForbiddenEmpty: ptr.To[bool](true),
			},
		},
	}

	httpmock.RegisterResponder("GET", urlConfig, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, config)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.ConfigGet()
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), config, resp)

	// When error
	httpmock.RegisterResponder("GET", urlConfig, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.ConfigGet()
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestConfigUpdate() {

	urlConfigUpdate := fmt.Sprintf("%s/config", urlConfig)

	config := &opensearch.SecurityConfig{
		Dynamic: opensearch.SecurityConfigDynamic{
			DoNotFailOnForbidden:      ptr.To[bool](true),
			DoNotFailOnForbiddenEmpty: ptr.To[bool](true),
		},
	}

	httpmock.RegisterResponder("PUT", urlConfigUpdate, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.ConfigUpdate(config)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlConfigUpdate, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.ConfigUpdate(config)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestConfigDiff() {
	var actual, expected, original *opensearch.SecurityConfig

	expected = &opensearch.SecurityConfig{
		Dynamic: opensearch.SecurityConfigDynamic{
			DoNotFailOnForbidden:      ptr.To[bool](true),
			DoNotFailOnForbiddenEmpty: ptr.To[bool](true),
		},
	}

	// When config not exist yet
	actual = nil
	diff, err := t.opensearchHandler.ConfigDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When config is the same
	actual = &opensearch.SecurityConfig{
		Dynamic: opensearch.SecurityConfigDynamic{
			DoNotFailOnForbidden:      ptr.To[bool](true),
			DoNotFailOnForbiddenEmpty: ptr.To[bool](true),
		},
	}
	diff, err = t.opensearchHandler.ConfigDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When config is not the same
	expected.Dynamic.Http = &opensearch.SecurityConfigHttp{
		AnonymousAuthEnabled: ptr.To[bool](true),
	}

	diff, err = t.opensearchHandler.ConfigDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When elastic add default values
	expected = &opensearch.SecurityConfig{
		Dynamic: opensearch.SecurityConfigDynamic{
			DoNotFailOnForbidden:      ptr.To[bool](true),
			DoNotFailOnForbiddenEmpty: ptr.To[bool](true),
		},
	}

	original = &opensearch.SecurityConfig{
		Dynamic: opensearch.SecurityConfigDynamic{
			DoNotFailOnForbidden:      ptr.To[bool](true),
			DoNotFailOnForbiddenEmpty: ptr.To[bool](true),
		},
	}

	actual = &opensearch.SecurityConfig{
		Dynamic: opensearch.SecurityConfigDynamic{
			DoNotFailOnForbidden:      ptr.To[bool](true),
			DoNotFailOnForbiddenEmpty: ptr.To[bool](true),
			License:                   ptr.To[string]("test"),
		},
	}

	diff, err = t.opensearchHandler.ConfigDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), actual, diff.Patched)

}
