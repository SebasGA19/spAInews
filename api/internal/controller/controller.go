package controller

import (
	"context"
	"database/sql"
	"github.com/SebasGA19/spAInews/api/internal/controller/email"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
	Mongo         *mongo.Client
	SQL           *sql.DB
	Session       *redis.Client
	PendingEmails *redis.Client
	Email         *email.Email
	ctx           context.Context
}

func (c *Controller) Close() error {
	_ = c.Mongo.Disconnect(c.ctx)
	_ = c.SQL.Close()
	_ = c.Session.Close()
	return c.PendingEmails.Close()
}

func NewController(
	mongoClient *mongo.Client,
	sqlDB *sql.DB,
	session *redis.Client,
	pendingEmails *redis.Client,
	emailClient *email.Email,
) *Controller {
	return &Controller{
		Mongo:         mongoClient,
		SQL:           sqlDB,
		Session:       session,
		PendingEmails: pendingEmails,
		Email:         emailClient,
		ctx:           context.Background(),
	}
}
