package domain

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func getDbClient() *mongo.Client {

	url, _ := os.LookupEnv("DB_URL")

	client, err  := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	return client
}
