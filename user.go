package opensearchhandler

import (
	"context"

	"github.com/disaster37/generic-objectmatcher/patch"
	"github.com/disaster37/opensearch/v2"
	jsonIterator "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// UserCreate permit to create new user
func (h *OpensearchHandlerImpl) UserUpdate(name string, user *opensearch.SecurityPutUser) (err error) {

	if _, err = h.client.SecurityPutUser(name).Body(user).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when update user '%s'", name)
	}

	return nil
}

// UserDelete permit to delete the user
func (h *OpensearchHandlerImpl) UserDelete(name string) (err error) {

	if _, err = h.client.SecurityDeleteUser(name).Do(context.Background()); err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when delete user '%s'", name)
	}

	return nil
}

// UserGet permit to get the user
func (h *OpensearchHandlerImpl) UserGet(name string) (user *opensearch.SecurityUser, err error) {

	userResp, err := h.client.SecurityGetUser(name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "Error when get user '%s'", name)
	}

	tmp := (*userResp)[name]
	return &tmp, nil
}

// UserDiff permit to check if 2 users are the same
func (h *OpensearchHandlerImpl) UserDiff(actualObject, expectedObject, originalObject *opensearch.SecurityPutUser) (patchResult *patch.PatchResult, err error) {
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
