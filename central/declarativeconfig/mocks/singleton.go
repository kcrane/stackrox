// Code generated by MockGen. DO NOT EDIT.
// Source: singleton.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	proto "github.com/gogo/protobuf/proto"
	gomock "github.com/golang/mock/gomock"
)

// MockReconciliationErrorReporter is a mock of ReconciliationErrorReporter interface.
type MockReconciliationErrorReporter struct {
	ctrl     *gomock.Controller
	recorder *MockReconciliationErrorReporterMockRecorder
}

// MockReconciliationErrorReporterMockRecorder is the mock recorder for MockReconciliationErrorReporter.
type MockReconciliationErrorReporterMockRecorder struct {
	mock *MockReconciliationErrorReporter
}

// NewMockReconciliationErrorReporter creates a new mock instance.
func NewMockReconciliationErrorReporter(ctrl *gomock.Controller) *MockReconciliationErrorReporter {
	mock := &MockReconciliationErrorReporter{ctrl: ctrl}
	mock.recorder = &MockReconciliationErrorReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReconciliationErrorReporter) EXPECT() *MockReconciliationErrorReporterMockRecorder {
	return m.recorder
}

// ProcessError mocks base method.
func (m *MockReconciliationErrorReporter) ProcessError(protoValue proto.Message, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProcessError", protoValue, err)
}

// ProcessError indicates an expected call of ProcessError.
func (mr *MockReconciliationErrorReporterMockRecorder) ProcessError(protoValue, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessError", reflect.TypeOf((*MockReconciliationErrorReporter)(nil).ProcessError), protoValue, err)
}