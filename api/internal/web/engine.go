package web

import (
	"github.com/SebasGA19/spAInews/api/internal/database"
	"github.com/gin-gonic/gin"
)

func NewEngine(database *database.Database) *gin.Engine {
	engine := gin.Default()
	controller := NewController(database)
	engine.POST(RegisterURI, controller.Register)
	engine.POST(LoginURI, controller.Login)
	return engine
}
