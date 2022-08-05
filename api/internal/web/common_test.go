package web

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/SebasGA19/spAInews/api/internal/controller"
	"github.com/SebasGA19/spAInews/api/internal/controller/email"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"net/http/httptest"
	"net/smtp"
	"time"
)

func mariaDB() *sql.DB {
	mariaURL := "spainews:spainews@tcp(127.0.0.1:3306)/spainews?parseTime=true"
	sqlDB, err := sql.Open("mysql", mariaURL)
	if err != nil {
		log.Fatal(err)
	}
	pingError := sqlDB.Ping()
	if pingError != nil {
		log.Fatal(pingError)
	}
	return sqlDB
}

func mongoDB() *mongo.Client {
	mongoURL := "mongodb://api:api@127.0.0.1:27017/spainews"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, connectionError := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if connectionError != nil {
		log.Fatal(connectionError)
	}
	var readPref readpref.ReadPref
	pingError := mongoClient.Ping(context.Background(), &readPref)
	if pingError != nil {
		log.Fatal(pingError)
	}
	return mongoClient
}

func redisClients() (sessions *redis.Client, pendingEmails *redis.Client) {
	redisAddress := "127.0.0.1:6379"
	sessions = redis.NewClient(
		&redis.Options{
			Addr: redisAddress,
			DB:   0,
		},
	)
	if _, err := sessions.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err)
	}
	pendingEmails = redis.NewClient(
		&redis.Options{
			Addr: redisAddress,
			DB:   1,
		},
	)
	if _, err := pendingEmails.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err)
	}
	return sessions, pendingEmails
}

func emailSettings() (address, from string, auth smtp.Auth) {
	auth = smtp.PlainAuth(
		"",
		"dashboard@dev-spainews.co",
		"password",
		"",
	)
	return fmt.Sprintf("%s:%s", "", ""), "dashboard@dev-spainews.co", auth
}

func newTestController() *controller.Controller {
	sqlDB := mariaDB()
	mongoClient := mongoDB()
	sessions, pendingEmails := redisClients()
	e := email.NewEmail(emailSettings())
	return controller.NewController(mongoClient, sqlDB, sessions, pendingEmails, e)
}

func getRequest(uri string, header http.Header, jsonPayload any) *http.Request {
	buffer, marshalError := json.Marshal(jsonPayload)
	if marshalError != nil {
		panic(marshalError)
	}
	body := bytes.NewBuffer(buffer)
	req := httptest.NewRequest(http.MethodGet, uri, body)
	if header != nil {
		req.Header = header
	}
	return req
}

func postRequest(uri string, header http.Header, jsonPayload any) *http.Request {
	buffer, marshalError := json.Marshal(jsonPayload)
	if marshalError != nil {
		panic(marshalError)
	}
	body := bytes.NewBuffer(buffer)
	req := httptest.NewRequest(http.MethodPost, uri, body)
	if header != nil {
		req.Header = header
	}
	return req
}

func deleteRequest(uri string, header http.Header, jsonPayload any) *http.Request {
	buffer, marshalError := json.Marshal(jsonPayload)
	if marshalError != nil {
		panic(marshalError)
	}
	body := bytes.NewBuffer(buffer)
	req := httptest.NewRequest(http.MethodDelete, uri, body)
	if header != nil {
		req.Header = header
	}
	return req
}

func putRequest(uri string, header http.Header, jsonPayload any) *http.Request {
	buffer, marshalError := json.Marshal(jsonPayload)
	if marshalError != nil {
		panic(marshalError)
	}
	body := bytes.NewBuffer(buffer)
	req := httptest.NewRequest(http.MethodPut, uri, body)
	if header != nil {
		req.Header = header
	}
	return req
}
