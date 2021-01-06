package app

import (
	"encoding/json"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"github.com/luschnat-ziegler/cc_backend_go/service"
	"net/http"
)

type AuthHandlers struct {
	service service.AuthService
}

func (ah *AuthHandlers) logInUser (w http.ResponseWriter, r *http.Request) {

	var logInRequest dto.LogInRequest
	err := json.NewDecoder(r.Body).Decode(&logInRequest)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, errs.NewBadRequestError("Body parsing error").AsMessage())
	}

	result, appError := ah.service.LogIn(logInRequest)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, result)
	}
}