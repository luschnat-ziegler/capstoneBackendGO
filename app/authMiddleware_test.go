package app

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"github.com/luschnat-ziegler/cc_backend_go/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupAuthMiddlewareTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockAuthService = service.NewMockAuthService(ctrl)
	am := AuthMiddleware{mockAuthService}

	testHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "handler entered")
	}

	router = mux.NewRouter()
	router.Use(am.authorizationHandler())
	router.HandleFunc("/auth_required/{id}", testHandler).
		Name("GetUser").
		Methods(http.MethodGet)
	router.HandleFunc("/no_auth_required", testHandler).Name("no_auth_required")
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_should_serve_next_if_route_is_not_restricted(t *testing.T) {

	// Arrange
	teardown := setupAuthMiddlewareTest(t)
	defer teardown()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/no_auth_required", nil)
	recorder := httptest.NewRecorder()

	// Act
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Errorf("Wrong code returned. 200 expected, returned: %v", recorder.Code)
	}

	expected := "handler entered"
	if expected != recorder.Body.String() {
		t.Errorf("Wrong body written, expected 'handler returned', got %v", recorder.Body.String())
	}
}

func Test_should_return_401_if_route_is_protected_and_no_token_in_header(t *testing.T) {

	// Arrange
	teardown := setupAuthMiddlewareTest(t)
	defer teardown()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/auth_required/test_id", nil)
	recorder := httptest.NewRecorder()

	// Act
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Wrong code returned. 401 expected, returned: %v", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	expected := `{"message":"Token missing"}`
	if resBody != expected {
		t.Errorf("Wrong body written, expected {\"message\":\"Token missing\"}, got %v", recorder.Body.String())
	}
}

func Test_should_return_401_if_route_is_protected_and_token_is_invalid(t *testing.T) {
	// Arrange
	teardown := setupAuthMiddlewareTest(t)
	defer teardown()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/auth_required/test_id", nil)
	request.Header.Set("Authorization", "test_token")
	recorder := httptest.NewRecorder()

	mockAppError := errs.NewUnauthorizedError("invalid token")
	mockAuthService.EXPECT().Verify("test_token").Return(nil, mockAppError)

	// Act
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Wrong code returned. 401 expected, returned: %v", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	expected := `{"message":"invalid token"}`
	if resBody != expected {
		t.Errorf("Wrong body written, expected {\"message\":\"invalid token\"}, got %v", recorder.Body.String())
	}
}

func Test_should_return_401_if_route_is_protected_and_token_does_not_match_id(t *testing.T) {

	// Arrange
	teardown := setupAuthMiddlewareTest(t)
	defer teardown()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/auth_required/test_id", nil)
	request.Header.Set("Authorization", "test_token")
	recorder := httptest.NewRecorder()

	mockId := "wrong_test_id"
	mockAuthService.EXPECT().Verify("test_token").Return(&mockId, nil)

	// Act
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Wrong code returned. 401 expected, returned: %v", recorder.Code)
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	expected := `{"message":"Token not matching requested user"}`
	if resBody != expected {
		t.Errorf("Wrong body written, expected {\"message\":\"Token not matching requested user\"}, got %v", recorder.Body.String())
	}
}

func Test_should_serve_next_if_route_is_protected_and_auth_is_valid_and_matching(t *testing.T) {
	// Arrange
	teardown := setupAuthMiddlewareTest(t)
	defer teardown()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/auth_required/test_id", nil)
	request.Header.Set("Authorization", "test_token")
	recorder := httptest.NewRecorder()

	mockId := "test_id"
	mockAuthService.EXPECT().Verify("test_token").Return(&mockId, nil)

	// Act
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Errorf("Wrong code returned. 200 expected, returned: %v", recorder.Code)
	}

	expected := "handler entered"
	if expected != recorder.Body.String() {
		t.Errorf("Wrong body written, expected 'handler returned', got %v", recorder.Body.String())
	}
}
