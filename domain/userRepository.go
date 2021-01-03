package domain

import (
	"context"
	"fmt"
	"github.com/luschnat-ziegler/cc_backend_go/dto"
	"github.com/luschnat-ziegler/cc_backend_go/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type UserRepositoryDB struct {}

func (userRepositoryDB UserRepositoryDB) ById(id string) (*User, *error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := getDbClient()

	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("countrycheck").Collection("user")

	var user User
	objectId, _ := primitive.ObjectIDFromHex(id)
	err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	return &user, nil
}

func (userRepositoryDB UserRepositoryDB) Save(user User) (string, *errs.AppError) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := getDbClient()

	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	hashPw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashPw)

	collection := client.Database("countrycheck").Collection("user")
	result, err := collection.InsertOne(ctx, user)
	if err != nil {log.Fatal(err)}

	resultAsString := result.InsertedID.(primitive.ObjectID).Hex()

	return resultAsString, nil
}

func (userRepositoryDB UserRepositoryDB) UpdateWeights(request dto.SetUserWeightsRequest) (*dto.SetUserWeightsResponse, *error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := getDbClient()

	err := client.Connect(ctx)
	if err != nil {
		return nil, &err
	}
	defer client.Disconnect(ctx)

	collection := client.Database("countrycheck").Collection("user")

	objectId, _ := primitive.ObjectIDFromHex(request.Id)

	filter := bson.M{"_id": objectId}

	update := bson.M{
		"$set": bson.M{
			"weightenvironment": request.WeightEnvironment,
			"weightgender_": request.WeightGender,
			"weightlgbtq": request.WeightLgbtq,
			"weightequality": request.WeightEquality,
			"weightcorruption": request.WeightCorruption,
			"weightfreedom": request.WeightFreedom,
		},
	}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, &err
	}

	if res.ModifiedCount == 0 && res.MatchedCount > 0 {
		return &dto.SetUserWeightsResponse{
			Matched: true,
			Updated: false,
		}, nil
	} else if res.ModifiedCount > 0 {
		return &dto.SetUserWeightsResponse{
			Matched: true,
			Updated: true,
		}, nil
	} else {
		return &dto.SetUserWeightsResponse{
			Matched: false,
			Updated: false,
		}, nil
	}
}

func NewUserRepositoryDb() UserRepositoryDB {
	return UserRepositoryDB{}
}

