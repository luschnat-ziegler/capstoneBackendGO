package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"net/http"
	"os"
)

type AuthMiddleware struct {}

func (am AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			currentRouteName := mux.CurrentRoute(r).GetName()
			if !((currentRouteName == "GetUser") || (currentRouteName == "UpdateUserWeights")) {
				next.ServeHTTP(w, r)
			} else if authHeader := r.Header.Get("Authorization"); authHeader == "" {
				appError := errs.NewUnauthorizedError("Token missing")
				writeResponse(w, appError.Code, appError.AsMessage())
			} else if secret, ok := os.LookupEnv("JWT_SECRET"); !ok {
				appError := errs.NewUnexpectedError("Unexpected server error")
				writeResponse(w, appError.Code, appError.AsMessage())
			} else if token , err := jwt.Parse(getTokenFromHeader(authHeader), func(token *jwt.Token) (interface{}, error) {return []byte(secret), nil}); err != nil {
				appError := errs.NewUnauthorizedError("Token invalid")
				writeResponse(w, appError.Code, appError.AsMessage())
			} else if token.Claims.(jwt.MapClaims)["sub"].(string) != mux.Vars(r)["id"] {
				appError := errs.NewUnauthorizedError("Token not matching user")
				writeResponse(w, appError.Code, appError.AsMessage())
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}