package service

import (
	"github.com/luschnat-ziegler/cc_backend_go/domain"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
)

type AuthService interface {
	Login()

}

type DefaultCountryService struct {
	repo domain.CountryRepository
}