// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/disaster37/opensearch-handler/v2 (interfaces: OpensearchHandler)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination=mocks/opensearch_handler.go -package=mocks github.com/disaster37/opensearch-handler/v2 OpensearchHandler
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	patch "github.com/disaster37/generic-objectmatcher/patch"
	opensearch "github.com/disaster37/opensearch/v2"
	logrus "github.com/sirupsen/logrus"
	gomock "go.uber.org/mock/gomock"
)

// MockOpensearchHandler is a mock of OpensearchHandler interface.
type MockOpensearchHandler struct {
	ctrl     *gomock.Controller
	recorder *MockOpensearchHandlerMockRecorder
}

// MockOpensearchHandlerMockRecorder is the mock recorder for MockOpensearchHandler.
type MockOpensearchHandlerMockRecorder struct {
	mock *MockOpensearchHandler
}

// NewMockOpensearchHandler creates a new mock instance.
func NewMockOpensearchHandler(ctrl *gomock.Controller) *MockOpensearchHandler {
	mock := &MockOpensearchHandler{ctrl: ctrl}
	mock.recorder = &MockOpensearchHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpensearchHandler) EXPECT() *MockOpensearchHandlerMockRecorder {
	return m.recorder
}

// ActionGroupDelete mocks base method.
func (m *MockOpensearchHandler) ActionGroupDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionGroupDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActionGroupDelete indicates an expected call of ActionGroupDelete.
func (mr *MockOpensearchHandlerMockRecorder) ActionGroupDelete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionGroupDelete", reflect.TypeOf((*MockOpensearchHandler)(nil).ActionGroupDelete), arg0)
}

// ActionGroupDiff mocks base method.
func (m *MockOpensearchHandler) ActionGroupDiff(arg0, arg1, arg2 *opensearch.SecurityPutActionGroup) (*patch.PatchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionGroupDiff", arg0, arg1, arg2)
	ret0, _ := ret[0].(*patch.PatchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionGroupDiff indicates an expected call of ActionGroupDiff.
func (mr *MockOpensearchHandlerMockRecorder) ActionGroupDiff(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionGroupDiff", reflect.TypeOf((*MockOpensearchHandler)(nil).ActionGroupDiff), arg0, arg1, arg2)
}

// ActionGroupGet mocks base method.
func (m *MockOpensearchHandler) ActionGroupGet(arg0 string) (*opensearch.SecurityActionGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionGroupGet", arg0)
	ret0, _ := ret[0].(*opensearch.SecurityActionGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionGroupGet indicates an expected call of ActionGroupGet.
func (mr *MockOpensearchHandlerMockRecorder) ActionGroupGet(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionGroupGet", reflect.TypeOf((*MockOpensearchHandler)(nil).ActionGroupGet), arg0)
}

// ActionGroupUpdate mocks base method.
func (m *MockOpensearchHandler) ActionGroupUpdate(arg0 string, arg1 *opensearch.SecurityPutActionGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionGroupUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActionGroupUpdate indicates an expected call of ActionGroupUpdate.
func (mr *MockOpensearchHandlerMockRecorder) ActionGroupUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionGroupUpdate", reflect.TypeOf((*MockOpensearchHandler)(nil).ActionGroupUpdate), arg0, arg1)
}

// Client mocks base method.
func (m *MockOpensearchHandler) Client() *opensearch.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(*opensearch.Client)
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockOpensearchHandlerMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockOpensearchHandler)(nil).Client))
}

// ClusterHealth mocks base method.
func (m *MockOpensearchHandler) ClusterHealth() (*opensearch.ClusterHealthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterHealth")
	ret0, _ := ret[0].(*opensearch.ClusterHealthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClusterHealth indicates an expected call of ClusterHealth.
func (mr *MockOpensearchHandlerMockRecorder) ClusterHealth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterHealth", reflect.TypeOf((*MockOpensearchHandler)(nil).ClusterHealth))
}

// ComponentTemplateDelete mocks base method.
func (m *MockOpensearchHandler) ComponentTemplateDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentTemplateDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ComponentTemplateDelete indicates an expected call of ComponentTemplateDelete.
func (mr *MockOpensearchHandlerMockRecorder) ComponentTemplateDelete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentTemplateDelete", reflect.TypeOf((*MockOpensearchHandler)(nil).ComponentTemplateDelete), arg0)
}

// ComponentTemplateDiff mocks base method.
func (m *MockOpensearchHandler) ComponentTemplateDiff(arg0, arg1, arg2 *opensearch.IndicesGetComponentTemplate) (*patch.PatchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentTemplateDiff", arg0, arg1, arg2)
	ret0, _ := ret[0].(*patch.PatchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComponentTemplateDiff indicates an expected call of ComponentTemplateDiff.
func (mr *MockOpensearchHandlerMockRecorder) ComponentTemplateDiff(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentTemplateDiff", reflect.TypeOf((*MockOpensearchHandler)(nil).ComponentTemplateDiff), arg0, arg1, arg2)
}

// ComponentTemplateGet mocks base method.
func (m *MockOpensearchHandler) ComponentTemplateGet(arg0 string) (*opensearch.IndicesGetComponentTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentTemplateGet", arg0)
	ret0, _ := ret[0].(*opensearch.IndicesGetComponentTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComponentTemplateGet indicates an expected call of ComponentTemplateGet.
func (mr *MockOpensearchHandlerMockRecorder) ComponentTemplateGet(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentTemplateGet", reflect.TypeOf((*MockOpensearchHandler)(nil).ComponentTemplateGet), arg0)
}

// ComponentTemplateUpdate mocks base method.
func (m *MockOpensearchHandler) ComponentTemplateUpdate(arg0 string, arg1 *opensearch.IndicesGetComponentTemplate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentTemplateUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ComponentTemplateUpdate indicates an expected call of ComponentTemplateUpdate.
func (mr *MockOpensearchHandlerMockRecorder) ComponentTemplateUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentTemplateUpdate", reflect.TypeOf((*MockOpensearchHandler)(nil).ComponentTemplateUpdate), arg0, arg1)
}

// ConfigDiff mocks base method.
func (m *MockOpensearchHandler) ConfigDiff(arg0, arg1, arg2 *opensearch.SecurityConfig) (*patch.PatchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigDiff", arg0, arg1, arg2)
	ret0, _ := ret[0].(*patch.PatchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigDiff indicates an expected call of ConfigDiff.
func (mr *MockOpensearchHandlerMockRecorder) ConfigDiff(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigDiff", reflect.TypeOf((*MockOpensearchHandler)(nil).ConfigDiff), arg0, arg1, arg2)
}

// ConfigGet mocks base method.
func (m *MockOpensearchHandler) ConfigGet() (*opensearch.SecurityGetConfigResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigGet")
	ret0, _ := ret[0].(*opensearch.SecurityGetConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigGet indicates an expected call of ConfigGet.
func (mr *MockOpensearchHandlerMockRecorder) ConfigGet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigGet", reflect.TypeOf((*MockOpensearchHandler)(nil).ConfigGet))
}

// ConfigUpdate mocks base method.
func (m *MockOpensearchHandler) ConfigUpdate(arg0 *opensearch.SecurityConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigUpdate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfigUpdate indicates an expected call of ConfigUpdate.
func (mr *MockOpensearchHandlerMockRecorder) ConfigUpdate(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigUpdate", reflect.TypeOf((*MockOpensearchHandler)(nil).ConfigUpdate), arg0)
}

// IndexTemplateDelete mocks base method.
func (m *MockOpensearchHandler) IndexTemplateDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexTemplateDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// IndexTemplateDelete indicates an expected call of IndexTemplateDelete.
func (mr *MockOpensearchHandlerMockRecorder) IndexTemplateDelete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexTemplateDelete", reflect.TypeOf((*MockOpensearchHandler)(nil).IndexTemplateDelete), arg0)
}

// IndexTemplateDiff mocks base method.
func (m *MockOpensearchHandler) IndexTemplateDiff(arg0, arg1, arg2 *opensearch.IndicesGetIndexTemplate) (*patch.PatchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexTemplateDiff", arg0, arg1, arg2)
	ret0, _ := ret[0].(*patch.PatchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IndexTemplateDiff indicates an expected call of IndexTemplateDiff.
func (mr *MockOpensearchHandlerMockRecorder) IndexTemplateDiff(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexTemplateDiff", reflect.TypeOf((*MockOpensearchHandler)(nil).IndexTemplateDiff), arg0, arg1, arg2)
}

// IndexTemplateGet mocks base method.
func (m *MockOpensearchHandler) IndexTemplateGet(arg0 string) (*opensearch.IndicesGetIndexTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexTemplateGet", arg0)
	ret0, _ := ret[0].(*opensearch.IndicesGetIndexTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IndexTemplateGet indicates an expected call of IndexTemplateGet.
func (mr *MockOpensearchHandlerMockRecorder) IndexTemplateGet(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexTemplateGet", reflect.TypeOf((*MockOpensearchHandler)(nil).IndexTemplateGet), arg0)
}

// IndexTemplateUpdate mocks base method.
func (m *MockOpensearchHandler) IndexTemplateUpdate(arg0 string, arg1 *opensearch.IndicesGetIndexTemplate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexTemplateUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IndexTemplateUpdate indicates an expected call of IndexTemplateUpdate.
func (mr *MockOpensearchHandlerMockRecorder) IndexTemplateUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexTemplateUpdate", reflect.TypeOf((*MockOpensearchHandler)(nil).IndexTemplateUpdate), arg0, arg1)
}

// IngestPipelineDelete mocks base method.
func (m *MockOpensearchHandler) IngestPipelineDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngestPipelineDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// IngestPipelineDelete indicates an expected call of IngestPipelineDelete.
func (mr *MockOpensearchHandlerMockRecorder) IngestPipelineDelete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngestPipelineDelete", reflect.TypeOf((*MockOpensearchHandler)(nil).IngestPipelineDelete), arg0)
}

// IngestPipelineDiff mocks base method.
func (m *MockOpensearchHandler) IngestPipelineDiff(arg0, arg1, arg2 *opensearch.IngestGetPipeline) (*patch.PatchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngestPipelineDiff", arg0, arg1, arg2)
	ret0, _ := ret[0].(*patch.PatchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IngestPipelineDiff indicates an expected call of IngestPipelineDiff.
func (mr *MockOpensearchHandlerMockRecorder) IngestPipelineDiff(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngestPipelineDiff", reflect.TypeOf((*MockOpensearchHandler)(nil).IngestPipelineDiff), arg0, arg1, arg2)
}

// IngestPipelineGet mocks base method.
func (m *MockOpensearchHandler) IngestPipelineGet(arg0 string) (*opensearch.IngestGetPipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngestPipelineGet", arg0)
	ret0, _ := ret[0].(*opensearch.IngestGetPipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IngestPipelineGet indicates an expected call of IngestPipelineGet.
func (mr *MockOpensearchHandlerMockRecorder) IngestPipelineGet(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngestPipelineGet", reflect.TypeOf((*MockOpensearchHandler)(nil).IngestPipelineGet), arg0)
}

// IngestPipelineUpdate mocks base method.
func (m *MockOpensearchHandler) IngestPipelineUpdate(arg0 string, arg1 *opensearch.IngestGetPipeline) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngestPipelineUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IngestPipelineUpdate indicates an expected call of IngestPipelineUpdate.
func (mr *MockOpensearchHandlerMockRecorder) IngestPipelineUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngestPipelineUpdate", reflect.TypeOf((*MockOpensearchHandler)(nil).IngestPipelineUpdate), arg0, arg1)
}

// RoleDelete mocks base method.
func (m *MockOpensearchHandler) RoleDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RoleDelete indicates an expected call of RoleDelete.
func (mr *MockOpensearchHandlerMockRecorder) RoleDelete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleDelete", reflect.TypeOf((*MockOpensearchHandler)(nil).RoleDelete), arg0)
}

// RoleDiff mocks base method.
func (m *MockOpensearchHandler) RoleDiff(arg0, arg1, arg2 *opensearch.SecurityPutRole) (*patch.PatchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleDiff", arg0, arg1, arg2)
	ret0, _ := ret[0].(*patch.PatchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoleDiff indicates an expected call of RoleDiff.
func (mr *MockOpensearchHandlerMockRecorder) RoleDiff(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleDiff", reflect.TypeOf((*MockOpensearchHandler)(nil).RoleDiff), arg0, arg1, arg2)
}

// RoleGet mocks base method.
func (m *MockOpensearchHandler) RoleGet(arg0 string) (*opensearch.SecurityRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleGet", arg0)
	ret0, _ := ret[0].(*opensearch.SecurityRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoleGet indicates an expected call of RoleGet.
func (mr *MockOpensearchHandlerMockRecorder) RoleGet(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleGet", reflect.TypeOf((*MockOpensearchHandler)(nil).RoleGet), arg0)
}

// RoleMappingDelete mocks base method.
func (m *MockOpensearchHandler) RoleMappingDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleMappingDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RoleMappingDelete indicates an expected call of RoleMappingDelete.
func (mr *MockOpensearchHandlerMockRecorder) RoleMappingDelete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleMappingDelete", reflect.TypeOf((*MockOpensearchHandler)(nil).RoleMappingDelete), arg0)
}

// RoleMappingDiff mocks base method.
func (m *MockOpensearchHandler) RoleMappingDiff(arg0, arg1, arg2 *opensearch.SecurityPutRoleMapping) (*patch.PatchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleMappingDiff", arg0, arg1, arg2)
	ret0, _ := ret[0].(*patch.PatchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoleMappingDiff indicates an expected call of RoleMappingDiff.
func (mr *MockOpensearchHandlerMockRecorder) RoleMappingDiff(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleMappingDiff", reflect.TypeOf((*MockOpensearchHandler)(nil).RoleMappingDiff), arg0, arg1, arg2)
}

// RoleMappingGet mocks base method.
func (m *MockOpensearchHandler) RoleMappingGet(arg0 string) (*opensearch.SecurityRoleMapping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleMappingGet", arg0)
	ret0, _ := ret[0].(*opensearch.SecurityRoleMapping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoleMappingGet indicates an expected call of RoleMappingGet.
func (mr *MockOpensearchHandlerMockRecorder) RoleMappingGet(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleMappingGet", reflect.TypeOf((*MockOpensearchHandler)(nil).RoleMappingGet), arg0)
}

// RoleMappingUpdate mocks base method.
func (m *MockOpensearchHandler) RoleMappingUpdate(arg0 string, arg1 *opensearch.SecurityPutRoleMapping) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleMappingUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RoleMappingUpdate indicates an expected call of RoleMappingUpdate.
func (mr *MockOpensearchHandlerMockRecorder) RoleMappingUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleMappingUpdate", reflect.TypeOf((*MockOpensearchHandler)(nil).RoleMappingUpdate), arg0, arg1)
}

// RoleUpdate mocks base method.
func (m *MockOpensearchHandler) RoleUpdate(arg0 string, arg1 *opensearch.SecurityPutRole) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RoleUpdate indicates an expected call of RoleUpdate.
func (mr *MockOpensearchHandlerMockRecorder) RoleUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleUpdate", reflect.TypeOf((*MockOpensearchHandler)(nil).RoleUpdate), arg0, arg1)
}

// SetLogger mocks base method.
func (m *MockOpensearchHandler) SetLogger(arg0 *logrus.Entry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogger", arg0)
}

// SetLogger indicates an expected call of SetLogger.
func (mr *MockOpensearchHandlerMockRecorder) SetLogger(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogger", reflect.TypeOf((*MockOpensearchHandler)(nil).SetLogger), arg0)
}

// SnapshotRepositoryDelete mocks base method.
func (m *MockOpensearchHandler) SnapshotRepositoryDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotRepositoryDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SnapshotRepositoryDelete indicates an expected call of SnapshotRepositoryDelete.
func (mr *MockOpensearchHandlerMockRecorder) SnapshotRepositoryDelete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotRepositoryDelete", reflect.TypeOf((*MockOpensearchHandler)(nil).SnapshotRepositoryDelete), arg0)
}

// SnapshotRepositoryDiff mocks base method.
func (m *MockOpensearchHandler) SnapshotRepositoryDiff(arg0, arg1, arg2 *opensearch.SnapshotRepositoryMetaData) (*patch.PatchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotRepositoryDiff", arg0, arg1, arg2)
	ret0, _ := ret[0].(*patch.PatchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotRepositoryDiff indicates an expected call of SnapshotRepositoryDiff.
func (mr *MockOpensearchHandlerMockRecorder) SnapshotRepositoryDiff(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotRepositoryDiff", reflect.TypeOf((*MockOpensearchHandler)(nil).SnapshotRepositoryDiff), arg0, arg1, arg2)
}

// SnapshotRepositoryGet mocks base method.
func (m *MockOpensearchHandler) SnapshotRepositoryGet(arg0 string) (*opensearch.SnapshotRepositoryMetaData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotRepositoryGet", arg0)
	ret0, _ := ret[0].(*opensearch.SnapshotRepositoryMetaData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotRepositoryGet indicates an expected call of SnapshotRepositoryGet.
func (mr *MockOpensearchHandlerMockRecorder) SnapshotRepositoryGet(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotRepositoryGet", reflect.TypeOf((*MockOpensearchHandler)(nil).SnapshotRepositoryGet), arg0)
}

// SnapshotRepositoryUpdate mocks base method.
func (m *MockOpensearchHandler) SnapshotRepositoryUpdate(arg0 string, arg1 *opensearch.SnapshotRepositoryMetaData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotRepositoryUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SnapshotRepositoryUpdate indicates an expected call of SnapshotRepositoryUpdate.
func (mr *MockOpensearchHandlerMockRecorder) SnapshotRepositoryUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotRepositoryUpdate", reflect.TypeOf((*MockOpensearchHandler)(nil).SnapshotRepositoryUpdate), arg0, arg1)
}

// TenantDelete mocks base method.
func (m *MockOpensearchHandler) TenantDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TenantDelete indicates an expected call of TenantDelete.
func (mr *MockOpensearchHandlerMockRecorder) TenantDelete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantDelete", reflect.TypeOf((*MockOpensearchHandler)(nil).TenantDelete), arg0)
}

// TenantDiff mocks base method.
func (m *MockOpensearchHandler) TenantDiff(arg0, arg1, arg2 *opensearch.SecurityPutTenant) (*patch.PatchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantDiff", arg0, arg1, arg2)
	ret0, _ := ret[0].(*patch.PatchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TenantDiff indicates an expected call of TenantDiff.
func (mr *MockOpensearchHandlerMockRecorder) TenantDiff(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantDiff", reflect.TypeOf((*MockOpensearchHandler)(nil).TenantDiff), arg0, arg1, arg2)
}

// TenantGet mocks base method.
func (m *MockOpensearchHandler) TenantGet(arg0 string) (*opensearch.SecurityTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantGet", arg0)
	ret0, _ := ret[0].(*opensearch.SecurityTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TenantGet indicates an expected call of TenantGet.
func (mr *MockOpensearchHandlerMockRecorder) TenantGet(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantGet", reflect.TypeOf((*MockOpensearchHandler)(nil).TenantGet), arg0)
}

// TenantUpdate mocks base method.
func (m *MockOpensearchHandler) TenantUpdate(arg0 string, arg1 *opensearch.SecurityPutTenant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TenantUpdate indicates an expected call of TenantUpdate.
func (mr *MockOpensearchHandlerMockRecorder) TenantUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantUpdate", reflect.TypeOf((*MockOpensearchHandler)(nil).TenantUpdate), arg0, arg1)
}

// UserDelete mocks base method.
func (m *MockOpensearchHandler) UserDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserDelete indicates an expected call of UserDelete.
func (mr *MockOpensearchHandlerMockRecorder) UserDelete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserDelete", reflect.TypeOf((*MockOpensearchHandler)(nil).UserDelete), arg0)
}

// UserDiff mocks base method.
func (m *MockOpensearchHandler) UserDiff(arg0, arg1, arg2 *opensearch.SecurityPutUser) (*patch.PatchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserDiff", arg0, arg1, arg2)
	ret0, _ := ret[0].(*patch.PatchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserDiff indicates an expected call of UserDiff.
func (mr *MockOpensearchHandlerMockRecorder) UserDiff(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserDiff", reflect.TypeOf((*MockOpensearchHandler)(nil).UserDiff), arg0, arg1, arg2)
}

// UserGet mocks base method.
func (m *MockOpensearchHandler) UserGet(arg0 string) (*opensearch.SecurityUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserGet", arg0)
	ret0, _ := ret[0].(*opensearch.SecurityUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserGet indicates an expected call of UserGet.
func (mr *MockOpensearchHandlerMockRecorder) UserGet(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserGet", reflect.TypeOf((*MockOpensearchHandler)(nil).UserGet), arg0)
}

// UserUpdate mocks base method.
func (m *MockOpensearchHandler) UserUpdate(arg0 string, arg1 *opensearch.SecurityPutUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserUpdate indicates an expected call of UserUpdate.
func (mr *MockOpensearchHandlerMockRecorder) UserUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserUpdate", reflect.TypeOf((*MockOpensearchHandler)(nil).UserUpdate), arg0, arg1)
}
