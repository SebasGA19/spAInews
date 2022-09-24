package controller

import (
	"context"
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

func (c *Controller) AddPendingRegistration(registrationData RegistrationData) error {
	code := common.RandString(32)
	registrationDataJson, marshalError := json.Marshal(registrationData)
	if marshalError != nil {
		return marshalError
	}
	status := c.RedisRegistrations.Set(context.Background(), code, registrationDataJson, 24*time.Hour)
	if err := status.Err(); err != nil {
		return err
	}
	return c.SMTP.SendConfirmationRegistration(code, registrationData.Email)
}

func (c *Controller) ConfirmEmail(code string) error {
	cmd := c.RedisRegistrations.GetDel(context.Background(), code)
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
