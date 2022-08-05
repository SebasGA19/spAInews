package main

import (
	"context"
	"database/sql"
	"github.com/SebasGA19/spAInews/api/internal/config"
	"github.com/SebasGA19/spAInews/api/internal/database"
	"github.com/SebasGA19/spAInews/api/internal/web"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MariaURL      string
	MongoURL      string
	ListenAddress string
)

func init() {
	MariaURL = os.Getenv(config.MariaURL)
	MongoURL = os.Getenv(config.MongoURL)
	ListenAddress = os.Getenv(config.ListenAddress)
}

/*
main: Set GIN_MODE=release for production
*/
func main() {
	sqlDB, err := sql.Open("mysql", MariaURL)
	if err != nil {
		log.Fatal(err)
	}
	pingError := sqlDB.Ping()
	if pingError != nil {
		log.Fatal(pingError)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, connectionError := mongo.Connect(ctx, options.Client().ApplyURI(MongoURL))
	if connectionError != nil {
		log.Fatal(connectionError)
	}
	var readPref readpref.ReadPref
	pingError = mongoClient.Ping(context.Background(), &readPref)
	if pingError != nil {
		log.Fatal(pingError)
	}
	db := database.NewDatabase(mongoClient, sqlDB)
	engine := web.NewEngine(db)
	log.Fatal(engine.Run(ListenAddress))
}
