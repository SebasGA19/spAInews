package controller

import (
	"encoding/json"
	"github.com/SebasGA19/spAInews/api/internal/common"
	"time"
)

type (
	RegistrationData struct {
		Username string
		Email    string
		Password string
	}
)

func (c *Controller) AddPendingEmail(registrationData RegistrationData) error {
	code := common.RandString(32)
	registrationDataJson, marshalError := json.Marshal(registrationData)
	if marshalError != nil {
		return marshalError
	}
	status := c.PendingEmails.Set(c.ctx, code, registrationDataJson, 24*time.Hour)
	if err := status.Err(); err != nil {
		return err
	}
	return c.Email.SendConfirmationEmail(code)
}

func (c *Controller) ConfirmEmail(code string) error {
	cmd := c.PendingEmails.GetDel(c.ctx, code)
	if err := cmd.Err(); err != nil {
		return err
	}
	var registrationData RegistrationData
	scanError := cmd.Scan(&registrationData)
	if scanError != nil {
		return scanError
	}
	return c.RegisterUser(registrationData.Username, registrationData.Email, registrationData.Password)
}
