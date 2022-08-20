package controller

import (
	"github.com/SebasGA19/spAInews/api/internal/common"
	"time"
)

func (c *Controller) AddPendingReset(userId int) error {
	resetCode := common.RandString(32)
	cmd := c.RedisResetPasswords.Set(c.ctx, resetCode, userId, 6*time.Hour)
	if err := cmd.Err(); err != nil {
		return err
	}
	return c.Email.SendResetCode(resetCode)
}

func (c Controller) ResetPassword(resetCode, newPassword string) error {
	cmd := c.RedisResetPasswords.GetDel(c.ctx, resetCode)
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
