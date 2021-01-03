package app

import (
	"encoding/json"
	"fmt"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
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
		fmt.Println("Error decoding request body")
	}

	result, e := ah.service.LogIn(logInRequest)
	if e != nil {
		writeResponse(w, http.StatusBadRequest, e.Message)
	} else {
		writeResponse(w, http.StatusOK, result)
	}
}