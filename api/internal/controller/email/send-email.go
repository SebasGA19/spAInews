//go:build !dev

package email

import (
	"fmt"
	"net/smtp"
)

func (e *Email) SendEmail(message []byte, to ...string) error {
	if e.Dev {
		fmt.Println(string(message))
		return nil
	}
	return smtp.SendMail(e.SMTPAddress, e.Auth, e.From, to, message)
}
