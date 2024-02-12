package opensearchhandler

import (
	"context"

	"emperror.dev/errors"
	"github.com/disaster37/generic-objectmatcher/patch"
	"github.com/disaster37/opensearch/v2"
	jsonIterator "github.com/json-iterator/go"
)

func (h *OpensearchHandlerImpl) MonitorCreate(monitor *opensearch.AlertingMonitor) (err error) {
	if _, err = h.client.AlertingPostMonitor().Body(monitor).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when create monitor '%s'", monitor.Name)
	}

	return nil
}
func (h *OpensearchHandlerImpl) MonitorUpdate(id string, sequenceNumber int64, pimaryTerm int64, monitor *opensearch.AlertingMonitor) (err error) {
	if _, err = h.client.AlertingPutMonitor(id).SequenceNumber(sequenceNumber).PrimaryTerm(pimaryTerm).Body(monitor).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when update monitor '%s'", monitor.Name)
	}

	return nil
}
func (h *OpensearchHandlerImpl) MonitorDelete(id string) (err error) {
	if _, err = h.client.AlertingDeleteMonitor(id).Do(context.Background()); err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when delete monitor '%s'", id)
	}
	return
}
func (h *OpensearchHandlerImpl) MonitorGet(name string) (monitor *opensearch.AlertingGetMonitor, err error) {
	res, err := h.client.AlertingSearchMonitor().SearchByName(name).Do(context.Background())
	if err != nil {
		return nil, errors.Wrapf(err, "Error when get monitor '%s'", name)
	}
	if len(res) == 0 {
		return nil, nil
	}

	return &(res[0]), nil
}
func (h *OpensearchHandlerImpl) MonitorDiff(actualObject, expectedObject, originalObject *opensearch.AlertingMonitor) (patchResult *patch.PatchResult, err error) {
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
