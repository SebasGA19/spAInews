package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (backend *Backend) AuthSession(ctx *gin.Context) {
	session := ctx.GetHeader(SessionHeader)
	userId, queryError := backend.Controller.QuerySession(session)
	if queryError != nil {
		log.Print(queryError)
		ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(PermissionDeniedErrorCode))
		return
	}
	ctx.Set(UserIdVariable, userId)
	ctx.Set(SessionVariable, session)
}
