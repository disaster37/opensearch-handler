package opensearchhandler

import (
	"context"

	"emperror.dev/errors"
	"github.com/disaster37/opensearch/v3"
)

// ClusterHealth permit to get the cluster health
func (h *OpensearchHandlerImpl) ClusterHealth() (health *opensearch.ClusterHealthResponse, err error) {
	health, err = h.client.ClusterHealth().Do(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "Error when get cluster health")
	}

	return health, nil
}

// EnableRoutingRebalance permit to enable cluster routing rebalance
// It put `cluster.routing.rebalance.enable` to all
func (h *OpensearchHandlerImpl) EnableRoutingRebalance() (err error) {
	settings := map[string]interface{}{
		"persistent": map[string]interface{}{
			"cluster.routing.rebalance.enable": "all",
		},
	}

	if _, err = h.client.ClusterPutSetting().Body(settings).Do(context.Background()); err != nil {
		return errors.Wrap(err, "Error when enable routing rebalance")
	}

	return nil
}

// DisableRoutingRebalance permit to disable cluster routing rebalance
// It put `cluster.routing.rebalance.enable` to none
func (h *OpensearchHandlerImpl) DisableRoutingRebalance() (err error) {
	settings := map[string]interface{}{
		"persistent": map[string]interface{}{
			"cluster.routing.rebalance.enable": "none",
		},
	}

	if _, err = h.client.ClusterPutSetting().Body(settings).Do(context.Background()); err != nil {
		return errors.Wrap(err, "Error when disable routing rebalance")
	}

	return nil
}

// EnableRoutingAllocation permit to enable cluster routing allocation
// It put `cluster.routing.allocation.enable` to all
func (h *OpensearchHandlerImpl) EnableRoutingAllocation() (err error) {
	settings := map[string]interface{}{
		"persistent": map[string]interface{}{
			"cluster.routing.allocation.enable": "all",
		},
	}

	if _, err = h.client.ClusterPutSetting().Body(settings).Do(context.Background()); err != nil {
		return errors.Wrap(err, "Error when enable routing allocation")
	}

	return nil
}

// DisableRoutingAllocation permit to disable cluster routing allocation
// It put `cluster.routing.allocation.enable` to primaries
func (h *OpensearchHandlerImpl) DisableRoutingAllocation() (err error) {
	settings := map[string]interface{}{
		"persistent": map[string]interface{}{
			"cluster.routing.allocation.enable": "primaries",
		},
	}

	if _, err = h.client.ClusterPutSetting().Body(settings).Do(context.Background()); err != nil {
		return errors.Wrap(err, "Error when disable routing allocation")
	}

	return nil
}
