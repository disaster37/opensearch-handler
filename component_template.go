package opensearchhandler

import (
	"context"

	"github.com/disaster37/generic-objectmatcher/patch"
	localpatch "github.com/disaster37/opensearch-handler/v2/patch"
	"github.com/disaster37/opensearch/v2"
	jsonIterator "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// ComponentTemplateUpdate permit to update component template
func (h *OpensearchHandlerImpl) ComponentTemplateUpdate(name string, component *opensearch.IndicesGetComponentTemplate) (err error) {

	if _, err := h.client.IndexPutComponentTemplate(name).BodyJson(component).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when update Component template '%s'", name)
	}

	return nil
}

// ComponentTemplateDelete permit to delete component template
func (h *OpensearchHandlerImpl) ComponentTemplateDelete(name string) (err error) {

	if _, err = h.client.IndexDeleteComponentTemplate(name).Do(context.Background()); err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when delete component template '%s'", name)
	}

	return nil

}

// ComponentTemplateGet permit to get component template
func (h *OpensearchHandlerImpl) ComponentTemplateGet(name string) (component *opensearch.IndicesGetComponentTemplate, err error) {

	indexComponentTemplateResp, err := h.client.IndexGetComponentTemplate(name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "Error when get component template '%s'", name)
	}

	if len(indexComponentTemplateResp.ComponentTemplates) == 0 {
		return nil, nil
	}

	return indexComponentTemplateResp.ComponentTemplates[0].ComponentTemplate, nil
}

// ComponentTemplateDiff permit to check if 2 component template are the same
func (h *OpensearchHandlerImpl) ComponentTemplateDiff(actualObject, expectedObject, originalObject *opensearch.IndicesGetComponentTemplate) (patchResult *patch.PatchResult, err error) {
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

	return patch.DefaultPatchMaker.Calculate(actualObject, expectedObject, originalObject, localpatch.ConvertComponentTemplateSetting)
}
