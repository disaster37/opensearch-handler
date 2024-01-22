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

var urlTenant = fmt.Sprintf("%s/_plugins/_security/api/tenants/test", baseURL)

func (t *OpensearchHandlerTestSuite) TestTenantGet() {

	result := make(map[string]opensearch.SecurityTenant)
	tenant := &opensearch.SecurityTenant{
		SecurityPutTenant: opensearch.SecurityPutTenant{
			Description: ptr.To[string]("test"),
		},
	}
	result["test"] = *tenant

	httpmock.RegisterResponder("GET", urlTenant, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, result)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.TenantGet("test")
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), tenant, resp)

	// When error
	httpmock.RegisterResponder("GET", urlTenant, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.TenantGet("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestTenantDelete() {

	httpmock.RegisterResponder("DELETE", urlTenant, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.TenantDelete("test")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("DELETE", urlTenant, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.TenantDelete("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestTenantUpdate() {
	tenant := &opensearch.SecurityPutTenant{
		Description: ptr.To[string]("test"),
	}

	httpmock.RegisterResponder("PUT", urlTenant, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.TenantUpdate("test", tenant)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlTenant, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.TenantUpdate("test", tenant)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestTenantDiff() {
	var actual, expected, original *opensearch.SecurityPutTenant

	expected = &opensearch.SecurityPutTenant{
		Description: ptr.To[string]("test"),
	}

	// When tenant not exist yet
	actual = nil
	diff, err := t.opensearchHandler.TenantDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When tenant is the same
	actual = &opensearch.SecurityPutTenant{
		Description: ptr.To[string]("test"),
	}
	diff, err = t.opensearchHandler.TenantDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When tenant is not the same
	expected.Description = ptr.To[string]("test2")

	diff, err = t.opensearchHandler.TenantDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When elastic add default values
	expected = &opensearch.SecurityPutTenant{}

	original = &opensearch.SecurityPutTenant{}

	actual = &opensearch.SecurityPutTenant{
		Description: ptr.To[string]("test"),
	}

	diff, err = t.opensearchHandler.TenantDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), actual, diff.Patched)

}
