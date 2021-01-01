package app

import (
	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/cc_backend_go/domain"
	"github.com/luschnat-ziegler/cc_backend_go/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	dbClient := getDbClient()

	countryRepositoryDb := domain.NewCountryRepositoryDb(dbClient)
	ch := CountryHandlers{service.NewCountryService(countryRepositoryDb)}

	router.HandleFunc("/countries", ch.getAllCountries).Methods(http.MethodGet).
		Name("GetAllCountries")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getDbClient() *mongo.Client {

	client, err  := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mlz:horst714@cluster0.z1fxp.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	return client
}
