package web

type (
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

const (
	SessionHeader            = "Session"
	ConfirmAccountCodeHeader = "Confirm-Code"
)

const (
	RegisterURI       = "/register"
	ConfirmAccountURI = "/confirm"
	LoginURI          = "/login"
)

const (
	InternalServerErrorCode = iota
	UsernameNotAvailableErrorCode
	EmailNotAvailableErrorCode
	CredentialsNotSubmitted
	InvalidCredentialsErrorCode
	InternalServerError          = "Internal server error"
	UsernameNotAvailableError    = "Username not available"
	EmailNotAvailableError       = "Email not available"
	CredentialsNotSubmittedError = "Credentials not submitted"
	InvalidCredentialsError      = "invalid username or password"
)

var (
	Errors = map[int]string{
		InternalServerErrorCode:       InternalServerError,
		UsernameNotAvailableErrorCode: UsernameNotAvailableError,
		EmailNotAvailableErrorCode:    EmailNotAvailableError,
		CredentialsNotSubmitted:       CredentialsNotSubmittedError,
		InvalidCredentialsErrorCode:   InvalidCredentialsError,
	}
)

func NewError(code int) Error {
	return Error{
		Code:    code,
		Message: Errors[code],
	}
}
