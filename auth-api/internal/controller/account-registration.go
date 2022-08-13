package controller

import (
	"encoding/json"
	"github.com/SebasGA19/spAInews/auth-api/internal/common"
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
	status := c.RedisRegistrations.Set(c.ctx, code, registrationDataJson, 24*time.Hour)
	if err := status.Err(); err != nil {
		return err
	}
	return c.Email.SendConfirmationEmail(code)
}

func (c *Controller) ConfirmEmail(code string) error {
	cmd := c.RedisRegistrations.GetDel(c.ctx, code)
	if err := cmd.Err(); err != nil {
		return err
	}
	rawRegistrationData, bytesError := cmd.Bytes()
	if bytesError != nil {
		return bytesError
	}
	var registrationData RegistrationData
	unmarshallError := json.Unmarshal(rawRegistrationData, &registrationData)
	if unmarshallError != nil {
		return unmarshallError
	}
	return c.RegisterUser(registrationData.Username, registrationData.Email, registrationData.Password)
}
