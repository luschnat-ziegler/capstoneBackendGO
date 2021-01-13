package app

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"github.com/luschnat-ziegler/cc_backend_go/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func stringPtr(s string) *string {
	return &s
}

var ah AuthHandlers
var mockAuthService *service.MockAuthService
var router *mux.Router

func setupAuthHandlersTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockAuthService = service.NewMockAuthService(ctrl)
	ah = AuthHandlers{mockAuthService}
	router = mux.NewRouter()
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_status_code_200_and_success_response_with_valid_request(t *testing.T) {

	// Arrange
	teardown := setupAuthHandlersTest(t)
	defer teardown()

	router.HandleFunc("/login", ah.logInUser)

	mockRequestDto := dto.LogInRequest{
		Email:    "test@test.de",
		Password: "test",
	}

	mockResponseDto := dto.LogInResponse{
		Success: true,
		Token:   stringPtr("test_token"),
	}

	mockAuthService.EXPECT().LogIn(mockRequestDto).Return(&mockResponseDto, nil)

	jsonBody := []byte(`{"email":"test@test.de","password":"test"}`)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while writing the code")
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	expected := `{"Success":true,"Token":"test_token"}`

	if expected != resBody {
		t.Error("Response body not matching")
	}
}

func Test_should_return_code_400_and_appropriate_error_with_invalid_json_in_request (t *testing.T) {

	// Arrange
	teardown := setupAuthHandlersTest(t)
	defer teardown()

	router.HandleFunc("/login", ah.logInUser)

	jsonBody := []byte(`"email":"test@test.de","password":"test"}`)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusBadRequest {
		t.Error("Failed while writing the code")
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	expected := `{"message":"Body parsing error"}`

	if expected != resBody {
		t.Error("Response body not matching")
	}
}

func Test_should_return_code_422_and_validation_error_message_with_invalid_request (t *testing.T) {

	// Arrange
	teardown := setupAuthHandlersTest(t)
	defer teardown()

	router.HandleFunc("/login", ah.logInUser)

	jsonBody := []byte(`{"email":"test","password":"test"}`)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert

	if recorder.Code != http.StatusUnprocessableEntity {
		t.Error("Failed while writing the code")
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	expected := `{"invalid_fields":["Email"]}`

	if expected != resBody {
		t.Error("Response error message not matching")
	}
}

func Test_should_return_error_with_code_if_service_returns_error (t *testing.T) {

	// Arrange
	teardown := setupAuthHandlersTest(t)
	defer teardown()

	router.HandleFunc("/login", ah.logInUser)

	jsonBody := []byte(`{"email":"test@test.de","password":"test"}`)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))

	mockRequestDto := dto.LogInRequest{
		Email:    "test@test.de",
		Password: "test",
	}

	mockAppError := errs.NewUnexpectedError("unexpected error")
	mockAuthService.EXPECT().LogIn(mockRequestDto).Return(nil, mockAppError)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while writing the code")
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	expected := `{"message":"unexpected error"}`

	if expected != resBody {
		t.Error("Response error message not matching")
	}
}
