package domain

import (
	"context"
	"fmt"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type CountryRepositoryDB struct {
	client *mongo.Client
}

func (countryRepositoryDB CountryRepositoryDB) FindAll() ([]Country, *errs.AppError) {

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	url, _ := os.LookupEnv("DB_URL")
	client, err  := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, errs.NewUnexpectedError("error connecting to DB")
	}
	defer client.Disconnect(ctx)

	collection := client.Database("countrycheck").Collection("countries")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err)
		return nil, errs.NewUnexpectedError("error querying country collection")
	}

	defer func() {
		err = cursor.Close(ctx)
		if err != nil {
			log.Println("Could not close cursor")
		}
	}()

	var output []Country
	for cursor.Next(ctx) {
		var country Country
		err := cursor.Decode(&country)
		if err != nil { log.Fatal(err) }
		output = append(output, country)
	}

	return output, nil
}

func NewCountryRepositoryDb(dbClient *mongo.Client) CountryRepositoryDB {
	return CountryRepositoryDB{dbClient}
}