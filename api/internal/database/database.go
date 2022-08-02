package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Mongo *mongo.Client
}

func NewDatabase(client *mongo.Client) *Database {
	return &Database{
		Mongo: client,
	}
}
