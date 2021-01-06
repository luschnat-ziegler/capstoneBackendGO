package service

import (
	"github.com/luschnat-ziegler/cc_backend_go/domain"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
)

type UserService interface {
	CreateUser(request dto.CreateUserRequest) (*dto.CreateUserResponse, *errs.AppError)
	GetUser(string) (*dto.GetUserResponse, *errs.AppError)
	UpdateWeights(request dto.SetUserWeightsRequest) (*dto.SetUserWeightsResponse, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (s DefaultUserService) GetUser(id string) (*dto.GetUserResponse, *errs.AppError) {
	user, err := s.repo.ById(id)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error in ById Method")
	}
	return user.ToGetUserResponse(), nil
}

func (s DefaultUserService) CreateUser(createUserRequest dto.CreateUserRequest) (*dto.CreateUserResponse, *errs.AppError) {
	user := domain.NewUser(createUserRequest)
	result, err := s.repo.Save(user)
	if err != nil {
		return nil, err
	}

	return &dto.CreateUserResponse{Id: *result}, nil
}

func (s DefaultUserService) UpdateWeights(setUserWeightsRequest dto.SetUserWeightsRequest) (*dto.SetUserWeightsResponse, *errs.AppError) {
	res, err := s.repo.UpdateWeights(setUserWeightsRequest)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error in update method")
	} else {
		return res, nil
	}
}

func NewUserService(repo domain.UserRepository) DefaultUserService {
	return DefaultUserService{repo}
}