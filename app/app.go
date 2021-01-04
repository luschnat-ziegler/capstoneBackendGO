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
	ah := AuthHandlers{service.NewAuthService(userRepositoryDB)}

	router.HandleFunc("/countries", ch.getAllCountries).
		Methods(http.MethodGet).
		Name("GetAllCountries")

	router.HandleFunc("/user", uh.CreateUser).
		Methods(http.MethodPost).
		Name("CreateUser")

	router.HandleFunc("/user/", uh.GetUserById).
		Methods(http.MethodGet).
		Name("GetUser")

	router.HandleFunc("/user", uh.UpdateUserWeights).
		Methods(http.MethodPatch).
		Name("UpdateUserWeights")

	router.HandleFunc("/login", ah.logInUser).
		Methods(http.MethodPost).
		Name("LogInUser")

	am := AuthMiddleware{}
	router.Use(am.authorizationHandler())

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}