// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eiizu/go-service/usecase (interfaces: StoreBook)

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "github.com/ae-tech-behind/turbo-dollop/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStoreBook is a mock of StoreBook interface
type MockStoreBook struct {
	ctrl     *gomock.Controller
	recorder *MockStoreBookMockRecorder
}

// MockStoreBookMockRecorder is the mock recorder for MockStoreBook
type MockStoreBookMockRecorder struct {
	mock *MockStoreBook
}

// NewMockStoreBook creates a new mock instance
func NewMockStoreBook(ctrl *gomock.Controller) *MockStoreBook {
	mock := &MockStoreBook{ctrl: ctrl}
	mock.recorder = &MockStoreBookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStoreBook) EXPECT() *MockStoreBookMockRecorder {
	return m.recorder
}

// CreateBook mocks base method
func (m *MockStoreBook) CreateBook(arg0 entity.Book) (*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", arg0)
	ret0, _ := ret[0].(*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook
func (mr *MockStoreBookMockRecorder) CreateBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockStoreBook)(nil).CreateBook), arg0)
}

// DeleteBook mocks base method
func (m *MockStoreBook) DeleteBook(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook
func (mr *MockStoreBookMockRecorder) DeleteBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockStoreBook)(nil).DeleteBook), arg0)
}

// GetBook mocks base method
func (m *MockStoreBook) GetBook(arg0 string) (*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBook", arg0)
	ret0, _ := ret[0].(*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBook indicates an expected call of GetBook
func (mr *MockStoreBookMockRecorder) GetBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBook", reflect.TypeOf((*MockStoreBook)(nil).GetBook), arg0)
}

// GetBooks mocks base method
func (m *MockStoreBook) GetBooks() ([]entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooks")
	ret0, _ := ret[0].([]entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooks indicates an expected call of GetBooks
func (mr *MockStoreBookMockRecorder) GetBooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooks", reflect.TypeOf((*MockStoreBook)(nil).GetBooks))
}

// UpdateBook mocks base method
func (m *MockStoreBook) UpdateBook(arg0 string, arg1 entity.Book) (*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", arg0, arg1)
	ret0, _ := ret[0].(*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBook indicates an expected call of UpdateBook
func (mr *MockStoreBookMockRecorder) UpdateBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockStoreBook)(nil).UpdateBook), arg0, arg1)
}
