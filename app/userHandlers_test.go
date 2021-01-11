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

// ##############
// Get user by id
// ##############

func Test_getUserById_should_return_code_200_and_corresponding_response_with_valid_request (t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)
	mockService := service.NewMockUserService(ctrl)
	uh := UserHandlers{mockService}
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", uh.GetUserById)

	mockResponseDto := dto.GetUserResponse{
		Email:             "test@test.de",
		FirstName:         "Test",
		LastName:          "McTestface",
		WeightEnvironment: 2,
		WeightGender:      3,
		WeightLgbtq:       4,
		WeightEquality:    1,
		WeightCorruption:  0,
		WeightFreedom:     2,
	}

	mockService.EXPECT().GetUser("5feb5e9b3cf54dae7d98873e").Return(&mockResponseDto, nil)
	request, _ := http.NewRequest(http.MethodGet, "/user/5feb5e9b3cf54dae7d98873e", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while writing the code")
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]

	expected := `{"email":"test@test.de","first_name":"Test","last_name":"McTestface","weight_environment":2,"weight_gender":3,"WeightLgbtq":4,"weight_equality":1,"weight_corruption":0,"weight_freedom":2}`
	if expected != resBody {
		t.Error("Response body not matching")
	}
}

func Test_getUserById_should_return_error_code_and_error_message_response_when_service_method_returns_error (t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mockService := service.NewMockUserService(ctrl)
	uh := UserHandlers{mockService}
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", uh.GetUserById)

	mockService.EXPECT().GetUser("invalid_id").Return(nil, errs.NewUnexpectedError("unexpected server error"))
	request, _ := http.NewRequest(http.MethodGet, "/user/invalid_id", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while writing the code")
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]

	expected := `{"message":"unexpected server error"}`
	if expected != resBody {
		t.Error("Response body not matching")
	}
}

// ###########
// Create User
// ###########

func Test_createUser_should_return_code_201_and_corresponding_response_with_valid_request (t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)
	mockService := service.NewMockUserService(ctrl)
	uh := UserHandlers{mockService}
	router := mux.NewRouter()
	router.HandleFunc("/user", uh.CreateUser).Methods(http.MethodPost)

	mockRequestDto := dto.CreateUserRequest{
		Email:     "test@test.de",
		Password:  "password",
		FirstName: "Testy",
		LastName:  "McTestface",
	}

	mockResponseDto := dto.CreateUserResponse{Id: "returned_id"}

	jsonBody := []byte(`{"email":"test@test.de","password":"password","first_name":"Testy","last_name":"McTestface"}`)
	mockService.EXPECT().CreateUser(mockRequestDto).Return(&mockResponseDto, nil)
	request, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(jsonBody))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusCreated {
		t.Error("Failed while writing the code")
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	expected := `{"user_id":"returned_id"}`

	if expected != resBody {
		t.Error("Response body not matching")
	}
}

func Test_createUser_should_return_code_400_and_error_message_with_invalid_json_body (t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)
	mockService := service.NewMockUserService(ctrl)
	uh := UserHandlers{mockService}
	router := mux.NewRouter()
	router.HandleFunc("/user", uh.CreateUser).Methods(http.MethodPost)

	jsonBody := []byte(`"email":"test@test.de","password":"password","first_name":"Testy","last_name":"McTestface"}`)
	request, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(jsonBody))

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

func Test_createUser_should_return_code_422_and_validation_message_with_invalid_request (t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)
	mockService := service.NewMockUserService(ctrl)
	uh := UserHandlers{mockService}
	router := mux.NewRouter()
	router.HandleFunc("/user", uh.CreateUser).Methods(http.MethodPost)

	jsonBody := []byte(`{"email":"test@test.de","password":"pass","first_name":"Testy","last_name":"McTestface"}`)
	request, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(jsonBody))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusUnprocessableEntity {
		t.Error("Failed while writing the code")
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	expected := `{"invalid_fields":["Password"]}`

	if expected != resBody {
		t.Error("Response body not matching")
	}
}

func Test_createUser_should_return_error_with_code_if_service_returns_error (t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)
	mockService := service.NewMockUserService(ctrl)
	uh := UserHandlers{mockService}
	router := mux.NewRouter()
	router.HandleFunc("/user", uh.CreateUser).Methods(http.MethodPost)

	mockRequestDto := dto.CreateUserRequest{
		Email:     "test@test.de",
		Password:  "password",
		FirstName: "Testy",
		LastName:  "McTestface",
	}

	jsonBody := []byte(`{"email":"test@test.de","password":"password","first_name":"Testy","last_name":"McTestface"}`)
	mockService.EXPECT().CreateUser(mockRequestDto).Return(nil, errs.NewUnexpectedError("unexpected server error"))
	request, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(jsonBody))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while writing the code")
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	expected := `{"message":"unexpected server error"}`

	if expected != resBody {
		t.Error("Response body not matching")
	}
}

// #################
// updateUserWeights
// #################

func Test_updateUserWeights_should_return_code_200_and_corresponding_response_with_valid_request (t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)
	mockService := service.NewMockUserService(ctrl)
	uh := UserHandlers{mockService}
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", uh.UpdateUserWeights).Methods(http.MethodPatch)

	mockRequestDto := dto.SetUserWeightsRequest{
		Id:                "valid_test_id",
		WeightEnvironment: intPtr(1),
		WeightGender:      intPtr(2),
		WeightLgbtq:       intPtr(3),
		WeightEquality:    intPtr(0),
		WeightCorruption:  intPtr(2),
		WeightFreedom:     intPtr(4),
	}

	mockResponseDto := dto.SetUserWeightsResponse{
		Matched: true,
		Updated: true,
	}

	jsonBody := []byte(`{"weight_environment":1,"weight_gender":2,"weight_lgbtq":3,"weight_corruption":2,"weight_freedom":4,"weight_equality":0}`)
	mockService.EXPECT().UpdateWeights(mockRequestDto).Return(&mockResponseDto, nil)
	request, _ := http.NewRequest(http.MethodPatch, "/user/valid_test_id", bytes.NewBuffer(jsonBody))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while writing the code")
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	expected := `{"matched":true,"updated":true}`

	if expected != resBody {
		t.Error("Response body not matching")
	}
}

func Test_updateUserWeights_should_return_code_400_and_error_message_with_invalid_json_body (t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mockService := service.NewMockUserService(ctrl)
	uh := UserHandlers{mockService}
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", uh.UpdateUserWeights).Methods(http.MethodPatch)

	invalidJsonBody := []byte(`"weight_environment":1,"weight_gender":2,"weight_lgbtq":3,"weight_corruption":2,"weight_freedom":4,"weight_equality":0}`)
	request, _ := http.NewRequest(http.MethodPatch, "/user/valid_test_id", bytes.NewBuffer(invalidJsonBody))

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

func Test_updateUserWeights_should_return_code_422_and_validation_message_with_invalid_request (t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mockService := service.NewMockUserService(ctrl)
	uh := UserHandlers{mockService}
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", uh.UpdateUserWeights).Methods(http.MethodPatch)

	jsonBody := []byte(`{"weight_environment":10,"weight_gender":2,"weight_lgbtq":3,"weight_corruption":2,"weight_freedom":4,"weight_equality":0}`)
	request, _ := http.NewRequest(http.MethodPatch, "/user/valid_test_id", bytes.NewBuffer(jsonBody))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusUnprocessableEntity {
		t.Error("Failed while writing the code")
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	expected := `{"invalid_fields":["WeightEnvironment"]}`

	if expected != resBody {
		t.Error("Response body not matching")
	}
}

func Test_updateUserWeights_should_return_error_with_code_if_service_returns_error (t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)
	mockService := service.NewMockUserService(ctrl)
	uh := UserHandlers{mockService}
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", uh.UpdateUserWeights).Methods(http.MethodPatch)

	mockRequestDto := dto.SetUserWeightsRequest{
		Id:                "valid_test_id",
		WeightEnvironment: intPtr(1),
		WeightGender:      intPtr(2),
		WeightLgbtq:       intPtr(3),
		WeightEquality:    intPtr(0),
		WeightCorruption:  intPtr(2),
		WeightFreedom:     intPtr(4),
	}

	jsonBody := []byte(`{"weight_environment":1,"weight_gender":2,"weight_lgbtq":3,"weight_corruption":2,"weight_freedom":4,"weight_equality":0}`)
	mockService.EXPECT().UpdateWeights(mockRequestDto).Return(nil, errs.NewUnexpectedError("unexpected server error"))
	request, _ := http.NewRequest(http.MethodPatch, "/user/valid_test_id", bytes.NewBuffer(jsonBody))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while writing the code")
	}

	resBody := recorder.Body.String()
	resBody = resBody[:len(resBody)-1]
	expected := `{"message":"unexpected server error"}`

	if expected != resBody {
		t.Error("Response body not matching")
	}
}
