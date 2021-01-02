package app

import (

	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/cc_backend_go/domain"
	"github.com/luschnat-ziegler/cc_backend_go/service"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	countryRepositoryDb := domain.NewCountryRepositoryDb()
	userRepositoryDB := domain.NewUserRepositoryDb()
	ch := CountryHandlers{service.NewCountryService(countryRepositoryDb)}
	uh := UserHandlers{service.NewUserService(userRepositoryDB)}

	router.HandleFunc("/countries", ch.getAllCountries).
		Methods(http.MethodGet).
		Name("GetAllCountries")

	router.HandleFunc("/user", uh.CreateUser).
		Methods(http.MethodPost).
		Name("CreateUser")

	router.HandleFunc("/user/{user_id}", uh.GetUserById).
		Methods(http.MethodGet).
		Name("GetUser")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}