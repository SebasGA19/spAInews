package controller

import (
	"context"
	"database/sql"
	"github.com/SebasGA19/spAInews/auth-api/internal/controller/email"
	"github.com/go-redis/redis/v9"
)

type Controller struct {
	SQL                 *sql.DB
	RedisSession        *redis.Client
	RedisRegistrations  *redis.Client
	RedisConfirmEmails  *redis.Client
	RedisResetPasswords *redis.Client
	Email               *email.Email
	ctx                 context.Context
}

func (c *Controller) Close() error {
	_ = c.SQL.Close()
	_ = c.RedisSession.Close()
	return c.RedisRegistrations.Close()
}

func NewController(
	sqlDB *sql.DB,
	redisSession *redis.Client,
	redisRegistrations *redis.Client,
	redisConfirmEmails *redis.Client,
	redisResetPasswords *redis.Client,
	emailClient *email.Email,
) *Controller {
	return &Controller{
		SQL:                 sqlDB,
		RedisSession:        redisSession,
		RedisRegistrations:  redisRegistrations,
		RedisConfirmEmails:  redisConfirmEmails,
		RedisResetPasswords: redisResetPasswords,
		Email:               emailClient,
		ctx:                 context.Background(),
	}
}
