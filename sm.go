package opensearchhandler

import (
	"context"

	"emperror.dev/errors"
	"github.com/disaster37/generic-objectmatcher/patch"
	"github.com/disaster37/opensearch/v2"
	jsonIterator "github.com/json-iterator/go"
)

func (h *OpensearchHandlerImpl) SmCreate(name string, policy *opensearch.SmPutPolicy) (err error) {
	if _, err = h.client.SmPostPolicy(name).Body(policy).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when create SM policy '%s'", name)
	}

	return nil
}

func (h *OpensearchHandlerImpl) SmUpdate(name string, sequenceNumber int64, pimaryTerm int64, policy *opensearch.SmPutPolicy) (err error) {
	if _, err = h.client.SmPutPolicy(name).SequenceNumber(sequenceNumber).PrimaryTerm(pimaryTerm).Body(policy).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when update SM policy '%s'", name)
	}

	return nil
}

func (h *OpensearchHandlerImpl) SmDelete(name string) (err error) {
	if _, err = h.client.SmDeletePolicy(name).Do(context.Background()); err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when delete SM policy '%s'", name)
	}

	return nil
}

func (h *OpensearchHandlerImpl) SmGet(name string) (policy *opensearch.SmGetPolicyResponse, err error) {
	res, err := h.client.SmGetPolicy(name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "Error when get SM policy '%s'", name)
	}

	return res, nil
}

func (h *OpensearchHandlerImpl) SmDiff(actualObject, expectedObject, originalObject *opensearch.SmPutPolicy) (patchResult *patch.PatchResult, err error) {
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
