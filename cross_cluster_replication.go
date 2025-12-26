package opensearchhandler

import (
	"context"

	"emperror.dev/errors"
	"github.com/disaster37/generic-objectmatcher/patch"
	"github.com/disaster37/opensearch/v3"
	jsonIterator "github.com/json-iterator/go"
)

// CrossClusterReplicationStart permit to start CCR rule
func (h *OpensearchHandlerImpl) CrossClusterReplicationStart(name string, ccrRule *opensearch.CcrRule) (err error) {

	// Check if rule already exist
	ccrStatus, err := h.CrossClusterReplicationStatus(name)
	if err != nil {
		return errors.Wrapf(err, "Error when get CCR rule '%s'", name)
	}
	if ccrStatus != nil {
		h.log.Debugf("CCR rule '%s' already exist", name)
		// Check if on pause to resume or delete it
		if ccrStatus.Status == string(opensearch.CcrStatusPaused) {
			h.log.Debugf("CCR rule '%s' is paused", name)
			if _, err = h.Client().CcrResumeRule(name).Do(context.Background()); err != nil {
				h.log.Warnf("Error when resume CCR rule '%s', we will delete it", name)

				if err = h.CrossClusterReplicationStop(name); err != nil {
					return errors.Wrapf(err, "Error when delete CCR rule '%s'", name)
				}

				h.log.Infof("CCR rule '%s' deleted", name)

			} else {
				h.log.Infof("CCR rule '%s' resumed", name)
				return nil
			}
		}
	}

	resp, err := h.Client().CcrStartRule(name).Body(ccrRule).Do(context.Background())
	if err != nil {
		return errors.Wrapf(err, "Error when start CCR rule '%s'", name)
	}
	if !resp.Acknowledged {
		return errors.Errorf("CCR rule '%s' not started", name)
	}

	return nil
}

// CrossClusterReplicationStop permit to stop CCR rule
func (h *OpensearchHandlerImpl) CrossClusterReplicationStop(name string) (err error) {
	resp, err := h.Client().CcrStopRule(name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when stop CCR rule '%s'", name)
	}
	if !resp.Acknowledged {
		return errors.Errorf("CCR rule '%s' not stopped", name)
	}

	return nil

}

func (h *OpensearchHandlerImpl) CrossClusterReplicationPause(name string) (err error) {

	resp, err := h.Client().CcrPauseRule(name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when pause CCR rule '%s'", name)
	}
	if !resp.Acknowledged {
		return errors.Errorf("CCR rule '%s' not paused", name)
	}

	return nil
}

// CrossClusterReplicationStatus permit to get CCR status
func (h *OpensearchHandlerImpl) CrossClusterReplicationStatus(name string) (ccrStatus *opensearch.CcrStatusRuleResponse, err error) {

	ccrStatus, err = h.Client().CcrStatusRule(name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "Error when get CCR rule '%s'", name)
	}

	return ccrStatus, nil

}

// CrossClusterReplicationDiff permit to diff a CCR rule
func (h *OpensearchHandlerImpl) CrossClusterReplicationDiff(actualObject, expectedObject, originalObject *opensearch.CcrRule) (patchResult *patch.PatchResult, err error) {

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

// CrossClusterReplicationAutoFollowCreate permit to create a CCR auto follow rule
func (h *OpensearchHandlerImpl) CrossClusterReplicationAutoFollowCreate(ccrRule *opensearch.CcrAutoFollowRule) (err error) {

	resp, err := h.Client().CcrPostAutoFollow().Body(ccrRule).Do(context.Background())
	if err != nil {
		return errors.Wrapf(err, "Error when create CCR auto follow rule '%s'", ccrRule.Name)
	}
	if !resp.Acknowledged {
		return errors.Errorf("CCR auto follow rule '%s' not created", ccrRule.Name)
	}

	return nil
}

// CrossClusterReplicationAutoFollowStatus permit to get CCR auto follow status
func (h *OpensearchHandlerImpl) CrossClusterReplicationAutoFollowStatus(name string) (ccrStatus *opensearch.CcrAutoFollowStatus, err error) {
	resp, err := h.Client().CcrAutoFollowStatus().Do(context.Background())
	if err != nil {
		return nil, errors.Wrapf(err, "Error when get CCR auto follow rule '%s'", name)
	}

	for _, ccrStatus := range resp.AutofollowStats {
		if ccrStatus.Name == name {
			return &ccrStatus, nil
		}
	}

	return nil, nil
}

// CrossClusterReplicationAutoFollowDelete permit to delete a CCR auto follow rule
func (h *OpensearchHandlerImpl) CrossClusterReplicationAutoFollowDelete(name, leaderAlias string) (err error) {

	resp, err := h.Client().CcrDeleteAutoFollow(leaderAlias, name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when delete CCR auto follow rule '%s'", name)
	}
	if !resp.Acknowledged {
		return errors.Errorf("CCR auto follow rule '%s' not deleted", name)
	}

	return nil

}

func (h *OpensearchHandlerImpl) CrossClusterReplicationAutoFollowDiff(actualObject, expectedObject, originalObject *opensearch.CcrAutoFollowRule) (patchResult *patch.PatchResult, err error) {

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
