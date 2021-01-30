package service

import (
	"fmt"
	"github.com/golang/mock/gomock"
	domain2 "github.com/luschnat-ziegler/cc_backend_go/domain"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"github.com/luschnat-ziegler/cc_backend_go/mocks/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

var mockCountryRepository *domain.MockCountryRepository
var countryService CountryService

func setupCountryServiceTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockCountryRepository = domain.NewMockCountryRepository(ctrl)
	countryService = NewCountryService(mockCountryRepository)
	return func() {
		countryService = nil
		defer ctrl.Finish()
	}
}

func Test_GetAll_should_return_slice_of_pointers_to_GetCountryResponse_and_nil_if_repo_method_returns_no_error(t *testing.T) {

	// Arrange
	teardown := setupCountryServiceTest(t)
	defer teardown()

	mockCountry := domain2.Country{
		ID:          primitive.ObjectID{},
		Name:        "Testistan",
		Region:      "TestRegion",
		Freedom:     intPtr(34),
		Gender:      intPtr(45),
		Lgbtq:       intPtr(23),
		Environment: nil,
		Corruption:  intPtr(15),
		Inequality:  intPtr(67),
		Total:       intPtr(36),
	}

	mockSlice := make([]domain2.Country, 0)
	mockSlice = append(mockSlice, mockCountry)

	mockCountryRepository.EXPECT().FindAll().Return(mockSlice, nil)

	// Act
	result, err := countryService.GetAll()

	// Assert
	if err != nil {
		t.Error("Error returned, nil expected.")
	}

	resultType := fmt.Sprintf("%T", result)
	if resultType != "[]dto.GetCountryResponse" {
		t.Error("Wrong type returned")
	}

	if result != nil && (result[0].Name != mockCountry.Name || result[0].Environment != nil) {
		t.Error("Returned data das not match mock")
	}
}

func Test_GetAll_should_return_nil_and_AppError_if_repo_method_returns_error(t *testing.T) {

	// Arrange
	teardown := setupCountryServiceTest(t)
	defer teardown()

	mockAppError := errs.NewUnexpectedError("unexpected server error")

	mockCountryRepository.EXPECT().FindAll().Return(nil, mockAppError)

	// Act
	result, err := countryService.GetAll()

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
