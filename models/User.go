package models

import (
	"context"
	"os"
	"simpleCrudMongoDB/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Username string `bson:"username"`
	Email string `bson:"email"`
	Address string `bson:"address"`
}

type IUser interface {
	CreateUser (data User) (*mongo.InsertOneResult, error)
	FindUser (filter interface{}) (*User, error)
	FindUsers () ([]*User, error)
	UpdateUser (filter interface{}, data User) (*User, error)
	Delete (filter interface{}) error
}

type DB struct {
	*db.Connection
}

var mongoContext, _ = context.WithTimeout(context.Background(), 30 * time.Second)
var userCollection *mongo.Collection

func (u *User) init() error {
	var dbConnection db.IConnection

	dbConnection = &db.Connection{
		ConnectionString: os.Getenv("MONGO_HOST"),
		MongoContext: mongoContext,
	}

	client, err := dbConnection.Connect()
	if err != nil {
		return err
	}

	userCollection = client.Database(os.Getenv("DB_NAME")).Collection("users")
	return nil
}

func (u *User) CreateUser(data User) (*mongo.InsertOneResult, error) {
	err := u.init()
	if err != nil {
		return nil, err
	}

	res, err := userCollection.InsertOne(mongoContext, data)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *User) FindUser (filter interface{}) (*User, error) {
	err := u.init()
	if err != nil {
		return nil, err
	}
	filter = bson.M{
		"username": filter.(User).Username,
	}
	res := userCollection.FindOne(mongoContext, filter)
	if err := res.Err(); err != nil {
		return nil, err
	}

	userData := new(User)
	err = res.Decode(userData)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (u *User) FindUsers() ([]*User, error) {
	panic("implement me")
}

func (u *User) UpdateUser (filter interface{}, data User) (*User, error) {
	err := u.init()
	if err != nil {
		return nil, err
	}
	filter = bson.M{
		"username": filter.(User).Username,
	}
	updateData := bson.M{
		"username": data.Username,
	}
	res := userCollection.FindOneAndUpdate(mongoContext, filter, bson.M{"$set": updateData})
	if err = res.Err(); err != nil {
		return nil, err
	}
	userData := new(User)
	err = res.Decode(userData)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (u *User) Delete (filter interface{}) error {
	err := u.init()
	if err != nil {
		return err
	}
	filter = bson.M{
		"username": filter.(User).Username,
	}
	res := userCollection.FindOneAndDelete(mongoContext, filter)
	if err = res.Err(); err != nil {
		return err
	}

	return nil
}