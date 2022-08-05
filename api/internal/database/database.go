package database

import (
	"database/sql"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Mongo *mongo.Client
	SQL   *sql.DB
}

func NewDatabase(mongoClient *mongo.Client, sqlDB *sql.DB) *Database {
	return &Database{
		Mongo: mongoClient,
		SQL:   sqlDB,
	}
}
