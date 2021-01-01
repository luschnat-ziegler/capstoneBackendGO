package service

import (
	"github.com/luschnat-ziegler/cc_backend_go/domain"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
)

type CountryService interface {
	GetAll() ([]*dto.GetCountryResponse, *errs.AppError)
}

type DefaultCountryService struct {
	repo domain.CountryRepository
}

func (s DefaultCountryService) GetAll() ([]*dto.GetCountryResponse, *errs.AppError) {
	countries, err := s.repo.FindAll()
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected server error")
	}
	response := make([]*dto.GetCountryResponse, 0)
	for _, c := range countries {
		response = append(response, c.ToGetCountryResponseDto())
	}
	return response, nil
}

func NewCountryService(repository domain.CountryRepository) DefaultCountryService {
	return DefaultCountryService{repository}
}