package web

import (
	"github.com/SebasGA19/spAInews/auth-api/internal/controller"
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
	ChangePassword struct {
		OldPassword string `json:"old-password"`
		NewPassword string `json:"new-password"`
	}
	AccountInformationResponse struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	GetWordsResponse struct {
		Automatic bool     `json:"automatic"`
		Words     []string `json:"words"`
	}
	UpdateWords struct {
		Automatic bool     `json:"automatic"`
		Words     []string `json:"words"`
	}
	RequestPasswordReset struct {
		Email string `json:"email"`
	}
	ResetPassword struct {
		NewPassword string `json:"new-password"`
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
	confirmCode := ctx.GetHeader(ConfirmCodeHeader)
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
	session, found := ctx.Get(SessionVariable)
	if !found {
		ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(PermissionDeniedErrorCode))
		return
	}
	destroyError := backend.Controller.DestroySession(session.(string))
	if destroyError != nil {
		log.Print(destroyError)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	// Done
	ctx.Done()
}

func (backend *Backend) ChangePassword(ctx *gin.Context) {
	userId, found := ctx.Get(UserIdVariable)
	if !found {
		ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(PermissionDeniedErrorCode))
		return
	}
	var changePassword ChangePassword
	bindError := ctx.Bind(&changePassword)
	if bindError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	updatePasswordError := backend.Controller.ChangePassword(userId.(int), changePassword.OldPassword, changePassword.NewPassword)
	if updatePasswordError != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(PermissionDeniedErrorCode))
		return
	}
	// Done
	ctx.Done()
}

func (backend *Backend) AccountInformation(ctx *gin.Context) {
	userId, found := ctx.Get(UserIdVariable)
	if !found {
		ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(PermissionDeniedErrorCode))
		return
	}
	username, password, getInformationError := backend.Controller.Account(userId.(int))
	if getInformationError != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(PermissionDeniedErrorCode))
		return
	}
	ctx.JSON(http.StatusOK, AccountInformationResponse{
		Username: username,
		Email:    password,
	})
}

func (backend *Backend) GetWords(ctx *gin.Context) {
	userId, found := ctx.Get(UserIdVariable)
	if !found {
		ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(PermissionDeniedErrorCode))
		return
	}
	userWords, auto, getWordsError := backend.Controller.GetUserWords(userId.(int))
	if getWordsError != nil {
		log.Print(getWordsError)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	ctx.JSON(http.StatusOK, GetWordsResponse{
		Automatic: auto,
		Words:     userWords,
	})
	ctx.Done()
}

func (backend *Backend) PostWords(ctx *gin.Context) {
	userId, found := ctx.Get(UserIdVariable)
	if !found {
		ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(PermissionDeniedErrorCode))
		return
	}
	var updateWords UpdateWords
	bindError := ctx.Bind(&updateWords)
	if bindError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	setWordsError := backend.Controller.SetUserWords(userId.(int), updateWords.Words, updateWords.Automatic)
	if setWordsError != nil {
		log.Print(setWordsError)
		ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(PermissionDeniedErrorCode))
		return
	}
	ctx.Done()
}

func (backend *Backend) RequestResetPassword(ctx *gin.Context) {
	var requestPasswordReset RequestPasswordReset
	err := ctx.Bind(&requestPasswordReset)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	userId, queryUserIdError := backend.Controller.QueryUserIdByEmail(requestPasswordReset.Email)
	if queryUserIdError != nil {
		// Do not return errors to prevent email enumeration
		ctx.Done()
		return
	}
	addPendingError := backend.Controller.AddPendingReset(userId)
	if addPendingError != nil {
		log.Print(addPendingError)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	ctx.Done()
}

func (backend *Backend) ConfirmResetPassword(ctx *gin.Context) {
	resetCode := ctx.GetHeader(ConfirmCodeHeader)
	var resetPassword ResetPassword
	bindError := ctx.Bind(&resetPassword)
	if bindError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	resetError := backend.Controller.ResetPassword(resetCode, resetPassword.NewPassword)
	if resetError != nil {
		log.Print(resetError)
		ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(PermissionDeniedErrorCode))
		return
	}
	ctx.Done()
}
