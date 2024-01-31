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

var urlAudit = fmt.Sprintf("%s/_plugins/_security/api/audit", baseURL)

func (t *OpensearchHandlerTestSuite) TestAuditGet() {

	audit := &opensearch.SecurityGetAuditResponse{
		Config: opensearch.SecurityAudit{
			Enabled: ptr.To[bool](true),
			Audit: opensearch.SecurityAuditSpec{
				IgnoreUsers: []string{"test"},
			},
		},
	}

	httpmock.RegisterResponder("GET", urlAudit, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, audit)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.SecurityAuditGet()
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), audit, resp)

	// When error
	httpmock.RegisterResponder("GET", urlAudit, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.SecurityAuditGet()
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestAuditUpdate() {

	urlAuditUpdate := fmt.Sprintf("%s/config", urlAudit)

	audit := &opensearch.SecurityAudit{
		Enabled: ptr.To[bool](true),
		Audit: opensearch.SecurityAuditSpec{
			IgnoreUsers: []string{"test"},
		},
	}
	httpmock.RegisterResponder("PUT", urlAuditUpdate, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.SecurityAuditUpdate(audit)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlAuditUpdate, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.SecurityAuditUpdate(audit)
	assert.Error(t.T(), err)
}


func (t *OpensearchHandlerTestSuite) TestAuditDiff() {
	var actual, expected, original *opensearch.SecurityAudit

	expected = &opensearch.SecurityAudit{
		Enabled: ptr.To[bool](true),
		Audit: opensearch.SecurityAuditSpec{
			IgnoreUsers: []string{"test"},
		},
	}

	// When audit not exist yet
	actual = nil
	diff, err := t.opensearchHandler.SecurityAuditDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When audit is the same
	actual = &opensearch.SecurityAudit{
		Enabled: ptr.To[bool](true),
		Audit: opensearch.SecurityAuditSpec{
			IgnoreUsers: []string{"test"},
		},
	}
	diff, err = t.opensearchHandler.SecurityAuditDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When audit is not the same
	expected.Audit.IgnoreRequests = []string{"test"}

	diff, err = t.opensearchHandler.SecurityAuditDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When elastic add default values
	expected = &opensearch.SecurityAudit{
		Enabled: ptr.To[bool](true),
		Audit: opensearch.SecurityAuditSpec{
			IgnoreUsers: []string{"test"},
		},
	}

	original = &opensearch.SecurityAudit{
		Enabled: ptr.To[bool](true),
		Audit: opensearch.SecurityAuditSpec{
			IgnoreUsers: []string{"test"},
		},
	}
	actual =&opensearch.SecurityAudit{
		Enabled: ptr.To[bool](true),
		Audit: opensearch.SecurityAuditSpec{
			IgnoreUsers: []string{"test"},
			EnableRest: ptr.To[bool](true),
		},
	}

	diff, err = t.opensearchHandler.SecurityAuditDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), actual, diff.Patched)

}
