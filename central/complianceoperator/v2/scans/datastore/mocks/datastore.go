// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/datastore.go -source datastore.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	storage "github.com/stackrox/rox/generated/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockDataStore is a mock of DataStore interface.
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore.
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance.
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// DeleteScan mocks base method.
func (m *MockDataStore) DeleteScan(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScan", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteScan indicates an expected call of DeleteScan.
func (mr *MockDataStoreMockRecorder) DeleteScan(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScan", reflect.TypeOf((*MockDataStore)(nil).DeleteScan), ctx, id)
}

// GetScan mocks base method.
func (m *MockDataStore) GetScan(ctx context.Context, id string) (*storage.ComplianceOperatorScanV2, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScan", ctx, id)
	ret0, _ := ret[0].(*storage.ComplianceOperatorScanV2)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetScan indicates an expected call of GetScan.
func (mr *MockDataStoreMockRecorder) GetScan(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScan", reflect.TypeOf((*MockDataStore)(nil).GetScan), ctx, id)
}

// GetScansByCluster mocks base method.
func (m *MockDataStore) GetScansByCluster(ctx context.Context, clusterID string) ([]*storage.ComplianceOperatorScanV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScansByCluster", ctx, clusterID)
	ret0, _ := ret[0].([]*storage.ComplianceOperatorScanV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScansByCluster indicates an expected call of GetScansByCluster.
func (mr *MockDataStoreMockRecorder) GetScansByCluster(ctx, clusterID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScansByCluster", reflect.TypeOf((*MockDataStore)(nil).GetScansByCluster), ctx, clusterID)
}

// UpsertScan mocks base method.
func (m *MockDataStore) UpsertScan(ctx context.Context, result *storage.ComplianceOperatorScanV2) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertScan", ctx, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertScan indicates an expected call of UpsertScan.
func (mr *MockDataStoreMockRecorder) UpsertScan(ctx, result any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertScan", reflect.TypeOf((*MockDataStore)(nil).UpsertScan), ctx, result)
}