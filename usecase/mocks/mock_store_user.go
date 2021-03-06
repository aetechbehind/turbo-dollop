// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ae-tech-behind/turbo-dollop/usecase (interfaces: StoreUser)

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "github.com/ae-tech-behind/turbo-dollop/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStoreUser is a mock of StoreUser interface
type MockStoreUser struct {
	ctrl     *gomock.Controller
	recorder *MockStoreUserMockRecorder
}

// MockStoreUserMockRecorder is the mock recorder for MockStoreUser
type MockStoreUserMockRecorder struct {
	mock *MockStoreUser
}

// NewMockStoreUser creates a new mock instance
func NewMockStoreUser(ctrl *gomock.Controller) *MockStoreUser {
	mock := &MockStoreUser{ctrl: ctrl}
	mock.recorder = &MockStoreUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStoreUser) EXPECT() *MockStoreUserMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockStoreUser) CreateUser(arg0 entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockStoreUserMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStoreUser)(nil).CreateUser), arg0)
}

// DeleteUser mocks base method
func (m *MockStoreUser) DeleteUser(arg0 string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockStoreUserMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStoreUser)(nil).DeleteUser), arg0)
}

// GetUser mocks base method
func (m *MockStoreUser) GetUser(arg0 string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockStoreUserMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStoreUser)(nil).GetUser), arg0)
}

// GetUsers mocks base method
func (m *MockStoreUser) GetUsers() ([]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockStoreUserMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockStoreUser)(nil).GetUsers))
}

// UpdateUser mocks base method
func (m *MockStoreUser) UpdateUser(arg0 entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockStoreUserMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStoreUser)(nil).UpdateUser), arg0)
}
