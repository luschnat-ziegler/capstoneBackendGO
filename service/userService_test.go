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

var mockUserRepository *domain.MockUserRepository
var userService UserService

func setupUserServiceTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockUserRepository = domain.NewMockUserRepository(ctrl)
	userService = NewUserService(mockUserRepository)
	return func() {
		userService = nil
		defer ctrl.Finish()
	}
}

// ##########
// CreateUser
// ##########

func Test_CreateUser_returns_CreateUserResponse_and_nil_when_Repo_method_returns_no_error(t *testing.T) {

	// Arrange
	teardown := setupUserServiceTest(t)
	defer teardown()

	mockCreateUserRequest := dto.CreateUserRequest{
		Email:     "test@test.de",
		Password:  "password",
		FirstName: "Testy",
		LastName:  "McTestface",
	}

	var user = domain2.NewUser(mockCreateUserRequest)

	id := "test_id"
	mockUserRepository.EXPECT().Save(user).Return(&id, nil)

	//Act
	result, err := userService.CreateUser(mockCreateUserRequest)

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

func Test_CreateUser_returns_nil_and_AppError_if_repo_method_returns_error(t *testing.T) {

	// Arrange
	teardown := setupUserServiceTest(t)
	defer teardown()

	mockCreateUserRequest := dto.CreateUserRequest{
		Email:     "test@test.de",
		Password:  "password",
		FirstName: "Testy",
		LastName:  "McTestface",
	}

	var user = domain2.NewUser(mockCreateUserRequest)

	mockAppError := errs.NewUnexpectedError("unexpected server error")
	mockUserRepository.EXPECT().Save(user).Return(nil, mockAppError)

	// Act
	result, err := userService.CreateUser(mockCreateUserRequest)

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

func Test_GetUser_should_return_GetUserRequest_and_nil_called_with_valid_id(t *testing.T) {

	// Arrange

	teardown := setupUserServiceTest(t)
	defer teardown()

	user := domain2.User{
		ID:                primitive.ObjectID{},
		Email:             "test@test.de",
		Password:          "password",
		FirstName:         "Testy",
		LastName:          "McTestface",
		WeightEnvironment: 2,
		WeightGender:      2,
		WeightLgbtq:       1,
		WeightEquality:    3,
		WeightCorruption:  4,
		WeightFreedom:     0,
	}

	mockUserRepository.EXPECT().ById("test_id").Return(&user, nil)

	// Act
	result, err := userService.GetUser("test_id")

	// Assert
	if err != nil {
		t.Error("Error returned, nil expected.")
	}

	resultType := fmt.Sprintf("%T", result)
	if resultType != "*dto.GetUserResponse" {
		t.Error("Wrong type returned")
	}

	resultString := fmt.Sprintf("%v", result)
	expected := "&{test@test.de Testy McTestface 2 2 1 3 4 0}"
	if result != nil && resultString != expected {
		t.Error("Returned value does not match")
	}
}

func Test_GetUser_should_return_nil_and_AppError_if_repo_method_returns_error(t *testing.T) {

	// Arrange
	teardown := setupUserServiceTest(t)
	defer teardown()

	mockAppError := errs.NewUnexpectedError("unexpected server error")

	mockUserRepository.EXPECT().ById("test_id").Return(nil, mockAppError)

	// Act
	result, err := userService.GetUser("test_id")

	// Assert
	if result != nil {
		t.Error("Result not nil, nil expected")
	}

	if err == nil {
		t.Error("Error nil, *AppError expected")
	}

	errorString := fmt.Sprintf("%v", result)
	expected := "&{500 unexpected server error}"
	if result != nil && errorString != expected {
		t.Error("Returned error does not match")
	}
}

// #############
// UpdateWeights
// #############

func Test_UpdateWeights_should_return_SetUserWeightsResponse_and_nil_called_with_SetUserWeightsRequest(t *testing.T) {

	// Arrange
	teardown := setupUserServiceTest(t)
	defer teardown()

	mockSetUserWeightsRequest := dto.SetUserWeightsRequest{
		Id:                "test_id",
		WeightEnvironment: intPtr(1),
		WeightGender:      intPtr(2),
		WeightLgbtq:       intPtr(3),
		WeightEquality:    intPtr(0),
		WeightCorruption:  intPtr(4),
		WeightFreedom:     intPtr(2),
	}

	mockSetUserWeightsResponse := dto.SetUserWeightsResponse{
		Matched: true,
		Updated: true,
	}

	mockUserRepository.EXPECT().UpdateWeights(mockSetUserWeightsRequest).Return(&mockSetUserWeightsResponse, nil)

	// Act
	result, err := userService.UpdateWeights(mockSetUserWeightsRequest)

	// Assert
	if err != nil {
		t.Error("Error returned, nil expected.")
	}

	resultType := fmt.Sprintf("%T", result)
	if resultType != "*dto.SetUserWeightsResponse" {
		t.Error("Wrong type returned")
	}

	resultString := fmt.Sprintf("%v", result)
	expected := "&{true true}"
	if result != nil && resultString != expected {
		t.Error("Returned value does not match")
	}
}

func Test_UpdateWeights_should_return_nil_and_AppError_if_repo_method_returns_error(t *testing.T) {

	// Arrange
	teardown := setupUserServiceTest(t)
	defer teardown()

	mockSetUserWeightsRequest := dto.SetUserWeightsRequest{
		Id:                "test_id",
		WeightEnvironment: intPtr(1),
		WeightGender:      intPtr(2),
		WeightLgbtq:       intPtr(3),
		WeightEquality:    intPtr(0),
		WeightCorruption:  intPtr(4),
		WeightFreedom:     intPtr(2),
	}

	mockAppError := errs.NewUnexpectedError("unexpected server error")

	mockUserRepository.EXPECT().UpdateWeights(mockSetUserWeightsRequest).Return(nil, mockAppError)

	// Act
	result, err := userService.UpdateWeights(mockSetUserWeightsRequest)

	// Assert
	if result != nil {
		t.Error("Result not nil, nil expected")
	}

	if err == nil {
		t.Error("Error nil, *AppError expected")
	}

	errorString := fmt.Sprintf("%v", result)
	expected := "&{500 unexpected server error}"
	if result != nil && errorString != expected {
		t.Error("Returned error does not match")
	}
}

func intPtr(i int) *int {
	return &i
}
