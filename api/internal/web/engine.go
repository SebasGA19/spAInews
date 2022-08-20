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
					ConfirmCodeHeader,
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

	noAuthRequired := engine.Group("/")
	noAuthRequired.PUT(RegisterURI, backend.Register)
	noAuthRequired.POST(ConfirmRegistrationURI, backend.ConfirmAccount)
	noAuthRequired.GET(SessionURI, backend.Session)
	noAuthRequired.POST(ResetPasswordURI, backend.RequestResetPassword)
	noAuthRequired.POST(ConfirmResetPasswordURI, backend.ConfirmResetPassword)
	noAuthRequired.POST(ConfirmUpdateEmailURI, backend.ConfirmUpdateEmail)
	noAuthRequired.GET(NewsURI, backend.LatestNews)
	noAuthRequired.GET(SearchNewsURI, backend.SearchNews)

	authRequired := engine.Group("/", backend.AuthSession)
	authRequired.DELETE(SessionURI, backend.DeleteSession)
	authRequired.POST(PasswordURI, backend.ChangePassword)
	authRequired.GET(AccountURI, backend.AccountInformation)
	authRequired.GET(WordsURI, backend.GetWords)
	authRequired.POST(WordsURI, backend.PostWords)
	authRequired.POST(EmailURI, backend.RequestUpdateEmail)
	return engine
}
