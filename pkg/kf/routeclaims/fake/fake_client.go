// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/kf/pkg/kf/routeclaims/fake (interfaces: Client)

// Package fake is a generated GoMock package.
package fake

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/google/kf/pkg/apis/kf/v1alpha1"
	routeclaims "github.com/google/kf/pkg/kf/routeclaims"
	reflect "reflect"
)

// FakeClient is a mock of Client interface
type FakeClient struct {
	ctrl     *gomock.Controller
	recorder *FakeClientMockRecorder
}

// FakeClientMockRecorder is the mock recorder for FakeClient
type FakeClientMockRecorder struct {
	mock *FakeClient
}

// NewFakeClient creates a new mock instance
func NewFakeClient(ctrl *gomock.Controller) *FakeClient {
	mock := &FakeClient{ctrl: ctrl}
	mock.recorder = &FakeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *FakeClient) EXPECT() *FakeClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *FakeClient) Create(arg0 string, arg1 *v1alpha1.RouteClaim, arg2 ...routeclaims.CreateOption) (*v1alpha1.RouteClaim, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*v1alpha1.RouteClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *FakeClientMockRecorder) Create(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*FakeClient)(nil).Create), varargs...)
}

// Delete mocks base method
func (m *FakeClient) Delete(arg0, arg1 string, arg2 ...routeclaims.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *FakeClientMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*FakeClient)(nil).Delete), varargs...)
}

// Get mocks base method
func (m *FakeClient) Get(arg0, arg1 string, arg2 ...routeclaims.GetOption) (*v1alpha1.RouteClaim, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*v1alpha1.RouteClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *FakeClientMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*FakeClient)(nil).Get), varargs...)
}

// List mocks base method
func (m *FakeClient) List(arg0 string, arg1 ...routeclaims.ListOption) ([]v1alpha1.RouteClaim, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]v1alpha1.RouteClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *FakeClientMockRecorder) List(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*FakeClient)(nil).List), varargs...)
}

// Transform mocks base method
func (m *FakeClient) Transform(arg0, arg1 string, arg2 routeclaims.Mutator) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transform", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transform indicates an expected call of Transform
func (mr *FakeClientMockRecorder) Transform(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transform", reflect.TypeOf((*FakeClient)(nil).Transform), arg0, arg1, arg2)
}

// Update mocks base method
func (m *FakeClient) Update(arg0 string, arg1 *v1alpha1.RouteClaim, arg2 ...routeclaims.UpdateOption) (*v1alpha1.RouteClaim, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*v1alpha1.RouteClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *FakeClientMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*FakeClient)(nil).Update), varargs...)
}

// Upsert mocks base method
func (m *FakeClient) Upsert(arg0 string, arg1 *v1alpha1.RouteClaim, arg2 routeclaims.Merger) (*v1alpha1.RouteClaim, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.RouteClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert
func (mr *FakeClientMockRecorder) Upsert(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*FakeClient)(nil).Upsert), arg0, arg1, arg2)
}
