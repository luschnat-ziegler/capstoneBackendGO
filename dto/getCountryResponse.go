package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetCountryResponse struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"country,omitempty"`
	Region string `json:"region,omitempty" bson:"region,omitempty"`
	Freedom *int `json:"freedom" bson:"freedom"`
	Gender *int `json:"gender" bson:"gender"`
	Lgbtq *int `json:"lgbtq" bson:"lgbtq"`
	Environment *int `json:"environment" bson:"environment"`
	Corruption *int `json:"corruption" bson:"corruption"`
	Inequality *int `json:"inequality" bson:"inequality"`
	Total *int `json:"total" bson:"total"`
}
