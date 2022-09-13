package main

import (
	"fmt"
	"github.com/SebasGA19/spAInews/api/internal/common"
	"github.com/SebasGA19/spAInews/api/internal/controller"
	"github.com/SebasGA19/spAInews/api/internal/web"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func init() {
	if len(os.Args) < 3 {
		_, _ = fmt.Fprintf(os.Stderr, "%s: LISTEN_ADDRESS CONFIG_FILE [CONFIG_FILE ...]", os.Args[0])
		os.Exit(1)
	}
	initSetup()
}

/*
main: Set GIN_MODE=release for production
*/
func main() {
	gin.SetMode(gin.ReleaseMode)
	mariadb := common.ConnectSQL(
		SQLUser, SQLPass, SQLHost, SQLPort, SQLDB, SQLExtra, SQLTCP,
	)
	mongoCollection := common.ConnectMongo(MongoUser, MongoPass, MongoHost, MongoPort, MongoDB).Collection("news")
	sessions := common.ConnectRedis(RedisHost, RedisPort, 0)
	registrations := common.ConnectRedis(RedisHost, RedisPort, 1)
	confirmEmails := common.ConnectRedis(RedisHost, RedisPort, 2)
	resetPasswords := common.ConnectRedis(RedisHost, RedisPort, 3)
	smtp := common.ConnectSMTP(SMTPFrom, SMTPUser, SMTPPass, SMTPHost, SMTPPort)
	c := controller.NewController(mariadb, mongoCollection, sessions, registrations, confirmEmails, resetPasswords, smtp)
	engine := web.NewEngine(c)
	log.Fatal(engine.Run(os.Args[1]))
}
