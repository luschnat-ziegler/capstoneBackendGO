// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luschnat-ziegler/cc_backend_go/service (interfaces: UserService)

// Package service is a generated GoMock package.
package service

import (
	gomock "github.com/golang/mock/gomock"
	dto "github.com/luschnat-ziegler/cc_backend_go/dto"
	errs "github.com/luschnat-ziegler/cc_backend_go/errs"
	reflect "reflect"
)

// MockUserService is a mock of UserService interface
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockUserService) CreateUser(arg0 dto.CreateUserRequest) (*dto.CreateUserResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(*dto.CreateUserResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockUserServiceMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserService)(nil).CreateUser), arg0)
}

// GetUser mocks base method
func (m *MockUserService) GetUser(arg0 string) (*dto.GetUserResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*dto.GetUserResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockUserServiceMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserService)(nil).GetUser), arg0)
}

// UpdateWeights mocks base method
func (m *MockUserService) UpdateWeights(arg0 dto.SetUserWeightsRequest) (*dto.SetUserWeightsResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWeights", arg0)
	ret0, _ := ret[0].(*dto.SetUserWeightsResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// UpdateWeights indicates an expected call of UpdateWeights
func (mr *MockUserServiceMockRecorder) UpdateWeights(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWeights", reflect.TypeOf((*MockUserService)(nil).UpdateWeights), arg0)
}
