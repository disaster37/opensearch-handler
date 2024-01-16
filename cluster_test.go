package opensearchhandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/disaster37/opensearch/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var urlCluster = fmt.Sprintf("%s/_cluster", baseURL)

func (t *OpensearchHandlerTestSuite) TestClusterHealth() {

	rawHealth := `
	{
		"cluster_name" : "test",
		"status" : "green",
		"timed_out" : false,
		"number_of_nodes" : 15,
		"number_of_data_nodes" : 10,
		"active_primary_shards" : 166,
		"active_shards" : 340,
		"relocating_shards" : 0,
		"initializing_shards" : 0,
		"unassigned_shards" : 0,
		"delayed_unassigned_shards" : 0,
		"number_of_pending_tasks" : 0,
		"number_of_in_flight_fetch" : 0,
		"task_max_waiting_in_queue_millis" : 0,
		"active_shards_percent_as_number" : 100.0
	}
	`

	healthTest := &opensearch.ClusterHealthResponse{}
	if err := json.Unmarshal([]byte(rawHealth), healthTest); err != nil {
		panic(err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/health", urlCluster), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, rawHealth)
		return resp, nil
	})

	health, err := t.opensearchHandler.ClusterHealth()
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Empty(t.T(), cmp.Diff(healthTest, health))

	// When error
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/health", urlCluster), httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.ClusterHealth()
	assert.Error(t.T(), err)
}
