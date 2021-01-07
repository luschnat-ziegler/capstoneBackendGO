package domain

import (
	"context"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"github.com/luschnat-ziegler/cc_backend_go/logger"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type CountryRepositoryDB struct {
}

func (countryRepositoryDB CountryRepositoryDB) FindAll() ([]Country, *errs.AppError) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := getDbClient()

	err := client.Connect(ctx)
	if err != nil {
		logger.Error("Error connecting to database: " + err.Error())
		return nil, errs.NewUnexpectedError("Database Error")
	}
	defer client.Disconnect(ctx)

	collection := client.Database("countrycheck").Collection("countries")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		logger.Error("Error querying database")
		return nil, errs.NewUnexpectedError("Database Error: " + err.Error())
	}

	defer func() {
		err = cursor.Close(ctx)
		if err != nil {
			logger.Error("Error closing cursor: " + err.Error())
		}
	}()

	var output []Country
	for cursor.Next(ctx) {
		var country Country
		err := cursor.Decode(&country)
		if err != nil {
			logger.Error("Error decoding database object: " + err.Error())
			return nil, errs.NewUnexpectedError("Database Error")
		}
		output = append(output, country)
	}

	return output, nil
}

func NewCountryRepositoryDb() CountryRepositoryDB {
	return CountryRepositoryDB{}
}