package app

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"github.com/luschnat-ziegler/cc_backend_go/mocks/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func intPtr(i int) *int {
	return &i
}

func Test_should_return_countries_with_status_code_200(t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)
	mockService := service.NewMockCountryService(ctrl)
	ch := CountryHandlers{mockService}
	router := mux.NewRouter()

	dummyCountries := []*dto.GetCountryResponse{
		{
			ID: primitive.NewObjectID(),
			Name: "Testistan",
			Region: "Earth",
			Freedom: intPtr(34) ,
			Gender: intPtr(79),
			Lgbtq: intPtr(23),
			Environment: intPtr(45),
			Corruption: intPtr(59),
			Inequality: intPtr(90),
			Total: intPtr(45),
		},
		{
			ID: primitive.NewObjectID(),
			Name: "Testland",
			Region: "Earth",
			Freedom: intPtr(22) ,
			Gender: intPtr(25),
			Lgbtq: nil,
			Environment: intPtr(78),
			Corruption: intPtr(45),
			Inequality: intPtr(66),
			Total: intPtr(56),
		},
	}

	mockService.EXPECT().GetAll().Return(dummyCountries, nil)

	router = mux.NewRouter()
	router.HandleFunc("/countries", ch.getAllCountries)

	request, _ := http.NewRequest(http.MethodGet, "/countries", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while writing the code")
	}


	var resultCountries []*dto.GetCountryResponse
	if err := json.NewDecoder(recorder.Body).Decode(&resultCountries); err != nil {
		t.Error("Failed to re-decode response body")
	}

	for i, country := range resultCountries {
		if country.Name != dummyCountries[i].Name {
			t.Error("Response does not correspond to data")
			break
		}
	}
}

func Test_should_return_statue_code_500_with_error_message (t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockCountryService(ctrl)

	mockReturnError := errs.NewUnexpectedError("Unexpected server error")
	mockService.EXPECT().GetAll().Return(nil, mockReturnError)
	ch := CountryHandlers{mockService}

	router := mux.NewRouter()
	router.HandleFunc("/countries", ch.getAllCountries)

	request, _ := http.NewRequest(http.MethodGet, "/countries", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while writing the code")
	}

	var resultError errs.AppError
	if err := json.NewDecoder(recorder.Body).Decode(&resultError); err != nil {
		t.Error("Failed to re-decode response body")
	}

	if resultError.Message != mockReturnError.Message {
		t.Error("Response error message does not match")
	}
}