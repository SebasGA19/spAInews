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
	SessionResponse struct {
		Session string `json:"session"`
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

func (backend *Backend) Session(ctx *gin.Context) {
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(CredentialsNotSubmitted))
		return
	}
	// Login
	id, loginError := backend.Controller.Login(username, password)
	if loginError != nil {
		log.Print(loginError)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	if id < 1 {
		ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(InvalidCredentialsErrorCode))
		return
	}
	// Create session
	sessionCode, createSessionError := backend.Controller.CreateSession(id)
	if createSessionError != nil {
		log.Print(createSessionError)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	// JSON
	ctx.JSON(http.StatusOK,
		SessionResponse{
			Session: sessionCode,
		},
	)
}

func (backend *Backend) DeleteSession(ctx *gin.Context) {
	session := ctx.GetHeader(SessionHeader)
	destroyError := backend.Controller.DestroySession(session)
	if destroyError != nil {
		log.Print(destroyError)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	// Done
	ctx.Done()
}
