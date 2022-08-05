package email

import "net/smtp"

type Email struct {
	SMTPAddress string
	From        string
	Auth        smtp.Auth
	dev         bool
}

func NewEmail(address, from string, auth smtp.Auth) *Email {
	return &Email{
		SMTPAddress: address,
		From:        from,
		Auth:        auth,
		dev:         address == ":",
	}
}
