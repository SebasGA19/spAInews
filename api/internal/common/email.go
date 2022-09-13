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
	body := strings.NewReader(fmt.Sprintf("Subject: %s\r\n\r\n%s\r\n", subject, message))
	return smtp.SendMail(s.Addr, s.Auth, s.From, recipients, body)
}
