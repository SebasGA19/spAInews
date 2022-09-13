package controller

import (
	"encoding/json"
	"github.com/SebasGA19/spAInews/api/internal/common"
	"time"
)

type (
	UpdateEmail struct {
		UserId   int
		Password string
		NewEmail string
	}
)

func (c *Controller) AddUpdateEmail(updateEmail UpdateEmail) error {
	resetCode := common.RandString(32)
	updateEmailJson, marshalError := json.Marshal(updateEmail)
	if marshalError != nil {
		return marshalError
	}
	cmd := c.RedisConfirmEmails.Set(c.ctx, resetCode, updateEmailJson, 24*time.Hour)
	if err := cmd.Err(); err != nil {
		return err
	}
	return c.SMTP.SendConfirmationEmail(resetCode, updateEmail.NewEmail)
}

func (c *Controller) UpdateEmail(resetCode string) error {
	cmd := c.RedisConfirmEmails.GetDel(c.ctx, resetCode)
	if err := cmd.Err(); err != nil {
		return err
	}
	var rawUpdateEmail []byte
	scanError := cmd.Scan(&rawUpdateEmail)
	if scanError != nil {
		return scanError
	}
	var updateEmail UpdateEmail
	unmarshalError := json.Unmarshal(rawUpdateEmail, &updateEmail)
	if unmarshalError != nil {
		return unmarshalError
	}
	return c.ChangeEmail(updateEmail.UserId, updateEmail.NewEmail, updateEmail.Password)
}
