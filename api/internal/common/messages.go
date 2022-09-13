package common

import (
	"bytes"
	"embed"
	"html/template"
)

var (
	//go:embed email-templates
	Templates embed.FS
)

func (s *SMTP) SendConfirmationEmail(code string, to ...string) error {
	if s.Dev {
		s.ConfirmationCode = code
		return nil
	}
	t := template.Must(template.ParseFS(Templates, "email-templates/email-confirmation.html"))
	out := bytes.NewBuffer(nil)
	executeError := t.Execute(out,
		struct {
			Code string
		}{
			Code: code,
		},
	)
	if executeError != nil {
		panic(executeError)
	}
	return s.SendEmail(to, "Confirm Email", string(out.Bytes()))
}

func (s *SMTP) SendResetCode(code string, to ...string) error {
	if s.Dev {
		s.ConfirmationCode = code
		return nil
	}
	t := template.Must(template.ParseFS(Templates, "email-templates/reset-password.html"))
	out := bytes.NewBuffer(nil)
	executeError := t.Execute(out,
		struct {
			Code string
		}{
			Code: code,
		},
	)
	if executeError != nil {
		panic(executeError)
	}
	return s.SendEmail(to, "Reset password", string(out.Bytes()))
}
