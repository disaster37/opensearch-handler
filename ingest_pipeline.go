package opensearchhandler

import (
	"context"

	"github.com/disaster37/generic-objectmatcher/patch"
	"github.com/disaster37/opensearch/v2"
	jsonIterator "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// IngestPipelineUpdate permit to create or update ingest pipeline
func (h *OpensearchHandlerImpl) IngestPipelineUpdate(name string, pipeline *opensearch.IngestGetPipeline) (err error) {

	if _, err = h.client.IngestPutPipeline(name).BodyJson(pipeline).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when update ingest pipeline '%s'", name)
	}

	return nil

}

// IngestPipelineDelete permit to delete ingest pipeline
func (h *OpensearchHandlerImpl) IngestPipelineDelete(name string) (err error) {

	if _, err = h.client.IngestDeletePipeline(name).Do(context.Background()); err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when delete ingest pipeline '%s'", name)
	}

	return nil
}

// IngestPipelineGet permit to get ingest pipeline
func (h *OpensearchHandlerImpl) IngestPipelineGet(name string) (pipeline *opensearch.IngestGetPipeline, err error) {

	pipelineResp, err := h.client.IngestGetPipeline(name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "Error when get ingest pipeline '%s'", name)
	}

	return pipelineResp[name], nil
}

// IngestPipelineDiff permit to check if 2 ingest pipeline are the same
func (h *OpensearchHandlerImpl) IngestPipelineDiff(actualObject, expectedObject, originalObject *opensearch.IngestGetPipeline) (patchResult *patch.PatchResult, err error) {
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
