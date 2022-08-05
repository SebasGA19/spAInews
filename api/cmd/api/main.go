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
	sqlDB := mariaDB()
	mongoClient := mongoDB()
	sessions, pendingEmails := redisClients()
	e := email.NewEmail(emailSettings())
	db := controller.NewController(mongoClient, sqlDB, sessions, pendingEmails, e)
	engine := web.NewEngine(db)
	log.Fatal(engine.Run(os.Getenv(config.ListenAddress)))
}
