package web

import "github.com/SebasGA19/spAInews/api/internal/database"

type Controller struct {
	DB *database.Database
}

func NewController(db *database.Database) *Controller {
	return &Controller{
		DB: db,
	}
}
