package opensearchhandler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/disaster37/opensearch/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var urlRoleMapping = fmt.Sprintf("%s/_plugins/_security/api/rolesmapping/test", baseURL)

func (t *OpensearchHandlerTestSuite) TestRoleMappingGet() {

	result := make(opensearch.SecurityGetRoleMappingResponse)
	roleMapping := &opensearch.SecurityRoleMapping{
		BackendRoles: []string{"ad_group"},
	}
	result["test"] = *roleMapping

	httpmock.RegisterResponder("GET", urlRoleMapping, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, result)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.RoleMappingGet("test")
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), roleMapping, resp)

	// When error
	httpmock.RegisterResponder("GET", urlRoleMapping, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.RoleMappingGet("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestRoleMappingDelete() {

	httpmock.RegisterResponder("DELETE", urlRoleMapping, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.RoleMappingDelete("test")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("DELETE", urlRoleMapping, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.RoleMappingDelete("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestRoleMappingUpdate() {
	roleMapping := &opensearch.SecurityRoleMapping{
		BackendRoles: []string{"ad_group"},
	}

	httpmock.RegisterResponder("PUT", urlRoleMapping, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.RoleMappingUpdate("test", roleMapping)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlRoleMapping, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.RoleMappingUpdate("test", roleMapping)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestRoleMappingDiff() {
	var actual, expected, original *opensearch.SecurityRoleMapping

	expected = &opensearch.SecurityRoleMapping{
		BackendRoles: []string{"ad_group"},
	}

	// When role mapping not exist yet
	actual = nil
	diff, err := t.opensearchHandler.RoleMappingDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When role mapping is the same
	actual = &opensearch.SecurityRoleMapping{
		BackendRoles: []string{"ad_group"},
	}
	diff, err = t.opensearchHandler.RoleMappingDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When role mapping is not the same
	expected.BackendRoles = []string{"ad_group_reader"}
	diff, err = t.opensearchHandler.RoleMappingDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When Elastic add default value
	actual = &opensearch.SecurityRoleMapping{
		BackendRoles: []string{"ad_group"},
		Hidden:       true,
	}

	expected = &opensearch.SecurityRoleMapping{
		BackendRoles: []string{"ad_group"},
	}

	original = &opensearch.SecurityRoleMapping{
		BackendRoles: []string{"ad_group"},
	}

	diff, err = t.opensearchHandler.RoleMappingDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), actual, diff.Patched)

}
