package dto

type LogInRequest struct {
	Email string		`json:"email"`
	Password string		`json:"password"`
}
