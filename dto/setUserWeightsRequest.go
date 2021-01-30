package dto

import (
	"github.com/go-playground/validator"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
)

type SetUserWeightsRequest struct {
	Id                string `json:"user_id"`
	WeightEnvironment *int   `json:"weight_environment,omitempty" validate:"required,min=0,max=4"`
	WeightGender      *int   `json:"weight_gender,omitempty" validate:"required,min=0,max=4"`
	WeightLgbtq       *int   `json:"weight_lgbtq,omitempty" validate:"required,min=0,max=4"`
	WeightEquality    *int   `json:"weight_equality,omitempty" validate:"required,min=0,max=4"`
	WeightCorruption  *int   `json:"weight_corruption,omitempty" validate:"required,min=0,max=4"`
	WeightFreedom     *int   `json:"weight_freedom,omitempty" validate:"required,min=0,max=4"`
}

func (setUserWeightsRequest SetUserWeightsRequest) Validate() *errs.ValidationError {
	v = validator.New()
	err := v.Struct(setUserWeightsRequest)

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
