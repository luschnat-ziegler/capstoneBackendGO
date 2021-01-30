package domain

import (
	"context"
	"github.com/luschnat-ziegler/cc_backend_go/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var clientInstance *mongo.Client
var clientInstanceError error

func getDbClient() (*mongo.Client, context.Context, context.CancelFunc, error) {

	url, _ := os.LookupEnv("DB_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		logger.Error("Database init error: " + err.Error())
		clientInstanceError = err
	}

	clientInstance = client

	return clientInstance, ctx, cancel, clientInstanceError
}

func disconnectClient(client *mongo.Client, ctx context.Context) {
	err := client.Disconnect(ctx)
	if err != nil {
		logger.Error("Error disconnecting from db client: " + err.Error())
	}
}
