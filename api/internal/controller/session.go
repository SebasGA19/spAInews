package controller

import (
	"context"
	"github.com/SebasGA19/spAInews/api/internal/common"
	"time"
)

func (c *Controller) CreateSession(userId int) (string, error) {
	key := common.RandString(32)
	err := c.RedisSession.Set(context.Background(), key, userId, 24*30*time.Hour).Err()
	return key, err
}

func (c *Controller) DestroySession(session string) error {
	result := c.RedisSession.Del(context.Background(), session)
	removedKeys, err := result.Result()
	if removedKeys != 1 {
		err = NoSessionFoundError
	}
	return err
}

func (c *Controller) QuerySession(session string) (int, error) {
	status := c.RedisSession.Get(context.Background(), session)
	if err := status.Err(); err != nil {
		return -1, err
	}
	var userId int
	scanError := status.Scan(&userId)
	return userId, scanError
}
