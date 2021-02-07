// Code generated by MockGen. DO NOT EDIT.
// Source: ctl/project/container_api.go

// Package project_test is a generated GoMock package.
package project_test

import (
	context "context"
	container "github.com/foldsh/fold/ctl/container"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockContainerAPI is a mock of ContainerAPI interface
type MockContainerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockContainerAPIMockRecorder
}

// MockContainerAPIMockRecorder is the mock recorder for MockContainerAPI
type MockContainerAPIMockRecorder struct {
	mock *MockContainerAPI
}

// NewMockContainerAPI creates a new mock instance
func NewMockContainerAPI(ctrl *gomock.Controller) *MockContainerAPI {
	mock := &MockContainerAPI{ctrl: ctrl}
	mock.recorder = &MockContainerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContainerAPI) EXPECT() *MockContainerAPIMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockContainerAPI) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockContainerAPIMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockContainerAPI)(nil).Context))
}

// NewNetwork mocks base method
func (m *MockContainerAPI) NewNetwork(name string) *container.Network {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewNetwork", name)
	ret0, _ := ret[0].(*container.Network)
	return ret0
}

// NewNetwork indicates an expected call of NewNetwork
func (mr *MockContainerAPIMockRecorder) NewNetwork(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewNetwork", reflect.TypeOf((*MockContainerAPI)(nil).NewNetwork), name)
}

// NetworkExists mocks base method
func (m *MockContainerAPI) NetworkExists(net *container.Network) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkExists", net)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkExists indicates an expected call of NetworkExists
func (mr *MockContainerAPIMockRecorder) NetworkExists(net interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkExists", reflect.TypeOf((*MockContainerAPI)(nil).NetworkExists), net)
}

// CreateNetwork mocks base method
func (m *MockContainerAPI) CreateNetwork(net *container.Network) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNetwork", net)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNetwork indicates an expected call of CreateNetwork
func (mr *MockContainerAPIMockRecorder) CreateNetwork(net interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNetwork", reflect.TypeOf((*MockContainerAPI)(nil).CreateNetwork), net)
}

// RemoveNetwork mocks base method
func (m *MockContainerAPI) RemoveNetwork(net *container.Network) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveNetwork", net)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNetwork indicates an expected call of RemoveNetwork
func (mr *MockContainerAPIMockRecorder) RemoveNetwork(net interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNetwork", reflect.TypeOf((*MockContainerAPI)(nil).RemoveNetwork), net)
}

// NewContainer mocks base method
func (m *MockContainerAPI) NewContainer(name string, image container.Image, mounts ...container.Mount) *container.Container {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, image}
	for _, a := range mounts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewContainer", varargs...)
	ret0, _ := ret[0].(*container.Container)
	return ret0
}

// NewContainer indicates an expected call of NewContainer
func (mr *MockContainerAPIMockRecorder) NewContainer(name, image interface{}, mounts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, image}, mounts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewContainer", reflect.TypeOf((*MockContainerAPI)(nil).NewContainer), varargs...)
}

// GetContainer mocks base method
func (m *MockContainerAPI) GetContainer(name string) (*container.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainer", name)
	ret0, _ := ret[0].(*container.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainer indicates an expected call of GetContainer
func (mr *MockContainerAPIMockRecorder) GetContainer(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainer", reflect.TypeOf((*MockContainerAPI)(nil).GetContainer), name)
}

// RunContainer mocks base method
func (m *MockContainerAPI) RunContainer(con *container.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunContainer", con)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunContainer indicates an expected call of RunContainer
func (mr *MockContainerAPIMockRecorder) RunContainer(con interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunContainer", reflect.TypeOf((*MockContainerAPI)(nil).RunContainer), con)
}

// StopContainer mocks base method
func (m *MockContainerAPI) StopContainer(con *container.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopContainer", con)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopContainer indicates an expected call of StopContainer
func (mr *MockContainerAPIMockRecorder) StopContainer(con interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopContainer", reflect.TypeOf((*MockContainerAPI)(nil).StopContainer), con)
}

// RemoveContainer mocks base method
func (m *MockContainerAPI) RemoveContainer(con *container.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveContainer", con)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContainer indicates an expected call of RemoveContainer
func (mr *MockContainerAPIMockRecorder) RemoveContainer(con interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveContainer", reflect.TypeOf((*MockContainerAPI)(nil).RemoveContainer), con)
}

// AddToNetwork mocks base method
func (m *MockContainerAPI) AddToNetwork(n *container.Network, con *container.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToNetwork", n, con)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToNetwork indicates an expected call of AddToNetwork
func (mr *MockContainerAPIMockRecorder) AddToNetwork(n, con interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToNetwork", reflect.TypeOf((*MockContainerAPI)(nil).AddToNetwork), n, con)
}
