package domain

import (
	"context"
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

	var user User
	collection := client.Database("countrycheck").Collection("user")
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

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

func NewUserRepositoryDb() UserRepositoryDB {
	return UserRepositoryDB{}
}

