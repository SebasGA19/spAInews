package web

import (
	"github.com/SebasGA19/spAInews/api/internal/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type (
	Register struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (backend *Backend) Register(ctx *gin.Context) {
	var registerForm Register
	bindError := ctx.BindJSON(&registerForm)
	if bindError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	// Check username and email are available
	userAvailable, queryError := backend.Controller.UsernameAvailable(registerForm.Username)
	if queryError != nil {
		log.Print(queryError)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	if !userAvailable {
		ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(UsernameNotAvailableErrorCode))
		return
	}
	var emailAvailable bool
	emailAvailable, queryError = backend.Controller.EmailAvailable(registerForm.Email)
	if queryError != nil {
		log.Print(queryError)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	if !emailAvailable {
		ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(EmailNotAvailableErrorCode))
		return
	}
	// Create a pending registration entry and send it to the address
	addPendingEmailError := backend.Controller.AddPendingEmail(
		controller.RegistrationData{
			Username: registerForm.Username,
			Email:    registerForm.Email,
			Password: registerForm.Password,
		},
	)
	if addPendingEmailError != nil {
		log.Print(addPendingEmailError)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	// Done
	ctx.Done()
}

func (backend *Backend) ConfirmAccount(ctx *gin.Context) {
	confirmCode := ctx.GetHeader(ConfirmAccountCodeHeader)
	// Confirm account
	confirmationError := backend.Controller.ConfirmEmail(confirmCode)
	if confirmationError != nil {
		log.Print(confirmationError)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	// Done
	ctx.Done()
}

func (backend *Backend) Login(ctx *gin.Context) {

}
