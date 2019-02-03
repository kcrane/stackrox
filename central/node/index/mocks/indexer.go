// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/node/index (interfaces: Indexer)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	reflect "reflect"
)

// MockIndexer is a mock of Indexer interface
type MockIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockIndexerMockRecorder
}

// MockIndexerMockRecorder is the mock recorder for MockIndexer
type MockIndexerMockRecorder struct {
	mock *MockIndexer
}

// NewMockIndexer creates a new mock instance
func NewMockIndexer(ctrl *gomock.Controller) *MockIndexer {
	mock := &MockIndexer{ctrl: ctrl}
	mock.recorder = &MockIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexer) EXPECT() *MockIndexerMockRecorder {
	return m.recorder
}

// AddNode mocks base method
func (m *MockIndexer) AddNode(arg0 *storage.Node) error {
	ret := m.ctrl.Call(m, "AddNode", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNode indicates an expected call of AddNode
func (mr *MockIndexerMockRecorder) AddNode(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNode", reflect.TypeOf((*MockIndexer)(nil).AddNode), arg0)
}

// AddNodes mocks base method
func (m *MockIndexer) AddNodes(arg0 []*storage.Node) error {
	ret := m.ctrl.Call(m, "AddNodes", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNodes indicates an expected call of AddNodes
func (mr *MockIndexerMockRecorder) AddNodes(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNodes", reflect.TypeOf((*MockIndexer)(nil).AddNodes), arg0)
}

// DeleteNode mocks base method
func (m *MockIndexer) DeleteNode(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteNode", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNode indicates an expected call of DeleteNode
func (mr *MockIndexerMockRecorder) DeleteNode(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNode", reflect.TypeOf((*MockIndexer)(nil).DeleteNode), arg0)
}

// Search mocks base method
func (m *MockIndexer) Search(arg0 *v1.Query) ([]search.Result, error) {
	ret := m.ctrl.Call(m, "Search", arg0)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockIndexerMockRecorder) Search(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockIndexer)(nil).Search), arg0)
}
