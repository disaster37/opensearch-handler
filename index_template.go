package opensearchhandler

import (
	"context"

	"github.com/disaster37/generic-objectmatcher/patch"
	localpatch "github.com/disaster37/opensearch-handler/v2/patch"
	"github.com/disaster37/opensearch/v2"
	jsonIterator "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// IndexTemplateUpdate permit to create or update index template
func (h *OpensearchHandlerImpl) IndexTemplateUpdate(name string, template *opensearch.IndicesGetIndexTemplate) (err error) {

	if _, err = h.client.IndexPutIndexTemplate(name).BodyJson(template).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when update index template '%s'", name)
	}

	return nil

}

// IndexTemplateDelete permit to delete index template
func (h *OpensearchHandlerImpl) IndexTemplateDelete(name string) (err error) {

	if _, err = h.client.IndexDeleteIndexTemplate(name).Do(context.Background()); err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when delete index template '%s'", name)
	}

	return nil
}

// IndexTemplateGet permit to get index template
func (h *OpensearchHandlerImpl) IndexTemplateGet(name string) (template *opensearch.IndicesGetIndexTemplate, err error) {

	indexTemplate, err := h.client.IndexGetIndexTemplate(name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "Error when get index template '%s'", name)
	}

	if len(indexTemplate.IndexTemplates) == 0 {
		return nil, nil
	}

	return indexTemplate.IndexTemplates[0].IndexTemplate, nil
}

// IndexTemplateDiff permit to check if 2 index template is the same
func (h *OpensearchHandlerImpl) IndexTemplateDiff(actualObject, expectedObject, originalObject *opensearch.IndicesGetIndexTemplate) (patchResult *patch.PatchResult, err error) {
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

	return patch.DefaultPatchMaker.Calculate(actualObject, expectedObject, originalObject, localpatch.ConvertIndexTemplateSetting)
}
