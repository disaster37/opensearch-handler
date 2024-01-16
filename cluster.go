package opensearchhandler

import (
	"context"

	"github.com/disaster37/opensearch/v2"
	"github.com/pkg/errors"
)

// ClusterHealth permit to get the cluster health
func (h *OpensearchHandlerImpl) ClusterHealth() (health *opensearch.ClusterHealthResponse, err error) {

	health, err = h.client.ClusterHealth().Do(context.Background())

	if err != nil {
		return nil, errors.Wrap(err, "Error when get cluster health")
	}

	return health, nil
}
