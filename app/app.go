package app

import (

	"github.com/gorilla/mux"
	"github.com/luschnat-ziegler/cc_backend_go/domain"
	"github.com/luschnat-ziegler/cc_backend_go/service"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

func Start() {

	router := mux.NewRouter()
	dbClient := getDbClient()

	countryRepositoryDb := domain.NewCountryRepositoryDb(dbClient)
	ch := CountryHandlers{service.NewCountryService(countryRepositoryDb)}

	router.HandleFunc("/countries", ch.getAllCountries).Methods(http.MethodGet).
		Name("GetAllCountries")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func getDbClient() *mongo.Client {

	url, _ := os.LookupEnv("DB_URL")

	client, err  := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	return client
}
