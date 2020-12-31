package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type UserRepositoryDB struct {
	client *mongo.Client
}

func (userRepositoryDB UserRepositoryDB) ById(id string) (*User, *error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := userRepositoryDB.client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer userRepositoryDB.client.Disconnect(ctx)

	var user User
	collection := userRepositoryDB.client.Database("countrycheck").Collection("user")
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	return &user, nil
}

func (userRepositoryDB UserRepositoryDB) Save(user User) (*mongo.InsertOneResult, *error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := userRepositoryDB.client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer userRepositoryDB.client.Disconnect(ctx)

	hashPw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashPw)

	collection := userRepositoryDB.client.Database("countrycheck").Collection("user")
	result, err := collection.InsertOne(ctx, user)
	if err != nil {log.Fatal(err)}

	return result, nil
}

func NewUserRepositoryDb(dbClient *mongo.Client) UserRepositoryDB {
	return UserRepositoryDB{dbClient}
}