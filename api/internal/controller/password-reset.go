package controller

import (
	"context"
	"github.com/SebasGA19/spAInews/api/internal/common"
	"time"
)

func (c *Controller) AddPendingReset(userId int) error {
	resetCode := common.RandString(32)
	cmd := c.RedisResetPasswords.Set(context.Background(), resetCode, userId, 6*time.Hour)
	if err := cmd.Err(); err != nil {
		return err
	}
	_, email, err := c.Account(userId)
	if err != nil {
		return err
	}
	return c.SMTP.SendResetCode(resetCode, email)
}

func (c *Controller) ResetPassword(resetCode, newPassword string) error {
	cmd := c.RedisResetPasswords.GetDel(context.Background(), resetCode)
	if err := cmd.Err(); err != nil {
		return err
	}
	var userId int
	scanError := cmd.Scan(&userId)
	if scanError != nil {
		return scanError
	}
	return c.SQLResetPassword(userId, newPassword)
}
