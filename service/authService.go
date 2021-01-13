package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/luschnat-ziegler/cc_backend_go/domain"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"github.com/luschnat-ziegler/cc_backend_go/logger"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	"time"
)

//go:generate mockgen -destination=../mocks/service/mockAuthService.go -package=service github.com/luschnat-ziegler/cc_backend_go/service AuthService
type AuthService interface {
	LogIn(request dto.LogInRequest) (*dto.LogInResponse, *errs.AppError)
	Verify(authHeader string) (*string, *errs.AppError)
}

type DefaultAuthService struct {
	repo domain.UserRepository
}

func (s DefaultAuthService) Verify(authHeader string) (*string, *errs.AppError) {
	splitToken := strings.Split(authHeader, "Bearer")
	if len(splitToken) != 2 {
		return nil, errs.NewUnauthorizedError("Invalid auth header")
	}

	secret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return nil, errs.NewUnexpectedError("unexpected server error")
	}

	tokenString := strings.TrimSpace(splitToken[1])
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {return []byte(secret), nil})
	if err != nil {
		return nil, errs.NewUnauthorizedError("Invalid token")
	}

	idFromToken := token.Claims.(jwt.MapClaims)["sub"].(string)
	return &idFromToken, nil
}

func (s DefaultAuthService) LogIn (request dto.LogInRequest) (*dto.LogInResponse, *errs.AppError) {
	user, err := s.repo.ByEmail(request.Email)
	if err != nil {
		return nil, errs.NewNotFoundError("User not found")
	}

	if !checkPasswordHash(request.Password, user.Password) {
		return nil, errs.NewNotFoundError("Wrong Password")
	}

	id := user.ID.Hex()

	claims := jwt.MapClaims{
		"sub": 		id,
		"exp":      time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret, _ := os.LookupEnv("JWT_SECRET")
	signedTokenAsString, e := token.SignedString([]byte(secret))
	if e != nil {
		logger.Error("Error signing token: " + e.Error())
		return nil, errs.NewUnexpectedError("Error generating token")
	}

	return &dto.LogInResponse{
		Success: true,
		Token:   &signedTokenAsString,
	}, nil

}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewAuthService(repo domain.UserRepository) DefaultAuthService {
	return DefaultAuthService{repo}
}