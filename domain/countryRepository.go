package domain

import (
	"context"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type CountryRepositoryDB struct {
	client *mongo.Client
}

func (countryRepositoryDB CountryRepositoryDB) FindAll() ([]Country, *errs.AppError) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := countryRepositoryDB.client.Connect(ctx)
	if err != nil {
		return nil, errs.NewUnexpectedError("error connecting to DB")
	}
	defer countryRepositoryDB.client.Disconnect(ctx)

	collection := countryRepositoryDB.client.Database("countrycheck").Collection("countries")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
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