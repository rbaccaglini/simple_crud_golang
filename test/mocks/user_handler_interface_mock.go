// Code generated by MockGen. DO NOT EDIT.
// Source: internal/handlers/user/user_handler_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockUserHandlerInterface is a mock of UserHandlerInterface interface.
type MockUserHandlerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserHandlerInterfaceMockRecorder
}

// MockUserHandlerInterfaceMockRecorder is the mock recorder for MockUserHandlerInterface.
type MockUserHandlerInterfaceMockRecorder struct {
	mock *MockUserHandlerInterface
}

// NewMockUserHandlerInterface creates a new mock instance.
func NewMockUserHandlerInterface(ctrl *gomock.Controller) *MockUserHandlerInterface {
	mock := &MockUserHandlerInterface{ctrl: ctrl}
	mock.recorder = &MockUserHandlerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserHandlerInterface) EXPECT() *MockUserHandlerInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserHandlerInterface) CreateUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateUser", c)
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserHandlerInterfaceMockRecorder) CreateUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserHandlerInterface)(nil).CreateUser), c)
}

// DeleteUser mocks base method.
func (m *MockUserHandlerInterface) DeleteUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUser", c)
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserHandlerInterfaceMockRecorder) DeleteUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserHandlerInterface)(nil).DeleteUser), c)
}

// FindAllUser mocks base method.
func (m *MockUserHandlerInterface) FindAllUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindAllUser", c)
}

// FindAllUser indicates an expected call of FindAllUser.
func (mr *MockUserHandlerInterfaceMockRecorder) FindAllUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllUser", reflect.TypeOf((*MockUserHandlerInterface)(nil).FindAllUser), c)
}

// FindUserByEmail mocks base method.
func (m *MockUserHandlerInterface) FindUserByEmail(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindUserByEmail", c)
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserHandlerInterfaceMockRecorder) FindUserByEmail(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserHandlerInterface)(nil).FindUserByEmail), c)
}

// FindUserById mocks base method.
func (m *MockUserHandlerInterface) FindUserById(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindUserById", c)
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockUserHandlerInterfaceMockRecorder) FindUserById(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockUserHandlerInterface)(nil).FindUserById), c)
}

// Login mocks base method.
func (m *MockUserHandlerInterface) Login(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Login", c)
}

// Login indicates an expected call of Login.
func (mr *MockUserHandlerInterfaceMockRecorder) Login(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserHandlerInterface)(nil).Login), c)
}

// UpdateUser mocks base method.
func (m *MockUserHandlerInterface) UpdateUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateUser", c)
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserHandlerInterfaceMockRecorder) UpdateUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserHandlerInterface)(nil).UpdateUser), c)
}
