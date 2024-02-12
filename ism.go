package opensearchhandler

import (
	"context"

	"emperror.dev/errors"
	"github.com/disaster37/generic-objectmatcher/patch"
	"github.com/disaster37/opensearch/v2"
	jsonIterator "github.com/json-iterator/go"
)

// IsmCreate permit to crate new ISM policy
func (h *OpensearchHandlerImpl) IsmCreate(name string, policy *opensearch.IsmPutPolicy) (err error) {
	if _, err = h.client.IsmPutPolicy(name).Body(policy).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when create ISM policy '%s'", name)
	}

	return nil
}

// IsmUpdate permit to update ISM policy
func (h *OpensearchHandlerImpl) IsmUpdate(name string, sequenceNumber int64, pimaryTerm int64, policy *opensearch.IsmPutPolicy) (err error) {
	if _, err = h.client.IsmPutPolicy(name).SequenceNumber(sequenceNumber).PrimaryTerm(pimaryTerm).Body(policy).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when update ISM policy '%s'", name)
	}

	return nil
}

// IsmDelete permit to delete ISM policy
func (h *OpensearchHandlerImpl) IsmDelete(name string) (err error) {
	if _, err = h.client.IsmDeletePolicy(name).Do(context.Background()); err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when delete ISM policy '%s'", name)
	}

	return nil
}

// IsmGet permit to get ISM policy
func (h *OpensearchHandlerImpl) IsmGet(name string) (policy *opensearch.IsmGetPolicyResponse, err error) {
	res, err := h.client.IsmGetPolicy(name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "Error when get ISM policy '%s", name)
	}

	return res, nil
}

// IsmDiff permit to diff ISM policy
func (h *OpensearchHandlerImpl) IsmDiff(actualObject, expectedObject, originalObject *opensearch.IsmPutPolicy) (patchResult *patch.PatchResult, err error) {
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

	return patch.DefaultPatchMaker.Calculate(actualObject, expectedObject, originalObject)
}
