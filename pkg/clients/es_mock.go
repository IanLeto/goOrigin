// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/clients/es.go

// Package clients is a generated GoMock package.
package clients

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEsAPIs is a mock of EsAPIs interface
type MockEsAPIs struct {
	ctrl     *gomock.Controller
	recorder *MockEsAPIsMockRecorder
}

// MockEsAPIsMockRecorder is the mock recorder for MockEsAPIs
type MockEsAPIsMockRecorder struct {
	mock *MockEsAPIs
}

// NewMockEsAPIs creates a new mock instance
func NewMockEsAPIs(ctrl *gomock.Controller) *MockEsAPIs {
	mock := &MockEsAPIs{ctrl: ctrl}
	mock.recorder = &MockEsAPIsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEsAPIs) EXPECT() *MockEsAPIsMockRecorder {
	return m.recorder
}

// Ping mocks base method
func (m *MockEsAPIs) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockEsAPIsMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockEsAPIs)(nil).Ping))
}
