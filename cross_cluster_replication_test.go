package opensearchhandler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/disaster37/opensearch/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var (
	urlCcr           = fmt.Sprintf("%s/_plugins/_replication/follower-01", baseURL)
	urlCcrAutoFollow = fmt.Sprintf("%s/_plugins/_replication/_autofollow", baseURL)
)

func (t *OpensearchHandlerTestSuite) TestCrossClusterReplicationStatus() {

	url := fmt.Sprintf("%s/_status", urlCcr)

	ccrResp := &opensearch.CcrStatusRuleResponse{
		Status:        "SYNCING",
		Reason:        "User initiated",
		LeaderAlias:   "my-connection-name",
		LeaderIndex:   "leader-01",
		FollowerIndex: "follower-01",
		SyncingDetails: opensearch.CcrRuleSyncingDetails{
			LeaderCheckpoint:   1,
			FollowerCheckpoint: 1,
			SeqNumber:          1,
		},
	}

	httpmock.RegisterResponder("GET", url, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, ccrResp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.CrossClusterReplicationStatus("follower-01")
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), ccrResp, resp)

	// When error
	httpmock.RegisterResponder("GET", url, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.CrossClusterReplicationStatus("follower-01")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestCrossClusterReplicationStart() {
	url := fmt.Sprintf("%s/_start", urlCcr)

	resp := &opensearch.CcrStartRuleResponse{
		Acknowledged: true,
	}

	ccrRule := &opensearch.CcrRule{
		LeaderAlias: "my-connection-name",
		LeaderIndex: "leader-01",
		UseRoles: opensearch.CcrRuleUseRoles{
			LeaderClusterRole:   "all_access",
			FollowerClusterRole: "all_access",
		},
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/_status", urlCcr), func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, nil)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	httpmock.RegisterResponder("PUT", url, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, resp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	err := t.opensearchHandler.CrossClusterReplicationStart("follower-01", ccrRule)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("PUT", url, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.CrossClusterReplicationStart("follower-01", ccrRule)
	assert.Error(t.T(), err)

	// When acknoledge is false
	resp.Acknowledged = false
	httpmock.RegisterResponder("PUT", url, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, resp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	err = t.opensearchHandler.CrossClusterReplicationStart("follower-01", ccrRule)
	assert.Error(t.T(), err)

	// When ccr rule already exist and in pause
	resp.Acknowledged = true
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/_status", urlCcr), func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, opensearch.CcrStatusRuleResponse{
			Status: string(opensearch.CcrStatusPaused),
		})
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/_resume", urlCcr), func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, resp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	err = t.opensearchHandler.CrossClusterReplicationStart("follower-01", ccrRule)
	if err != nil {
		t.Fail(err.Error())
	}

	// When ccr rule already exist and in pause and failed to resume
	resp.Acknowledged = true
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/_status", urlCcr), func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, opensearch.CcrStatusRuleResponse{
			Status: string(opensearch.CcrStatusPaused),
		})
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/_resume", urlCcr), httpmock.NewErrorResponder(errors.New("fack error")))
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/_stop", urlCcr), func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, resp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

}

func (t *OpensearchHandlerTestSuite) TestCrossClusterReplicationStop() {
	url := fmt.Sprintf("%s/_stop", urlCcr)

	resp := &opensearch.CcrStartRuleResponse{
		Acknowledged: true,
	}

	httpmock.RegisterResponder("POST", url, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, resp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	err := t.opensearchHandler.CrossClusterReplicationStop("follower-01")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("POST", url, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.CrossClusterReplicationStop("follower-01")
	assert.Error(t.T(), err)

	// When acknoledge is false
	resp.Acknowledged = false
	httpmock.RegisterResponder("POST", url, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, resp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	err = t.opensearchHandler.CrossClusterReplicationStop("follower-01")
	assert.Error(t.T(), err)

}

func (t *OpensearchHandlerTestSuite) TestCrossClusterReplicationPause() {
	url := fmt.Sprintf("%s/_pause", urlCcr)

	resp := &opensearch.CcrPauseRuleResponse{
		Acknowledged: true,
	}

	httpmock.RegisterResponder("POST", url, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, resp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	err := t.opensearchHandler.CrossClusterReplicationPause("follower-01")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("POST", url, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.CrossClusterReplicationPause("follower-01")
	assert.Error(t.T(), err)

	// When acknoledge is false
	resp.Acknowledged = false
	httpmock.RegisterResponder("POST", url, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, resp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	err = t.opensearchHandler.CrossClusterReplicationPause("follower-01")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestCrossClusterReplicationDiff() {
	var actual, expected *opensearch.CcrRule

	expected = &opensearch.CcrRule{
		LeaderAlias: "my-connection-name",
		LeaderIndex: "leader-01",
		UseRoles: opensearch.CcrRuleUseRoles{
			LeaderClusterRole:   "all_access",
			FollowerClusterRole: "all_access",
		},
	}

	// When CCR not exist yet
	actual = nil
	diff, err := t.opensearchHandler.CrossClusterReplicationDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When CCR is the same
	actual = &opensearch.CcrRule{
		LeaderAlias: "my-connection-name",
		LeaderIndex: "leader-01",
		UseRoles: opensearch.CcrRuleUseRoles{
			LeaderClusterRole:   "all_access",
			FollowerClusterRole: "all_access",
		},
	}
	diff, err = t.opensearchHandler.CrossClusterReplicationDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())

	// When CCR is not the same
	expected.LeaderIndex = "leader-02"
	diff, err = t.opensearchHandler.CrossClusterReplicationDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())

}

func (t *OpensearchHandlerTestSuite) TestCrossClusterReplicationAutoFollowCreate() {

	resp := &opensearch.CcrPostAutoFollowResponse{
		Acknowledged: true,
	}

	ccrRule := &opensearch.CcrAutoFollowRule{
		LeaderAlias: "my-connection-name",
		Name:        "leader-rule",
		Pattern:     "leader-*",
		UseRoles: opensearch.CcrRuleUseRoles{
			LeaderClusterRole:   "all_access",
			FollowerClusterRole: "all_access",
		},
	}

	httpmock.RegisterResponder("POST", urlCcrAutoFollow, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, resp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	err := t.opensearchHandler.CrossClusterReplicationAutoFollowCreate(ccrRule)
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("POST", urlCcrAutoFollow, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.CrossClusterReplicationAutoFollowCreate(ccrRule)
	assert.Error(t.T(), err)

	// When acknoledge is false
	resp.Acknowledged = false
	httpmock.RegisterResponder("POST", urlCcrAutoFollow, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, resp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	err = t.opensearchHandler.CrossClusterReplicationAutoFollowCreate(ccrRule)
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestCrossClusterReplicationAutoFollowStatus() {
	url := fmt.Sprintf("%s/_plugins/_replication/autofollow_stats", baseURL)

	ccrResp := &opensearch.CcrAutoFollowStatusResponse{
		NumSuccessStartReplications: 2,
		NumFailedStartReplications:  0,
		NumFailedLeaderCalls:        0,
		FailedIndices:               []string{},
		AutofollowStats: []opensearch.CcrFollowStatusState{
			{
				Name:                        "leader-rule",
				Pattern:                     "leader-*",
				NumSuccessStartReplications: 2,
				NumFailedStartReplications:  0,
				NumFailedLeaderCalls:        0,
				FailedIndices:               []string{},
				LastExecutionTime: opensearch.UnixMilliTime{
					Time: time.UnixMilli(0),
				},
			},
			{
				Name:                        "log-rule",
				Pattern:                     "log-*",
				NumSuccessStartReplications: 2,
				NumFailedStartReplications:  0,
				NumFailedLeaderCalls:        0,
				FailedIndices:               []string{},
				LastExecutionTime: opensearch.UnixMilliTime{
					Time: time.UnixMilli(0),
				},
			},
		},
	}

	httpmock.RegisterResponder("GET", url, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, ccrResp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})

	resp, err := t.opensearchHandler.CrossClusterReplicationAutoFollowStatus("leader-rule")
	if err != nil {
		t.Fail(err.Error())
	}
	assert.Equal(t.T(), &ccrResp.AutofollowStats[0], resp)

	// When error
	httpmock.RegisterResponder("GET", url, httpmock.NewErrorResponder(errors.New("fack error")))
	_, err = t.opensearchHandler.CrossClusterReplicationAutoFollowStatus("leader-rule")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestCrossClusterReplicationAutoFollowDelete() {

	resp := &opensearch.CcrPauseRuleResponse{
		Acknowledged: true,
	}

	httpmock.RegisterResponder("DELETE", urlCcrAutoFollow, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, resp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	err := t.opensearchHandler.CrossClusterReplicationAutoFollowDelete("leader-rule", "my-connection-name")
	if err != nil {
		t.Fail(err.Error())
	}

	// When error
	httpmock.RegisterResponder("DELETE", urlCcrAutoFollow, httpmock.NewErrorResponder(errors.New("fack error")))
	err = t.opensearchHandler.CrossClusterReplicationAutoFollowDelete("leader-rule", "my-connection-name")
	assert.Error(t.T(), err)

	// When acknoledge is false
	resp.Acknowledged = false
	httpmock.RegisterResponder("DELETE", urlCcrAutoFollow, func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, resp)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	err = t.opensearchHandler.CrossClusterReplicationAutoFollowDelete("leader-rule", "my-connection-name")
	assert.Error(t.T(), err)
}

func (t *OpensearchHandlerTestSuite) TestCrossClusterReplicationAutoFollowDiff() {
	var actual, expected *opensearch.CcrAutoFollowRule

	expected = &opensearch.CcrAutoFollowRule{
		LeaderAlias: "my-connection-name",
		Name:        "leader-rule",
		Pattern:     "leader-*",
		UseRoles: opensearch.CcrRuleUseRoles{
			LeaderClusterRole:   "all_access",
			FollowerClusterRole: "all_access",
		},
	}

	// When CCR not exist yet
	actual = nil
	diff, err := t.opensearchHandler.CrossClusterReplicationAutoFollowDiff(actual, expected, nil)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
	assert.Equal(t.T(), expected, diff.Patched)

	// When CCR is the same
	actual = &opensearch.CcrAutoFollowRule{
		LeaderAlias: "my-connection-name",
		Name:        "leader-rule",
		Pattern:     "leader-*",
		UseRoles: opensearch.CcrRuleUseRoles{
			LeaderClusterRole:   "all_access",
			FollowerClusterRole: "all_access",
		},
	}
	diff, err = t.opensearchHandler.CrossClusterReplicationAutoFollowDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.True(t.T(), diff.IsEmpty())

	// When CCR is not the same
	expected.Pattern = "leader-02-*"
	diff, err = t.opensearchHandler.CrossClusterReplicationAutoFollowDiff(actual, expected, actual)
	if err != nil {
		t.Fail(err.Error())
	}
	assert.False(t.T(), diff.IsEmpty())
}
