package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	ConnectionString string
	options          *options.ClientOptions
	MongoContext     context.Context
}

type IConnection interface {
	Connect() (*mongo.Client, error)
}

func (c *Connection) Connect() (*mongo.Client, error) {
	c.options = options.Client().ApplyURI(c.ConnectionString)
	client, err := mongo.Connect(c.MongoContext, c.options)
	if err != nil {
		return nil, err
	}

	return client, nil
}