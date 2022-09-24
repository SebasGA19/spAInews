package controller

import (
	"database/sql"
	"github.com/SebasGA19/spAInews/api/internal/common"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
	SQL                 *sql.DB
	Mongo               *mongo.Collection
	RedisSession        *redis.Client
	RedisRegistrations  *redis.Client
	RedisConfirmEmails  *redis.Client
	RedisResetPasswords *redis.Client
	SMTP                *common.SMTP
}

func (c *Controller) Close() error {
	_ = c.SQL.Close()
	_ = c.RedisSession.Close()
	return c.RedisRegistrations.Close()
}

func NewController(
	sqlDB *sql.DB,
	mongoClient *mongo.Collection,
	redisSession *redis.Client,
	redisRegistrations *redis.Client,
	redisConfirmEmails *redis.Client,
	redisResetPasswords *redis.Client,
	smtpClient *common.SMTP,
) *Controller {
	return &Controller{
		SQL:                 sqlDB,
		Mongo:               mongoClient,
		RedisSession:        redisSession,
		RedisRegistrations:  redisRegistrations,
		RedisConfirmEmails:  redisConfirmEmails,
		RedisResetPasswords: redisResetPasswords,
		SMTP:                smtpClient,
	}
}
