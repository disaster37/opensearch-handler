package opensearchhandler

import (
	"context"

	"github.com/disaster37/generic-objectmatcher/patch"
	"github.com/disaster37/opensearch/v2"
	jsonIterator "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// ActionGroupUpdate permit to update an action group
func (h *OpensearchHandlerImpl) ActionGroupUpdate(name string, actionGroup *opensearch.SecurityPutActionGroup) (err error) {
	if _, err = h.client.SecurityPutActionGroup(name).Body(actionGroup).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when update action group '%s'", name)
	}
	return nil
}

// ActionGroupDelete permit to delete an action group
func (h *OpensearchHandlerImpl) ActionGroupDelete(name string) (err error) {
	if _, err = h.client.SecurityDeleteActionGroup(name).Do(context.Background()); err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when delete action group '%s'", name)
	}
	return nil
}

// ActionGroupGet permit to get an action group
func (h *OpensearchHandlerImpl) ActionGroupGet(name string) (actionGroup *opensearch.SecurityActionGroup, err error) {
	res, err := h.client.SecurityGetActionGroup(name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "Error when get action group '%s'", name)
	}

	ag := (*res)[name]

	return &ag, nil
}

// ActionGroupDiff permit to diff an action group (it use the 3-way diff)
func (h *OpensearchHandlerImpl) ActionGroupDiff(actualObject, expectedObject, originalObject *opensearch.SecurityPutActionGroup) (patchResult *patch.PatchResult, err error) {
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
