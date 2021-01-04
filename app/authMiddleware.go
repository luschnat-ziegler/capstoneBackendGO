package app

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
)

type AuthMiddleware struct {}

func (am AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			currentRouteName := mux.CurrentRoute(r).GetName()
			if !((currentRouteName == "GetUser") || (currentRouteName == "UpdateUserWeights")) {
				next.ServeHTTP(w, r)
			} else {
				authHeader := r.Header.Get("Authorization")
				if authHeader != "" {
					secret,_ := os.LookupEnv("JWT_SECRET")
					tokenString := getTokenFromHeader(authHeader)

					token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
						return []byte(secret), nil
					})
					if err != nil {
						fmt.Println("Error parsing token")
					}

					fmt.Println(token.Claims)
					next.ServeHTTP(w, r)
				} else {
					writeResponse(w, http.StatusUnauthorized, "missing token")
				}
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}