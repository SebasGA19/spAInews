package main

import (
	"github.com/SebasGA19/spAInews/auth-api/internal/config"
	"github.com/SebasGA19/spAInews/auth-api/internal/controller"
	"github.com/SebasGA19/spAInews/auth-api/internal/controller/email"
	"github.com/SebasGA19/spAInews/auth-api/internal/web"
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
	sessions, registrations, confirmEmails, resetPasswords := redisClients()
	e := email.NewEmail(emailSettings())
	db := controller.NewController(sqlDB, sessions, registrations, confirmEmails, resetPasswords, e)
	engine := web.NewEngine(db)
	log.Fatal(engine.Run(os.Getenv(config.ListenAddress)))
}
