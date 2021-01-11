package service

import (
	"fmt"
	"github.com/golang/mock/gomock"
	domain2 "github.com/luschnat-ziegler/cc_backend_go/domain"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"github.com/luschnat-ziegler/cc_backend_go/mocks/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

// ##########
// CreateUser
// ##########

func Test_CreateUser_returns_CreateUserResponse_and_nil_when_Repo_method_returns_no_error (t *testing.T) {
	
	// Arrange
	ctrl := gomock.NewController(t)
	mockRepository := domain.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepository)
	
	mockCreateUserRequest := dto.CreateUserRequest{
		Email:     "test@test.de",
		Password:  "password",
		FirstName: "Testy",
		LastName:  "McTestface",
	}

	var user = domain2.NewUser(mockCreateUserRequest)

	id := "test_id"
	mockRepository.EXPECT().Save(user).Return(&id, nil)

	//Act
	result, err := service.CreateUser(mockCreateUserRequest)

	// Assert
	if err != nil {
		t.Error("Error returned, nil expected.")
	}

	if result.Id != id {
		t.Error("Wrong return value")
	}

	resultType := fmt.Sprintf("%T", result)
	if resultType != "*dto.CreateUserResponse" {
		t.Error("Wrong type returned")
	}
}

func Test_CreateUser_returns_nil_and_AppError_if_repo_method_returns_error (t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)
	mockRepository := domain.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepository)

	mockCreateUserRequest := dto.CreateUserRequest{
		Email:     "test@test.de",
		Password:  "password",
		FirstName: "Testy",
		LastName:  "McTestface",
	}

	var user = domain2.NewUser(mockCreateUserRequest)

	mockAppError := errs.NewUnexpectedError("unexpected server error")
	mockRepository.EXPECT().Save(user).Return(nil, mockAppError)

	// Act
	result, err := service.CreateUser(mockCreateUserRequest)

	// Assert
	if err == nil {
		t.Error("Nil returned, *AppError expected")
	}

	if result != nil {
		t.Error("Result returned, nil expected")
	}

	errorType := fmt.Sprintf("%T", err)
	if errorType != "*errs.AppError" {
		t.Error("Wrong type returned")
	}

	if err != nil && (err.Code != 500 || err.Message != "unexpected server error") {
		t.Error("Wrong error code and/or message returned")
	}
}

// #######
// GetUser
// #######

func Test_GetUser_returns_GetUserResponse_and_nil_if_repo_method_returns_no_error (t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)
	mockRepository := domain.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepository)

	user := domain2.User{
		ID:                primitive.ObjectID{},
		Email:             "test@test.de",
		Password:          "password",
		FirstName:         "Testy",
		LastName:          "McTestface",
		WeightEnvironment: 1,
		WeightGender:      2,
		WeightLgbtq:       3,
		WeightEquality:    0,
		WeightCorruption:  4,
		WeightFreedom:     1,
	}

	mockRepository.EXPECT().ById("test_id").Return(&user, nil)

	// Act
	result, err := service.GetUser("test_id")

	// Assert
	if err != nil {
		t.Error("Error returned, nil expected.")
	}

	resultType := fmt.Sprintf("%T", result)
	if resultType != "*dto.GetUserResponse" {
		t.Error("Wrong type returned")
	}

	resultString := fmt.Sprintf("%v", result)
	if resultString != "&{test@test.de Testy McTestface 1 2 3 0 4 1}" {
		t.Error("Returned GetUserResponse does not match input data")
	}
}

func Test_GetUser_returns_nil_and_AppError_if_repo_method_returns_error (t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)
	mockRepository := domain.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepository)

	mockAppError := errs.NewUnexpectedError("unexpected server error")
	mockRepository.EXPECT().ById("test_id").Return(nil, mockAppError)

	// Act
	result, err := service.GetUser("test_id")

	// Assert
	if err == nil {
		t.Error("Nil returned, *AppError expected")
	}

	if result != nil {
		t.Error("Result returned, nil expected")
	}

	errorType := fmt.Sprintf("%T", err)
	if errorType != "*errs.AppError" {
		t.Error("Wrong type returned")
	}

	if err != nil && (err.Code != 500 || err.Message != "unexpected server error") {
		t.Error("Wrong error code and/or message returned")
	}
}

// #############
// UpdateWeights
// #############

