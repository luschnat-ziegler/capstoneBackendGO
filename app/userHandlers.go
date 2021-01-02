package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/service"
	"net/http"
)

type UserHandlers struct {
	service service.UserService
}

func (uh *UserHandlers) GetUserById(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["user_id"]

	user, err := uh.service.GetUser(id)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Message)
	} else {
		writeResponse(w, http.StatusOK, user)
	}
}

func (uh *UserHandlers) CreateUser(w http. ResponseWriter, r*http.Request) {

	var createUserRequest dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&createUserRequest)
	if err != nil {
		fmt.Println(err)
	}

	result, e := uh.service.CreateUser(createUserRequest)
	if e != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		writeResponse(w, http.StatusCreated, result)
	}
}