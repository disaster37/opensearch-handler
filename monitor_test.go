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

var urlMonitor = fmt.Sprintf("%s/_plugins/_alerting/monitors/test", baseURL)

func (t *OpensearchHandlerTestSuite) TestMonitorGet() {

	searchUrl := fmt.Sprintf("%s/_plugins/_alerting/monitors/_search", baseURL)

	monitor := &opensearch.AlertingGetMonitor{
		AlertingMonitor: opensearch.AlertingMonitor{
			Type: "monitor",
			Name: "test",
		},
	}
	result := map[string]any{
		"hits": map[string]any{
			"hits": []map[string]any{
				{
					"_source": monitor,
				},
			},
		},
	}

	httpmock.RegisterResponder("GET", searchUrl, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, result)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.MonitorGet("test")
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), monitor, resp)

	// When error
	httpmock.RegisterResponder("GET", searchUrl, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.MonitorGet("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestMonitorDelete() {

	httpmock.RegisterResponder("DELETE", urlMonitor, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.MonitorDelete("test")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("DELETE", urlMonitor, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.MonitorDelete("test")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestMonitorCreate() {

	urlMonitorPost := fmt.Sprintf("%s/_plugins/_alerting/monitors", baseURL)

	monitor := &opensearch.AlertingMonitor{
		Type: "monitor",
		Name: "test",
	}

	httpmock.RegisterResponder("POST", urlMonitorPost, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.MonitorCreate(monitor)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("POST", urlMonitorPost, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.MonitorCreate(monitor)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestMonitorUpdate() {
	urlMonitorPut := fmt.Sprintf("%s?if_seq_no=7&if_primary_term=1", urlMonitor)

	monitor := &opensearch.AlertingMonitor{
		Type: "monitor",
		Name: "test",
	}

	httpmock.RegisterResponder("PUT", urlMonitorPut, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, `{}`)
		return resp, nil
	})

	err := t.opensearchHandler.MonitorUpdate("test", 7, 1, monitor)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", urlMonitorPut, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.MonitorUpdate("test", 7, 1, monitor)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestMonitorDiff() {
	var actual, expected, original *opensearch.AlertingMonitor

	expected = &opensearch.AlertingMonitor{
		Type: "monitor",
		Name: "test",
	}

	// When monitor not exist yet
	actual = nil
	diff, err := t.opensearchHandler.MonitorDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When monitor is the same
	actual = &opensearch.AlertingMonitor{
		Type: "monitor",
		Name: "test",
	}
	diff, err = t.opensearchHandler.MonitorDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When monitor is not the same
	expected.Name = "test2"

	diff, err = t.opensearchHandler.MonitorDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When opensearch add default values
	expected = &opensearch.AlertingMonitor{
		Type: "monitor",
		Name: "test",
	}

	original = &opensearch.AlertingMonitor{
		Type: "monitor",
		Name: "test",
	}

	actual = &opensearch.AlertingMonitor{
		Type:    "monitor",
		Name:    "test",
		Enabled: ptr.To[bool](true),
	}

	diff, err = t.opensearchHandler.MonitorDiff(actual, expected, original)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), actual, diff.Patched)

}
