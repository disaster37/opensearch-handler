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

	// Role scope
	RoleUpdate(name string, role *opensearch.SecurityPutRole) (err error)
	RoleDelete(name string) (err error)
	RoleGet(name string) (role *opensearch.SecurityRole, err error)
	RoleDiff(actualObject, expectedObject, originalObject *opensearch.SecurityPutRole) (patchResult *patch.PatchResult, err error)

	// Role mapping scope
	RoleMappingUpdate(name string, roleMapping *opensearch.SecurityPutRoleMapping) (err error)
	RoleMappingDelete(name string) (err error)
	RoleMappingGet(name string) (roleMapping *opensearch.SecurityRoleMapping, err error)
	RoleMappingDiff(actualObject, expectedObject, originalObject *opensearch.SecurityPutRoleMapping) (patchResult *patch.PatchResult, err error)

	// User scope
	UserUpdate(name string, user *opensearch.SecurityPutUser) (err error)
	UserDelete(name string) (err error)
	UserGet(name string) (user *opensearch.SecurityUser, err error)
	UserDiff(actualObject, expectedObject, originalObject *opensearch.SecurityPutUser) (patchResult *patch.PatchResult, err error)

	// Action groups scope
	ActionGroupUpdate(name string, actionGroup *opensearch.SecurityPutActionGroup) (err error)
	ActionGroupDelete(name string) (err error)
	ActionGroupGet(name string) (actionGroup *opensearch.SecurityActionGroup, err error)
	ActionGroupDiff(actualObject, expectedObject, originalObject *opensearch.SecurityPutActionGroup) (patchResult *patch.PatchResult, err error)

	// Tenanats scope
	TenantUpdate(name string, tenant *opensearch.SecurityPutTenant) (err error)
	TenantDelete(name string) (err error)
	TenantGet(name string) (tenant *opensearch.SecurityTenant, err error)
	TenantDiff(actualObject, expectedObject, originalObject *opensearch.SecurityPutTenant) (patchResult *patch.PatchResult, err error)

	// Security config scope
	SecurityConfigUpdate(config *opensearch.SecurityConfig) (err error)
	SecurityConfigGet() (config *opensearch.SecurityGetConfigResponse, err error)
	SecurityConfigDiff(actualObject, expectedObject, originalObject *opensearch.SecurityConfig) (patchResult *patch.PatchResult, err error)

	// Security audit scope
	SecurityAuditUpdate(config *opensearch.SecurityAudit) (err error)
	SecurityAuditGet() (config *opensearch.SecurityGetAuditResponse, err error)
	SecurityAuditDiff(actualObject, expectedObject, originalObject *opensearch.SecurityAudit) (patchResult *patch.PatchResult, err error)

	// Index State management scope
	IsmCreate(name string, policy *opensearch.IsmPutPolicy) (err error)
	IsmUpdate(name string, sequenceNumber int64, pimaryTerm int64, policy *opensearch.IsmPutPolicy) (err error)
	IsmDelete(name string) (err error)
	IsmGet(name string) (policy *opensearch.IsmGetPolicyResponse, err error)
	IsmDiff(actualObject, expectedObject, originalObject *opensearch.IsmPutPolicy) (patchResult *patch.PatchResult, err error)

	// Snapshot management scope
	SmCreate(name string, policy *opensearch.SmPutPolicy) (err error)
	SmUpdate(name string, sequenceNumber int64, pimaryTerm int64, policy *opensearch.SmPutPolicy) (err error)
	SmDelete(name string) (err error)
	SmGet(name string) (policy *opensearch.SmGetPolicyResponse, err error)
	SmDiff(actualObject, expectedObject, originalObject *opensearch.SmPutPolicy) (patchResult *patch.PatchResult, err error)

	// Monitor management scope
	MonitorCreate(monitor *opensearch.AlertingMonitor) (err error)
	MonitorUpdate(id string, sequenceNumber int64, pimaryTerm int64, monitor *opensearch.AlertingMonitor) (err error)
	MonitorDelete(id string) (err error)
	MonitorGet(name string) (monitor *opensearch.AlertingGetMonitorResponse, err error)
	MonitorDiff(actualObject, expectedObject, originalObject *opensearch.AlertingMonitor) (patchResult *patch.PatchResult, err error)

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
