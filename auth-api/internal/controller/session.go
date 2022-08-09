package controller

import (
	"github.com/SebasGA19/spAInews/auth-api/internal/common"
	"time"
)

func (c *Controller) CreateSession(userId int) (string, error) {
	key := common.RandString(32)
	err := c.Session.Set(c.ctx, key, userId, 24*30*time.Hour).Err()
	return key, err
}

func (c *Controller) DestroySession(session string) error {

	result := c.Session.Del(c.ctx, session)
	removedKeys, err := result.Result()
	if removedKeys != 1 {
		err = NoSessionFoundError
	}
	return err
}

func (c *Controller) QuerySession(session string) (int, error) {
	status := c.Session.Get(c.ctx, session)
	if err := status.Err(); err != nil {
		return -1, err
	}
	var userId int
	scanError := status.Scan(&userId)
	return userId, scanError
}
