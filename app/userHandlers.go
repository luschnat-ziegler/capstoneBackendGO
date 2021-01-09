package app

import (
	"encoding/json"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"github.com/luschnat-ziegler/cc_backend_go/service"
	"net/http"
)

type UserHandlers struct {
	service service.UserService
}

func (uh *UserHandlers) GetUserById(w http.ResponseWriter, r *http.Request) {

	idAsString, err := userIdStringFromToken(r)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, errs.NewBadRequestError("Invalid token").AsMessage())
	} else {
		user, appError := uh.service.GetUser(*idAsString)
		if appError != nil {
			writeResponse(w, http.StatusBadRequest, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, user)
		}
	}
}

func (uh *UserHandlers) CreateUser(w http. ResponseWriter, r *http.Request) {

	var createUserRequest dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&createUserRequest)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, errs.NewBadRequestError("Body parsing error").AsMessage())
	} else {
		validationError := createUserRequest.Validate()
		if validationError != nil {
			writeResponse(w, validationError.Code, validationError.AsMessage())
		} else {
			result, appError := uh.service.CreateUser(createUserRequest)
			if appError != nil {
				writeResponse(w, appError.Code, appError.AsMessage())
			} else {
				writeResponse(w, http.StatusCreated, result)
			}
		}
	}
}

func (uh *UserHandlers) UpdateUserWeights(w http. ResponseWriter, r *http.Request) {

	var setUserWeightsRequest dto.SetUserWeightsRequest
	err := json.NewDecoder(r.Body).Decode(&setUserWeightsRequest)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, errs.NewBadRequestError("Body parsing error").AsMessage())
	} else {
		validationError := setUserWeightsRequest.Validate()
		if validationError != nil {
			writeResponse(w, validationError.Code, validationError.AsMessage())
		} else {
			idAsString, err := userIdStringFromToken(r)
			if err != nil {
				writeResponse(w, http.StatusBadRequest, errs.NewBadRequestError("Invalid token").AsMessage())
			} else {
				setUserWeightsRequest.Id = *idAsString
				result, appError := uh.service.UpdateWeights(setUserWeightsRequest)
				if appError != nil {
					writeResponse(w, appError.Code, appError.AsMessage())
				} else {
					writeResponse(w, http.StatusOK, result)
				}
			}
		}
	}
}