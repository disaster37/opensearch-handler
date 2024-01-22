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

var urlUser = fmt.Sprintf("%s/_plugins/_security/api/internalusers/test", baseURL)

func (t *OpensearchHandlerTestSuite) TestUserGet() {

	result := make(opensearch.SecurityGetUserResponse)
	user := &opensearch.SecurityUser{
		SecurityUserBase: opensearch.SecurityUserBase{
			SecurityRoles: []string{"kibana_user"},
		},
	}
	result["test"] = *user

	httpmock.RegisterResponder("GET", urlUser, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, result)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.UserGet("test")
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), user, resp)

	// When error
	httpmock.RegisterResponder("GET", urlUser, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.UserGet("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestUserDelete() {

	httpmock.RegisterResponder("DELETE", urlUser, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.UserDelete("test")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("DELETE", urlUser, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.UserDelete("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestUserUpdate() {
	user := &opensearch.SecurityPutUser{
		SecurityUserBase: opensearch.SecurityUserBase{
			SecurityRoles: []string{"kibana_user"},
		},
		Password: "password",
	}

	httpmock.RegisterResponder("PUT", urlUser, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.UserUpdate("test", user)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlUser, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.UserUpdate("test", user)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestUserDiff() {
	var actual, expected, original *opensearch.SecurityPutUser

	expected = &opensearch.SecurityPutUser{

		SecurityUserBase: opensearch.SecurityUserBase{
			SecurityRoles: []string{"kibana_user"},
		},
		Password: "password",
	}

	// When user not exist yet
	actual = nil
	diff, err := t.opensearchHandler.UserDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When user is the same
	actual = &opensearch.SecurityPutUser{
		SecurityUserBase: opensearch.SecurityUserBase{
			SecurityRoles: []string{"kibana_user"},
		},
		Password: "password",
	}
	diff, err = t.opensearchHandler.UserDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When user is not the same
	expected.SecurityRoles = []string{"fake"}
	diff, err = t.opensearchHandler.UserDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When Elastic add default value
	actual = &opensearch.SecurityPutUser{
		SecurityUserBase: opensearch.SecurityUserBase{
			SecurityRoles: []string{"kibana_user"},
			Description:   ptr.To[string]("test"),
		},
		Password: "password",
	}

	expected = &opensearch.SecurityPutUser{
		SecurityUserBase: opensearch.SecurityUserBase{
			SecurityRoles: []string{"kibana_user"},
		},
		Password: "password",
	}

	original = &opensearch.SecurityPutUser{
		SecurityUserBase: opensearch.SecurityUserBase{
			SecurityRoles: []string{"kibana_user"},
		},
		Password: "password",
	}

	diff, err = t.opensearchHandler.UserDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), actual, diff.Patched)

}
