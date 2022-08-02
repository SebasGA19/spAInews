package main

import (
	"context"
	"github.com/SebasGA19/spAInews/api/internal/config"
	"github.com/SebasGA19/spAInews/api/internal/database"
	"github.com/SebasGA19/spAInews/api/internal/web"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

var (
	ConnectionURL string
	ListenAddress string
)

func init() {
	ConnectionURL = os.Getenv(config.MongoURL)
	ListenAddress = os.Getenv(config.ListenAddress)
}

/*
main: Set GIN_MODE=release for production
*/
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, connectionError := mongo.Connect(ctx, options.Client().ApplyURI(ConnectionURL))
	if connectionError != nil {
		log.Fatal(connectionError)
	}
	var readPref readpref.ReadPref
	pingError := client.Ping(context.Background(), &readPref)
	if pingError != nil {
		log.Fatal(pingError)
	}
	db := database.NewDatabase(client)
	engine := web.NewEngine(db)
	err := engine.Run(ListenAddress)
	log.Fatal(err)
}
