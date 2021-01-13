package service

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	domain2 "github.com/luschnat-ziegler/cc_backend_go/domain"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/mocks/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

var authService AuthService

func setupAuthServiceTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockUserRepository = domain.NewMockUserRepository(ctrl)
	authService = NewAuthService(mockUserRepository)
	return func() {
		userService = nil
		defer ctrl.Finish()
	}
}

func Test_Login_should_return_LogInResponse_and_nil_if_repo_method_returns_correctly (t *testing.T) {

	teardown := setupAuthServiceTest(t)
	defer teardown()

	mockUser := domain2.User{
		ID:                primitive.ObjectID{},
		Email:             "test@test.de",
		Password:          "$2a$14$NaDK2oITrjL0LxemjMD8mu.QB8Kzp2dcKAW0CoE3FhSbV5ycDveTq",
		FirstName:         "Testy",
		LastName:          "McTestface",
		WeightEnvironment: 1,
		WeightGender:      2,
		WeightLgbtq:       3,
		WeightEquality:    4,
		WeightCorruption:  0,
		WeightFreedom:     2,
	}

	mockLogInRequest := dto.LogInRequest{
		Email:    "test@test.de",
		Password: "password",
	}

	mockUserRepository.EXPECT().ByEmail("test@test.de").Return(&mockUser, nil)

	// Act
	result, err := authService.LogIn(mockLogInRequest)

	// Assert
	if err != nil {
		t.Error("Error returned, nil expected.")
	}

	resultType := fmt.Sprintf("%T", result)
	if resultType != "*dto.LogInResponse" {
		t.Error("Wrong type returned")
	}

	if result != nil && result.Success != true {
		t.Error("Returned response Success message does not match. False returned, true expected")
	}
}

func Test_Login_should_return_nil_and_NotFoundError_if_password_does_not_match_hash (t *testing.T) {

	// Arrange

	teardown := setupAuthServiceTest(t)
	defer teardown()

	mockUser := domain2.User{
		ID:                primitive.ObjectID{},
		Email:             "test@test.de",
		Password:          "$2a$14$NaDK2oITrjL0LxemjMD8mu.QB8Kzp2dcKAW0CoE3FhSbV5ycDveTq",
		FirstName:         "Testy",
		LastName:          "McTestface",
		WeightEnvironment: 1,
		WeightGender:      2,
		WeightLgbtq:       3,
		WeightEquality:    4,
		WeightCorruption:  0,
		WeightFreedom:     2,
	}

	mockLogInRequest := dto.LogInRequest{
		Email:    "test@test.de",
		Password: "invalid_password",
	}

	mockUserRepository.EXPECT().ByEmail("test@test.de").Return(&mockUser, nil)

	// Act
	result, err := authService.LogIn(mockLogInRequest)

	// Assert
	if result != nil {
		t.Error("Result returned, nil expected")
	}

	if err == nil {
		t.Error("Nil returned, appError expected.")
	}

	appErrorType := fmt.Sprintf("%T", err)
	if appErrorType != "*errs.AppError" {
		t.Error("Wrong type returned")
	}

	appErrorString := fmt.Sprintf("%v", err)
	if err != nil && appErrorString != "&{404 Wrong Password}" {
		t.Error("returned error does not match")
	}
}

func Test_Login_should_return_nil_and_AppError_if_repo_method_returns_error (t *testing.T) {

	// Arrange
	teardown := setupAuthServiceTest(t)
	defer teardown()

	mockLogInRequest := dto.LogInRequest{
		Email:    "test@test.de",
		Password: "invalid_password",
	}

	mockError := errors.New("test")
	mockUserRepository.EXPECT().ByEmail("test@test.de").Return(nil, &mockError)

	// Act
	result, err := authService.LogIn(mockLogInRequest)

	// Assert
	if result != nil {
		t.Error("Result returned, nil expected")
	}

	if err == nil {
		t.Error("Nil returned, appError expected.")
	}

	appErrorType := fmt.Sprintf("%T", err)
	if appErrorType != "*errs.AppError" {
		t.Error("Wrong type returned")
	}

	appErrorString := fmt.Sprintf("%v", err)
	if err != nil && appErrorString != "&{404 User not found}" {
		t.Error("returned error does not match")
	}
}