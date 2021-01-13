// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luschnat-ziegler/cc_backend_go/service (interfaces: AuthService)

// Package service is a generated GoMock package.
package service

import (
	gomock "github.com/golang/mock/gomock"
	dto "github.com/luschnat-ziegler/cc_backend_go/dto"
	errs "github.com/luschnat-ziegler/cc_backend_go/errs"
	reflect "reflect"
)

// MockAuthService is a mock of AuthService interface
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// LogIn mocks base method
func (m *MockAuthService) LogIn(arg0 dto.LogInRequest) (*dto.LogInResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogIn", arg0)
	ret0, _ := ret[0].(*dto.LogInResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// LogIn indicates an expected call of LogIn
func (mr *MockAuthServiceMockRecorder) LogIn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogIn", reflect.TypeOf((*MockAuthService)(nil).LogIn), arg0)
}

// Verify mocks base method
func (m *MockAuthService) Verify(arg0 string) (*string, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// Verify indicates an expected call of Verify
func (mr *MockAuthServiceMockRecorder) Verify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockAuthService)(nil).Verify), arg0)
}
