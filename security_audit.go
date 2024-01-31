package opensearchhandler

import (
	"context"

	"github.com/disaster37/generic-objectmatcher/patch"
	localpatch "github.com/disaster37/opensearch-handler/v2/patch"
	"github.com/disaster37/opensearch/v2"
	jsonIterator "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// SecurityAuditUpdate permit to update the security audit
func (h *OpensearchHandlerImpl) SecurityAuditUpdate(audit *opensearch.SecurityAudit) (err error) {
	if _, err = h.client.SecurityPutAudit().Body(audit).Do(context.Background()); err != nil {
		return errors.Wrap(err, "Error when update security audit")
	}
	return nil
}

// SecurityAuditGet permit to get the security audit
func (h *OpensearchHandlerImpl) SecurityAuditGet() (audit *opensearch.SecurityGetAuditResponse, err error) {
	audit, err = h.client.SecurityGetAudit().Do(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "Error when gets security audit")
	}
	return audit, nil
}

// SecurityAuditDiff permit to diff a security audit (it use the 3-way diff)
func (h *OpensearchHandlerImpl) SecurityAuditDiff(actualObject, expectedObject, originalObject *opensearch.SecurityAudit) (patchResult *patch.PatchResult, err error) {
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
