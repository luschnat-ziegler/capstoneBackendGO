package service

import "github.com/luschnat-ziegler/cc_backend_go/domain"

type UserService interface {
	CreateUser(string) (domain.User, *error)
	GetUser(string) (domain.User, error)
}

