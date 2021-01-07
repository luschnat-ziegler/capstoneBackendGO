package domain

import (
	"github.com/luschnat-ziegler/cc_backend_go/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func getDbClient() *mongo.Client {

	url, _ := os.LookupEnv("DB_URL")

	client, err  := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		logger.Error("Database init error: " + err.Error())
	}

	return client
}
