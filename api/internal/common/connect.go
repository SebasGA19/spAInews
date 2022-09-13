package common

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func ConnectSQL(username, password, host, port, database, extraConfigs string, tcp bool) *sql.DB {
	var url string
	if tcp {
		url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	} else {
		url = fmt.Sprintf("%s:%s@%s:%s/%s", username, password, host, port, database)
	}
	if len(extraConfigs) > 0 {
		url = fmt.Sprintf("%s?%s", url, extraConfigs)
	}
	conn, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	pingErr := conn.Ping()
	if pingErr != nil {
		panic(pingErr)
	}
	return conn
}

func ConnectRedis(host, port string, db int) *redis.Client {
	addr := fmt.Sprintf("%s:%s", host, port)
	client := redis.NewClient(
		&redis.Options{
			Addr: addr,
			DB:   db,
		},
	)
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}
	return client
}

func ConnectSMTP(from, username, password, host, port string, dev bool) *SMTP {
	addr := fmt.Sprintf("%s:%s", host, port)
	client, dialError := smtp.Dial(addr)
	if dialError != nil {
		panic(dialError)
	}
	defer client.Close()
	auth := sasl.NewPlainClient("", username, password)
	authError := client.Auth(auth)
	if authError != nil {
		panic(authError)
	}
	return &SMTP{
		Addr: addr,
		From: from,
		Auth: auth,
	}
}

func ConnectMongo(username, password, host, port, database string) *mongo.Database {
	mongoUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", username, password, host, port, database)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, connectionError := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if connectionError != nil {
		panic(connectionError)
	}
	var readPref readpref.ReadPref
	pingError := mongoClient.Ping(context.Background(), &readPref)
	if pingError != nil {
		panic(pingError)
	}
	return mongoClient.Database(database)
}
