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
	InternalServerError       = "Internal server error"
	UsernameNotAvailableError = "Username not available"
	EmailNotAvailableError    = "Email not available"
)

var (
	Errors = map[int]string{
		InternalServerErrorCode:       InternalServerError,
		UsernameNotAvailableErrorCode: UsernameNotAvailableError,
		EmailNotAvailableErrorCode:    EmailNotAvailableError,
	}
)

func NewError(code int) Error {
	return Error{
		Code:    code,
		Message: Errors[code],
	}
}
