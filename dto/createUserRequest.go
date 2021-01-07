package dto

import (
	"github.com/go-playground/validator"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
)

type CreateUserRequest struct {
	Email string		`json:"email" validate:"required,email"`
	Password string		`json:"password" validate:"required,min=6,max=32"`
	FirstName string	`json:"first_name" validate:"required"`
	LastName string		`json:"last_name" validate:"required"`
}

var v *validator.Validate

func (createUserRequest CreateUserRequest) Validate() *errs.ValidationError {
	v = validator.New()
	err := v.Struct(createUserRequest)

	if err != nil {
		var invalidFields []string
		for _, e := range err.(validator.ValidationErrors) {
			invalidFields = append(invalidFields, e.Field())
		}
		return errs.NewValidationError(invalidFields)
	} else {
		return nil
	}
}