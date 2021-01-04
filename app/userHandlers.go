package app

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/service"
	"net/http"
	"os"
)

type UserHandlers struct {
	service service.UserService
}

func (uh *UserHandlers) GetUserById(w http.ResponseWriter, r *http.Request) {

	secret,_ := os.LookupEnv("JWT_SECRET")
	authHeader := r.Header.Get("Authorization")
	tokenString := getTokenFromHeader(authHeader)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("token invalid (userHandler)")
	}

	user, er := uh.service.GetUser(token.Claims.(jwt.MapClaims)["sub"].(string))
	if err != nil {
		writeResponse(w, http.StatusBadRequest, er.Message)
	} else {
		writeResponse(w, http.StatusOK, user)
	}
}

func (uh *UserHandlers) CreateUser(w http. ResponseWriter, r *http.Request) {

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

func (uh *UserHandlers) UpdateUserWeights(w http. ResponseWriter, r *http.Request) {

	var setUserWeightsRequest dto.SetUserWeightsRequest
	err := json.NewDecoder(r.Body).Decode(&setUserWeightsRequest)
	if err != nil {
		fmt.Println("Error decoding request body")
	}

	result, e := uh.service.UpdateWeights(setUserWeightsRequest)
	if e != nil {
		writeResponse(w, http.StatusBadRequest, e.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, result)
	}
}