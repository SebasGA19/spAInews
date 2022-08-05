package web

import (
	"github.com/SebasGA19/spAInews/api/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewEngine(c *controller.Controller) *gin.Engine {
	engine := gin.Default()
	backend := NewBackend(c)
	engine.PUT(RegisterURI, backend.Register)
	engine.POST(ConfirmAccountURI, backend.ConfirmAccount)
	engine.GET(SessionURI, backend.Session)
	engine.DELETE(SessionURI, backend.DeleteSession)
	return engine
}
