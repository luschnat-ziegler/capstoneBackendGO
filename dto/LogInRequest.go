package dto

import (
	"github.com/go-playground/validator"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
)

type LogInRequest struct {
	Email string		`json:"email" validation:"required,email"`
	Password string		`json:"password" validation:"required"`
}

func (logInRequest LogInRequest) Validate() *errs.ValidationError {
	v := validator.New()
	err := v.Struct(logInRequest)

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
