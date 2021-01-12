package domain

import (
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Country struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"country,omitempty"`
	Region string `json:"region,omitempty" bson:"region,omitempty"`
	Freedom *int `json:"freedom,omitempty" bson:"freedom,omitempty"`
	Gender *int `json:"gender,omitempty" bson:"gender,omitempty"`
	Lgbtq *int `json:"lgbtq,omitempty" bson:"lgbtq,omitempty"`
	Environment *int `json:"environment,omitempty" bson:"environment,omitempty"`
	Corruption *int `json:"corruption,omitempty" bson:"corruption,omitempty"`
	Inequality *int `json:"inequality,omitempty" bson:"inequality,omitempty"`
	Total *int `json:"total,omitempty" bson:"total,omitempty"`
}

//go:generate mockgen -destination=../mocks/domain/mockCountryRepository.go -package=domain github.com/luschnat-ziegler/cc_backend_go/domain CountryRepository
type CountryRepository interface {
	FindAll() ([]Country, *errs.AppError)
}

func (c Country) ToGetCountryResponseDto() dto.GetCountryResponse {
	return dto.GetCountryResponse{
		ID: c.ID,
		Name: c.Name,
		Region: c.Region,
		Freedom: c.Freedom,
		Gender: c.Gender,
		Lgbtq: c.Lgbtq,
		Environment: c. Environment,
		Corruption: c.Corruption,
		Inequality: c.Inequality,
		Total: c.Total,
	}
}

