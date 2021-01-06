package domain

import (
	"context"
	"fmt"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"go.mongodb.org/mongo-driver/bson"
	"log"
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
		return nil, errs.NewUnexpectedError("Database Error")
	}
	defer client.Disconnect(ctx)

	collection := client.Database("countrycheck").Collection("countries")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err)
		return nil, errs.NewUnexpectedError("Database Error")
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
		if err != nil {
			log.Println(err)
			return nil, errs.NewUnexpectedError("Database Error")
		}
		output = append(output, country)
	}

	return output, nil
}

func NewCountryRepositoryDb() CountryRepositoryDB {
	return CountryRepositoryDB{}
}