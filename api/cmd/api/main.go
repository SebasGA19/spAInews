package main

import (
	"github.com/SebasGA19/spAInews/api/internal/config"
	"github.com/SebasGA19/spAInews/api/internal/controller"
	"github.com/SebasGA19/spAInews/api/internal/controller/email"
	"github.com/SebasGA19/spAInews/api/internal/web"
	"log"
	"os"
)

func init() {
}

/*
main: Set GIN_MODE=release for production
*/
func main() {
	mariadb := mariaDB()
	mongoCollection := mongoSettings()
	sessions, registrations, confirmEmails, resetPasswords := redisClients()
	e := email.NewEmail(emailSettings())
	db := controller.NewController(mariadb, mongoCollection, sessions, registrations, confirmEmails, resetPasswords, e)
	engine := web.NewEngine(db)
	log.Fatal(engine.Run(os.Getenv(config.ListenAddress)))
}