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

var urlActionGroup = fmt.Sprintf("%s/_plugins/_security/api/actiongroups/test", baseURL)

func (t *OpensearchHandlerTestSuite) TestActionGroupGet() {

	result := make(map[string]opensearch.SecurityActionGroup)
	ag := &opensearch.SecurityActionGroup{

		SecurityPutActionGroup: opensearch.SecurityPutActionGroup{
			AllowedActions: []string{"cluster_all"},
		},
	}
	result["test"] = *ag

	httpmock.RegisterResponder("GET", urlActionGroup, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, result)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.ActionGroupGet("test")
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), ag, resp)

	// When error
	httpmock.RegisterResponder("GET", urlActionGroup, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.ActionGroupGet("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestActionGroupDelete() {

	httpmock.RegisterResponder("DELETE", urlActionGroup, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.ActionGroupDelete("test")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("DELETE", urlActionGroup, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.ActionGroupDelete("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestActionGroupUpdate() {
	ag := &opensearch.SecurityPutActionGroup{
		AllowedActions: []string{"cluster_all"},
	}

	httpmock.RegisterResponder("PUT", urlActionGroup, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.ActionGroupUpdate("test", ag)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlActionGroup, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.ActionGroupUpdate("test", ag)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestActionGroupDiff() {
	var actual, expected, original *opensearch.SecurityPutActionGroup

	expected = &opensearch.SecurityPutActionGroup{
		AllowedActions: []string{"cluster_all"},
	}

	// When action group not exist yet
	actual = nil
	diff, err := t.opensearchHandler.ActionGroupDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When action group is the same
	actual = &opensearch.SecurityPutActionGroup{
		AllowedActions: []string{"cluster_all"},
	}
	diff, err = t.opensearchHandler.ActionGroupDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When action group is not the same
	expected.AllowedActions = []string{"cluster_monitor"}

	diff, err = t.opensearchHandler.ActionGroupDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When elastic add default values
	expected = &opensearch.SecurityPutActionGroup{
		AllowedActions: []string{"cluster_all"},
	}

	original = &opensearch.SecurityPutActionGroup{
		AllowedActions: []string{"cluster_all"},
	}

	actual = &opensearch.SecurityPutActionGroup{
		AllowedActions: []string{"cluster_all"},
		Description:    ptr.To[string]("test"),
	}

	diff, err = t.opensearchHandler.ActionGroupDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), actual, diff.Patched)

}
