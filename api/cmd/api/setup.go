package main

import (
	"fmt"
	"github.com/SebasGA19/spAInews/api/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

var (
	SQLUser  string
	SQLPass  string
	SQLHost  string
	SQLPort  string
	SQLDB    string
	SQLExtra string
	SQLTCP   bool
)

var (
	MongoUser string
	MongoPass string
	MongoHost string
	MongoPort string
	MongoDB   string
)

var (
	RedisHost string
	RedisPort string
)

var (
	SMTPFrom string
	SMTPUser string
	SMTPPass string
	SMTPHost string
	SMTPPort string
)

func initSetup() {
	m, err := godotenv.Read(os.Args[1:]...)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	// SQL
	SQLUser = m[config.SQLUser]
	SQLPass = m[config.SQLPass]
	SQLHost = m[config.SQLHost]
	SQLPort = m[config.SQLPort]
	SQLDB = m[config.SQLDB]
	SQLExtra = m[config.SQLExtra]
	SQLTCP = m[config.SQLTCP] == "true"
	// Mongo
	MongoUser = m[config.MongoUser]
	MongoPass = m[config.MongoPass]
	MongoHost = m[config.MongoHost]
	MongoPort = m[config.MongoPort]
	MongoDB = m[config.MongoDB]
	// Redis
	RedisHost = m[config.RedisHost]
	RedisPort = m[config.RedisPort]
	// SMTP
	SMTPFrom = m[config.SMTPFrom]
	SMTPUser = m[config.SMTPUser]
	SMTPPass = m[config.SMTPPass]
	SMTPHost = m[config.SMTPHost]
	SMTPPort = m[config.SMTPPort]
}
