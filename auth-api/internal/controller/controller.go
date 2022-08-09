package controller

import (
	"context"
	"database/sql"
	"github.com/SebasGA19/spAInews/auth-api/internal/controller/email"
	"github.com/go-redis/redis/v9"
)

type Controller struct {
	SQL           *sql.DB
	Session       *redis.Client
	PendingEmails *redis.Client
	Email         *email.Email
	ctx           context.Context
}

func (c *Controller) Close() error {
	_ = c.SQL.Close()
	_ = c.Session.Close()
	return c.PendingEmails.Close()
}

func NewController(
	sqlDB *sql.DB,
	session *redis.Client,
	pendingEmails *redis.Client,
	emailClient *email.Email,
) *Controller {
	return &Controller{
		SQL:           sqlDB,
		Session:       session,
		PendingEmails: pendingEmails,
		Email:         emailClient,
		ctx:           context.Background(),
	}
}
