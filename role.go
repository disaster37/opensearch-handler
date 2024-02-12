package opensearchhandler

import (
	"context"

	"emperror.dev/errors"
	"github.com/disaster37/generic-objectmatcher/patch"
	"github.com/disaster37/opensearch/v2"
	jsonIterator "github.com/json-iterator/go"
)

// RoleUpdate permit to update role
func (h *OpensearchHandlerImpl) RoleUpdate(name string, role *opensearch.SecurityPutRole) (err error) {

	if _, err = h.client.SecurityPutRole(name).Body(role).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when update role '%s'", name)
	}

	return nil
}

// RoleDelete permit to delete role
func (h *OpensearchHandlerImpl) RoleDelete(name string) (err error) {

	if _, err = h.client.SecurityDeleteRole(name).Do(context.Background()); err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when delete role '%s'", name)
	}

	return nil
}

// RoleGet permit to get role
func (h *OpensearchHandlerImpl) RoleGet(name string) (role *opensearch.SecurityRole, err error) {

	roleResp, err := h.client.SecurityGetRole(name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "Error when get role '%s'", name)
	}

	tmp := (*roleResp)[name]

	return &tmp, nil
}

// RoleDiff permit to check if 2 role are the same
func (h *OpensearchHandlerImpl) RoleDiff(actualObject, expectedObject, originalObject *opensearch.SecurityPutRole) (patchResult *patch.PatchResult, err error) {
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
