package dto

type GetUserResponse struct {
	Email             string `json:"email"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	WeightEnvironment int    `json:"weight_environment"`
	WeightGender      int    `json:"weight_gender"`
	WeightLgbtq       int    `json:"weight_lgbtq`
	WeightEquality    int    `json:"weight_equality"`
	WeightCorruption  int    `json:"weight_corruption"`
	WeightFreedom     int    `json:"weight_freedom"`
}
