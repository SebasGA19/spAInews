package email

import (
	"bytes"
	"embed"
	"html/template"
)

var (
	//go:embed email-templates
	Templates embed.FS
)

func (e *Email) SendConfirmationEmail(code string, to ...string) error {
	if e.Dev {
		e.ConfirmationCode = code
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
	return e.SendEmail(out.Bytes(), to...)
}
