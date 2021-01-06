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

func (setUserWeightsRequest SetUserWeightsRequest) HasValidWeights() bool {
	return (setUserWeightsRequest.WeightEnvironment >= 0 && setUserWeightsRequest.WeightEnvironment <= 4) &&
		(setUserWeightsRequest.WeightGender >= 0 && setUserWeightsRequest.WeightGender <= 4) &&
		(setUserWeightsRequest.WeightLgbtq >= 0 && setUserWeightsRequest.WeightLgbtq <= 4) &&
		(setUserWeightsRequest.WeightEquality >= 0 && setUserWeightsRequest.WeightEquality <= 4) &&
		(setUserWeightsRequest.WeightCorruption >= 0 && setUserWeightsRequest.WeightCorruption <= 4) &&
		(setUserWeightsRequest.WeightFreedom >= 0 && setUserWeightsRequest.WeightFreedom <= 4)
}