package opensearchhandler

import (
	"context"

	"emperror.dev/errors"
	"github.com/disaster37/generic-objectmatcher/patch"
	"github.com/disaster37/opensearch/v2"
	jsonIterator "github.com/json-iterator/go"
)

// SnapshotRepositoryUpdate permit to create or update snapshot repository
func (h *OpensearchHandlerImpl) SnapshotRepositoryUpdate(name string, repository *opensearch.SnapshotRepositoryMetaData) (err error) {

	if _, err = h.client.SnapshotCreateRepository(name).BodyJson(repository).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when update snapshot repository '%s'", name)
	}

	return nil
}

// SnapshotRepositoryDelete permit to delete snapshot repository
func (h *OpensearchHandlerImpl) SnapshotRepositoryDelete(name string) (err error) {

	if _, err = h.client.SnapshotDeleteRepository(name).Do(context.Background()); err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when delete snapshot repository '%s'", name)
	}

	return nil
}

// SnapshotRepositoryGet permit to get snapshot repository
func (h *OpensearchHandlerImpl) SnapshotRepositoryGet(name string) (repository *opensearch.SnapshotRepositoryMetaData, err error) {

	snapshotRepository, err := h.client.SnapshotGetRepository(name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "Error when delete snapshot repository '%s'", name)
	}

	return snapshotRepository[name], nil

}

// SnapshotRepositoryDiff permit to check if 2 repositories are the same
func (h *OpensearchHandlerImpl) SnapshotRepositoryDiff(actualObject, expectedObject, originalObject *opensearch.SnapshotRepositoryMetaData) (patchResult *patch.PatchResult, err error) {
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
