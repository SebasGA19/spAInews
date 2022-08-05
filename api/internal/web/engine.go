package web

import (
	"github.com/SebasGA19/spAInews/api/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewEngine(database *controller.Controller) *gin.Engine {
	engine := gin.Default()
	c := NewController(database)
	engine.POST(RegisterURI, c.Register)
	engine.POST(LoginURI, c.Login)
	return engine
}
