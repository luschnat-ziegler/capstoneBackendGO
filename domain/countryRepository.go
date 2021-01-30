package domain

import (
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"github.com/luschnat-ziegler/cc_backend_go/logger"
	"go.mongodb.org/mongo-driver/bson"
)

type CountryRepositoryDB struct {
}

func (countryRepositoryDB CountryRepositoryDB) FindAll() ([]Country, *errs.AppError) {

	client, ctx, cancel, err := getDbClient()
	if err != nil {
		logger.Error("Error connecting to database: " + err.Error())
		return nil, errs.NewUnexpectedError("Database Error")
	}
	defer disconnectClient(client, ctx)
	defer cancel()

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
