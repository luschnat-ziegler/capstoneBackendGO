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

func (userRepositoryDB UserRepositoryDB) ByEmail(email string) (*User, *error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := getDbClient()

	err := client.Connect(ctx)
	if err != nil {
		log.Println(err)
		return nil, &err
	}
	defer client.Disconnect(ctx)

	collection := client.Database("countrycheck").Collection("user")

	var user User
	err = collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, &err
	}

	return &user, nil
}

func (userRepositoryDB UserRepositoryDB) ById(id string) (*User, *errs.AppError) {

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

func (userRepositoryDB UserRepositoryDB) Save(user User) (*string, *errs.AppError) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := getDbClient()

	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("countrycheck").Collection("user")

	count, err :=collection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		return nil, errs.NewUnexpectedError("Database error")
	}

	if count > 0 {
		return nil, errs.NewConflictError("Existing user")
	}

	hashPw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashPw)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {log.Fatal(err)}

	resultAsString := result.InsertedID.(primitive.ObjectID).Hex()

	return &resultAsString, nil
}

func (userRepositoryDB UserRepositoryDB) UpdateWeights(request dto.SetUserWeightsRequest) (*dto.SetUserWeightsResponse, *errs.AppError) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := getDbClient()

	err := client.Connect(ctx)
	if err != nil {
		log.Println(err)
		return nil, errs.NewUnexpectedError("Database error")
	}
	defer client.Disconnect(ctx)

	collection := client.Database("countrycheck").Collection("user")

	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		log.Println(err)
		return nil, errs.NewUnexpectedError("ID parsing error")
	}

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
		return nil, errs.NewNotFoundError(err.Error())
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

