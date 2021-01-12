package service

import (
	"github.com/luschnat-ziegler/cc_backend_go/domain"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"github.com/luschnat-ziegler/cc_backend_go/logger"
)

//go:generate mockgen -destination=../mocks/service/mockCountryService.go -package=service github.com/luschnat-ziegler/cc_backend_go/service CountryService
type CountryService interface {
	GetAll() ([]dto.GetCountryResponse, *errs.AppError)
}

type DefaultCountryService struct {
	repo domain.CountryRepository
}

func (s DefaultCountryService) GetAll() ([]dto.GetCountryResponse, *errs.AppError) {
	countries, err := s.repo.FindAll()
	if err != nil {
		logger.Error("Error returned by CountryRepository.FindAll(): " + err.Message)
		return nil, errs.NewUnexpectedError("Error in findAll method")
	}
	response := make([]dto.GetCountryResponse, 0)
	for _, c := range countries {
		response = append(response, c.ToGetCountryResponseDto())
	}
	return response, nil
}

func NewCountryService(repository domain.CountryRepository) DefaultCountryService {
	return DefaultCountryService{repository}
}