package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/SebasGA19/spAInews/auth-api/internal/config"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/smtp"
	"os"
)

func mariaDB() *sql.DB {
	mariaURL := os.Getenv(config.MariaURL)
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

func redisClients() (sessions, registrations, confirmEmails, resetPasswords *redis.Client) {
	redisAddress := os.Getenv(config.RedisAddress)
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
	resetPasswords = redis.NewClient(
		&redis.Options{
			Addr: redisAddress,
			DB:   3,
		},
	)
	if _, err := resetPasswords.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err)
	}
	return sessions, registrations, confirmEmails, resetPasswords
}

func emailSettings() (address, from string, auth smtp.Auth) {
	auth = smtp.PlainAuth(
		"",
		os.Getenv(config.SMTPUsername),
		os.Getenv(config.SMTPPassword),
		os.Getenv(config.SMTPHost),
	)
	return fmt.Sprintf("%s:%s", os.Getenv(config.SMTPHost), os.Getenv(config.SMTPPort)), os.Getenv(config.SMTPFrom), auth
}
