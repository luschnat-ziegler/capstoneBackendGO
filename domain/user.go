package domain

import (
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string
	Password string
	FirstName string
	LastName string
	WeightEnvironment int
	WeightGender int
	WeightLgbtq int
	WeightEquality int
	WeightCorruption int
	WeightFreedom int
}

type UserRepository interface {
	ById(string) (*User, error)
	Save(user User) (*User, *error)
}

func NewUser(createUserRequest dto.CreateUserRequest) User {
	return User{
		Email: createUserRequest.Email,
		Password: createUserRequest.Password,
		FirstName: createUserRequest.FirstName,
		LastName: createUserRequest.LastName,
		WeightEnvironment: 2,
		WeightGender: 2,
		WeightLgbtq: 2,
		WeightEquality: 2,
		WeightCorruption: 2,
		WeightFreedom: 2,
	}
}