package dto

type GetUserResponse struct {
	Email string
	Password string
	FirstName string
	LastName string
	WeightEnvironment int
	WeightGender int
	WeightLgbtq int
	WeightEquality int
	WeightCorruption int
}
