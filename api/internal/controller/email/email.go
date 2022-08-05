package email

import "net/smtp"

type Email struct {
	SMTPAddress string
	From        string
	Auth        smtp.Auth
	// Dev Only
	Dev              bool
	ConfirmationCode string
}

func NewEmail(address, from string, auth smtp.Auth) *Email {
	return &Email{
		SMTPAddress: address,
		From:        from,
		Auth:        auth,
		Dev:         address == ":",
	}
}
