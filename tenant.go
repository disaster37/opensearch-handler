package opensearchhandler

import (
	"context"

	"github.com/disaster37/generic-objectmatcher/patch"
	"github.com/disaster37/opensearch/v2"
	jsonIterator "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// TenantUpdate permit to update a tenant
func (h *OpensearchHandlerImpl) TenantUpdate(name string, tenant *opensearch.SecurityPutTenant) (err error) {
	if _, err = h.client.SecurityPutTenant(name).Body(tenant).Do(context.Background()); err != nil {
		return errors.Wrapf(err, "Error when update tenant '%s'", name)
	}

	return nil
}

// TenantDelete permit to delete a tenant
func (h *OpensearchHandlerImpl) TenantDelete(name string) (err error) {
	if _, err = h.client.SecurityDeleteTenant(name).Do(context.Background()); err != nil {
		if opensearch.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "Error when delete tenant '%s'", name)
	}

	return nil
}

// TenantGet permit to get a tenant
func (h *OpensearchHandlerImpl) TenantGet(name string) (actionGroup *opensearch.SecurityTenant, err error) {
	res, err := h.client.SecurityGetTenant(name).Do(context.Background())
	if err != nil {
		if opensearch.IsNotFound(err) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "Error when get tenant '%s'", name)
	}

	tenant := (*res)[name]

	return &tenant, nil
}

// TenantDiff permit to diff a tenant (it use a 3-way diff)
func (h *OpensearchHandlerImpl) TenantDiff(actualObject, expectedObject, originalObject *opensearch.SecurityPutTenant) (patchResult *patch.PatchResult, err error) {
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
