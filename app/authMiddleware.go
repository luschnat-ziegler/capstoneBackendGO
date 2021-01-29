package app

import (
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"github.com/luschnat-ziegler/cc_backend_go/service"
	"net/http"
)

type AuthMiddleware struct {
	service service.AuthService
}

func (am AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			currentRouteName := mux.CurrentRoute(r).GetName()
			if !((currentRouteName == "GetUser") || (currentRouteName == "UpdateUserWeights")) {
				next.ServeHTTP(w, r)
				return
			} else if authHeader := r.Header.Get("Authorization"); authHeader == "" {
				appError := errs.NewUnauthorizedError("Token missing")
				writeResponse(w, appError.Code, appError.AsMessage())
				return
			} else if tokenUserId, appError := am.service.Verify(authHeader); appError != nil {
				writeResponse(w, appError.Code, appError.AsMessage())
				return
			} else if *tokenUserId != mux.Vars(r)["id"] {
				appError := errs.NewUnauthorizedError("Token not matching requested user")
				writeResponse(w, appError.Code, appError.AsMessage())
				return
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}