package opensearchhandler

import (
	"context"

	"github.com/disaster37/generic-objectmatcher/patch"
	localpatch "github.com/disaster37/opensearch-handler/v2/patch"
	"github.com/disaster37/opensearch/v2"
	jsonIterator "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// SecurityConfigUpdate permit to update the security config
func (h *OpensearchHandlerImpl) SecurityConfigUpdate(config *opensearch.SecurityConfig) (err error) {
	if _, err = h.client.SecurityPutConfig().Body(config).Do(context.Background()); err != nil {
		return errors.Wrap(err, "Error when update security config")
	}
	return nil
}

// SecurityConfigGet permit to get the security config
func (h *OpensearchHandlerImpl) SecurityConfigGet() (config *opensearch.SecurityGetConfigResponse, err error) {
	config, err = h.client.SecurityGetConfig().Do(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "Error when gets ecurity config")
	}
	return config, nil
}

// SecurityConfigDiff permit to diff a security config (it use the 3-way diff)
func (h *OpensearchHandlerImpl) SecurityConfigDiff(actualObject, expectedObject, originalObject *opensearch.SecurityConfig) (patchResult *patch.PatchResult, err error) {
	// If not yet exist
	if actualObject == nil {
		expected, err := jsonIterator.ConfigCompatibleWithStandardLibrary.Marshal(expectedObject)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to convert expected object to byte sequence")
		}

		return &patch.PatchResult{
			Patch:    expected,
			Current:  expected,
			Modified: expected,
			Original: nil,
			Patched:  expectedObject,
		}, nil
	}

	return patch.DefaultPatchMaker.Calculate(actualObject, expectedObject, originalObject, localpatch.RemoveEnvironmentVariableContend)
}
