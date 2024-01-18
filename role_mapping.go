package opensearchhandler

import (
	"context"

	"github.com/disaster37/generic-objectmatcher/patch"
	"github.com/disaster37/opensearch/v2"
	jsonIterator "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// RoleMappingUpdate permit to create or update role mapping
func (h *OpensearchHandlerImpl) RoleMappingUpdate(name string, roleMapping *opensearch.SecurityRoleMapping) (err error) {

	if _, err = h.client.SecurityPutRoleMapping(name).Body(roleMapping).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when update role mapping '%s'", name)
	}

	return nil
}

// RoleMappingDelete permit to delete role mapping
func (h *OpensearchHandlerImpl) RoleMappingDelete(name string) (err error) {

	if _, err = h.client.SecurityDeleteRoleMapping(name).Do(context.Background()); err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when delete role mapping '%s'", name)
	}

	return nil
}

// RoleMappingGet permit to get role mapping
func (h *OpensearchHandlerImpl) RoleMappingGet(name string) (roleMapping *opensearch.SecurityRoleMapping, err error) {

	roleMappingResp, err := h.client.SecurityGetRoleMapping(name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "Error when get role mapping '%s'", name)
	}

	tmp := (*roleMappingResp)[name]
	return &tmp, nil
}

// RoleMappingDiff permit to check if 2 role mapping are the same
func (h *OpensearchHandlerImpl) RoleMappingDiff(actualObject, expectedObject, originalObject *opensearch.SecurityRoleMapping) (patchResult *patch.PatchResult, err error) {
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
