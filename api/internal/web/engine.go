package web

import (
	"github.com/SebasGA19/spAInews/api/internal/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NewEngine(c *controller.Controller) *gin.Engine {
	engine := gin.Default()
	engine.Use(
		cors.New(
			cors.Config{
				AllowOrigins: []string{"*"},
				AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
				AllowHeaders: []string{
					"Origin",
					"Content-Type",
					"Content-Length",
					"Authorization",
					SessionHeader,
					ConfirmAccountCodeHeader,
				},
				ExposeHeaders:    []string{"Content-Length", "Content-Type"},
				AllowCredentials: true,
				AllowOriginFunc: func(origin string) bool {
					return true
				},
				MaxAge: 12 * time.Hour,
			},
		),
	)
	backend := NewBackend(c)
	engine.PUT(RegisterURI, backend.Register)
	engine.POST(ConfirmAccountURI, backend.ConfirmAccount)
	engine.GET(SessionURI, backend.Session)
	engine.DELETE(SessionURI, backend.DeleteSession)
	return engine
}
