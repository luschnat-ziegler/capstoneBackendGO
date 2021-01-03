package dto

type SetUserWeightsRequest struct {
	Id string				`json:"user_id"`
	WeightEnvironment int	`json:"weight_environment"`
	WeightGender int		`json:"weight_gender"`
	WeightLgbtq int			`json:"weight_lgbtq"`
	WeightEquality int		`json:"weight_equality"`
	WeightCorruption int	`json:"weight_corruption"`
	WeightFreedom int 		`json:"weight_freedom"`
}