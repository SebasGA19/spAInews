package common

import (
	"fmt"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"strings"
)

type SMTP struct {
	DashboardURL     string
	Dev              bool
	Addr             string
	From             string
	ConfirmationCode string
	Auth             sasl.Client
}

func (s *SMTP) SendEmail(recipients []string, subject, message string) error {
	if s.Dev {
		fmt.Println(message)
		return nil
	}
	to := recipients[0]
	if len(recipients) > 1 {
		to += ", " + strings.Join(recipients[1:], ", ")
	}
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";"
	body := strings.NewReader(fmt.Sprintf("To: %s\r\nSubject: %s\r\n%s\r\n%s\r\n", to, mime, subject, message))
	return smtp.SendMail(s.Addr, s.Auth, s.From, recipients, body)
}
