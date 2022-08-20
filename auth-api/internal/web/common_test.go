package web

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/SebasGA19/spAInews/auth-api/internal/controller"
	"github.com/SebasGA19/spAInews/auth-api/internal/controller/email"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"net/http/httptest"
	"net/smtp"
	"net/url"
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

func redisClients() (sessions *redis.Client, registrations, confirmEmails, passwordResets *redis.Client) {
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
	registrations = redis.NewClient(
		&redis.Options{
			Addr: redisAddress,
			DB:   1,
		},
	)
	if _, err := registrations.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err)
	}
	confirmEmails = redis.NewClient(
		&redis.Options{
			Addr: redisAddress,
			DB:   2,
		},
	)
	if _, err := confirmEmails.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err)
	}
	passwordResets = redis.NewClient(
		&redis.Options{
			Addr: redisAddress,
			DB:   3,
		},
	)
	if _, err := passwordResets.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err)
	}
	return sessions, registrations, confirmEmails, passwordResets
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
	mongoCollection := mongoSettings()
	sessions, registrations, confirmEmails, passwordResets := redisClients()
	e := email.NewEmail(emailSettings())
	return controller.NewController(sqlDB, mongoCollection, sessions, registrations, confirmEmails, passwordResets, e)
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

func mongoSettings() *mongo.Collection {
	mongoURL := "mongodb://spainews:spainews@127.0.0.1:27017/spainews"
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
	parseUrl, parseError := url.Parse(mongoURL)
	if parseError != nil {
		log.Fatal(parseError)
	}
	database := mongoClient.Database(parseUrl.RequestURI()[1:])
	return database.Collection("news")
}
