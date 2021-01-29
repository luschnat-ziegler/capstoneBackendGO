package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/cc_backend_go/domain"
	"github.com/luschnat-ziegler/cc_backend_go/logger"
	"github.com/luschnat-ziegler/cc_backend_go/service"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Start() {

	logger.Info("Application started...")

	sanityCheck()

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

	router.HandleFunc("/user/{id}", uh.GetUserById).
		Methods(http.MethodGet).
		Name("GetUser")

	router.HandleFunc("/user/{id}", uh.UpdateUserWeights).
		Methods(http.MethodPatch).
		Name("UpdateUserWeights")

	router.HandleFunc("/login", ah.logInUser).
		Methods(http.MethodPost).
		Name("LogInUser")

	am := AuthMiddleware{service.NewAuthService(userRepositoryDB)}
	middleWareHandleFunc := am.authorizationHandler()
	router.Use(middleWareHandleFunc)

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}


func sanityCheck() {
	envVars := []string{
		"DB_URL",
		"JWT_SECRET",
	}
	for _, envVar := range envVars {
		_, ok := os.LookupEnv(envVar)
		if !ok {
			log.Fatal(fmt.Sprintf("Environment variable %s not set in .env. Terminating application...", envVar))
		}
	}
}
