package opensearchhandler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/disaster37/opensearch/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var urlRole = fmt.Sprintf("%s/_plugins/_security/api/roles/test", baseURL)

func (t *OpensearchHandlerTestSuite) TestRoleGet() {

	result := make(map[string]opensearch.SecurityRole)
	role := &opensearch.SecurityRole{
		ClusterPermissions: []string{"all"},
		IndexPermissions: []opensearch.SecurityIndexPermissions{
			{
				IndexPatterns:  []string{"logstash-*"},
				AllowedActions: []string{"read"},
			},
		},
	}
	result["test"] = *role

	httpmock.RegisterResponder("GET", urlRole, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, result)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.RoleGet("test")
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), role, resp)

	// When error
	httpmock.RegisterResponder("GET", urlRole, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.RoleGet("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestRoleDelete() {

	httpmock.RegisterResponder("DELETE", urlRole, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.RoleDelete("test")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("DELETE", urlRole, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.RoleDelete("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestRoleUpdate() {
	role := &opensearch.SecurityRole{
		ClusterPermissions: []string{"all"},
		IndexPermissions: []opensearch.SecurityIndexPermissions{
			{
				IndexPatterns:  []string{"logstash-*"},
				AllowedActions: []string{"read"},
			},
		},
	}

	httpmock.RegisterResponder("PUT", urlRole, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.RoleUpdate("test", role)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlRole, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.RoleUpdate("test", role)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestRoleDiff() {
	var actual, expected, original *opensearch.SecurityRole

	expected = &opensearch.SecurityRole{
		ClusterPermissions: []string{"all"},
		IndexPermissions: []opensearch.SecurityIndexPermissions{
			{
				IndexPatterns:  []string{"logstash-*"},
				AllowedActions: []string{"read"},
			},
		},
	}

	// When role not exist yet
	actual = nil
	diff, err := t.opensearchHandler.RoleDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When role is the same
	actual = &opensearch.SecurityRole{
		ClusterPermissions: []string{"all"},
		IndexPermissions: []opensearch.SecurityIndexPermissions{
			{
				IndexPatterns:  []string{"logstash-*"},
				AllowedActions: []string{"read"},
			},
		},
	}
	diff, err = t.opensearchHandler.RoleDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When role is not the same
	expected.IndexPermissions = []opensearch.SecurityIndexPermissions{
		{
			IndexPatterns:  []string{"test-*"},
			AllowedActions: []string{"read"},
		},
	}

	diff, err = t.opensearchHandler.RoleDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When elastic add default values
	expected = &opensearch.SecurityRole{
		ClusterPermissions: []string{"all"},
		IndexPermissions: []opensearch.SecurityIndexPermissions{
			{
				IndexPatterns:  []string{"logstash-*"},
				AllowedActions: []string{"read"},
			},
		},
	}

	original = &opensearch.SecurityRole{
		ClusterPermissions: []string{"all"},
		IndexPermissions: []opensearch.SecurityIndexPermissions{
			{
				IndexPatterns:  []string{"logstash-*"},
				AllowedActions: []string{"read"},
			},
		},
	}

	actual = &opensearch.SecurityRole{
		ClusterPermissions: []string{"all"},
		IndexPermissions: []opensearch.SecurityIndexPermissions{
			{
				IndexPatterns:  []string{"logstash-*"},
				AllowedActions: []string{"read"},
			},
		},
		Reserved: true,
	}

	diff, err = t.opensearchHandler.RoleDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), actual, diff.Patched)

}
