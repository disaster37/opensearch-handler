package opensearchhandler

import (
	"github.com/disaster37/generic-objectmatcher/patch"
	"github.com/disaster37/opensearch/v2"
	"github.com/disaster37/opensearch/v2/config"
	"github.com/sirupsen/logrus"
)

type OpensearchHandler interface {
	Client() (client *opensearch.Client)

	// Snapshot repository scope
	SnapshotRepositoryUpdate(name string, repository *opensearch.SnapshotRepositoryMetaData) (err error)
	SnapshotRepositoryDelete(name string) (err error)
	SnapshotRepositoryGet(name string) (repository *opensearch.SnapshotRepositoryMetaData, err error)
	SnapshotRepositoryDiff(actualObject, expectedObject, originalObject *opensearch.SnapshotRepositoryMetaData) (patchResult *patch.PatchResult, err error)

	// Component template scope
	ComponentTemplateUpdate(name string, component *opensearch.IndicesGetComponentTemplate) (err error)
	ComponentTemplateDelete(name string) (err error)
	ComponentTemplateGet(name string) (component *opensearch.IndicesGetComponentTemplate, err error)
	ComponentTemplateDiff(actualObject, expectedObject, originalObject *opensearch.IndicesGetComponentTemplate) (patchResult *patch.PatchResult, err error)

	// Index template scope
	IndexTemplateUpdate(name string, template *opensearch.IndicesGetIndexTemplate) (err error)
	IndexTemplateDelete(name string) (err error)
	IndexTemplateGet(name string) (template *opensearch.IndicesGetIndexTemplate, err error)
	IndexTemplateDiff(actualObject, expectedObject, originalObject *opensearch.IndicesGetIndexTemplate) (patchResult *patch.PatchResult, err error)

	// Ingest pipline scope
	IngestPipelineUpdate(name string, pipeline *opensearch.IngestGetPipeline) (err error)
	IngestPipelineDelete(name string) (err error)
	IngestPipelineGet(name string) (pipeline *opensearch.IngestGetPipeline, err error)
	IngestPipelineDiff(actualObject, expectedObject, originalObject *opensearch.IngestGetPipeline) (patchResult *patch.PatchResult, err error)

	// Cluster scope
	ClusterHealth() (health *opensearch.ClusterHealthResponse, err error)

	SetLogger(log *logrus.Entry)
}

type OpensearchHandlerImpl struct {
	client *opensearch.Client
	log    *logrus.Entry
}

func NewOpensearchHandler(cfg *config.Config, log *logrus.Entry) (OpensearchHandler, error) {

	client, err := opensearch.NewClientFromConfig(cfg)
	if err != nil {
		return nil, err
	}

	return &OpensearchHandlerImpl{
		client: client,
		log:    log,
	}, nil
}

func (h *OpensearchHandlerImpl) SetLogger(log *logrus.Entry) {
	h.log = log
}

func (h *OpensearchHandlerImpl) Client() *opensearch.Client {
	return h.client
}
