package app

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

//go:generate mockgen -destination=../mocks/app/mockGetUserIdStringFromToken.go -package=app github.com/luschnat-ziegler/cc_backend_go/app GetUserIdStringFromToken
func GetUserIdStringFromToken(r *http.Request) (*string, error) {
	secret, _ := os.LookupEnv("JWT_SECRET")
	authHeader := r.Header.Get("Authorization")
	tokenString := getTokenFromHeader(authHeader)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	idAsString := token.Claims.(jwt.MapClaims)["sub"].(string)
	return &idAsString, nil
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func getTokenFromHeader(header string) string {
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}