// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GoogleCloudPlatform/kf/pkg/kf/fake (interfaces: LogTailer)

// Package fake is a generated GoMock package.
package fake

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// FakeLogTailer is a mock of LogTailer interface
type FakeLogTailer struct {
	ctrl     *gomock.Controller
	recorder *FakeLogTailerMockRecorder
}

// FakeLogTailerMockRecorder is the mock recorder for FakeLogTailer
type FakeLogTailerMockRecorder struct {
	mock *FakeLogTailer
}

// NewFakeLogTailer creates a new mock instance
func NewFakeLogTailer(ctrl *gomock.Controller) *FakeLogTailer {
	mock := &FakeLogTailer{ctrl: ctrl}
	mock.recorder = &FakeLogTailerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *FakeLogTailer) EXPECT() *FakeLogTailerMockRecorder {
	return m.recorder
}

// Tail mocks base method
func (m *FakeLogTailer) Tail(arg0 io.Writer, arg1, arg2 string, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tail", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Tail indicates an expected call of Tail
func (mr *FakeLogTailerMockRecorder) Tail(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tail", reflect.TypeOf((*FakeLogTailer)(nil).Tail), arg0, arg1, arg2, arg3)
}