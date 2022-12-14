package web

type (
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

const (
	UserIdVariable  = "UserId"
	SessionVariable = "Session"
)

const (
	SessionHeader     = "Session"
	ConfirmCodeHeader = "Confirm-Code"
)

const (
	RegisterURI             = "/register"
	ConfirmRegistrationURI  = "/confirm/registration"
	SessionURI              = "/session"
	PasswordURI             = "/password"
	AccountURI              = "/account"
	WordsURI                = "/words"
	ResetPasswordURI        = "/reset/password"
	ConfirmResetPasswordURI = "/confirm/reset/password"
	EmailURI                = "/email"
	ConfirmUpdateEmailURI   = "/confirm/email"
	NewsPageVariable        = "page"
	NewsURI                 = "/news/:" + NewsPageVariable
	SearchNewsURI           = "/search/:" + NewsPageVariable
)

const (
	InternalServerErrorCode = iota
	UsernameNotAvailableErrorCode
	EmailNotAvailableErrorCode
	CredentialsNotSubmitted
	InvalidCredentialsErrorCode
	PermissionDeniedErrorCode
	InternalServerError          = "Internal server error"
	UsernameNotAvailableError    = "Username not available"
	EmailNotAvailableError       = "Email not available"
	CredentialsNotSubmittedError = "Credentials not submitted"
	InvalidCredentialsError      = "Invalid username or password"
	PermissionDeniedError        = "Permission denied"
)

var (
	Errors = map[int]string{
		InternalServerErrorCode:       InternalServerError,
		UsernameNotAvailableErrorCode: UsernameNotAvailableError,
		EmailNotAvailableErrorCode:    EmailNotAvailableError,
		CredentialsNotSubmitted:       CredentialsNotSubmittedError,
		InvalidCredentialsErrorCode:   InvalidCredentialsError,
		PermissionDeniedErrorCode:     PermissionDeniedError,
	}
)

func NewError(code int) Error {
	return Error{
		Code:    code,
		Message: Errors[code],
	}
}
