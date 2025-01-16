package redis

import (
	"fmt"
	"go.uber.org/zap"
)

func SaveToken(uid int64, token string) (err error) {
	key := KeyUserToken + fmt.Sprintf("%d", uid)
	err = client.Set(ctx, key, token, 0).Err()
	if err != nil {
		zap.L().Error("SaveToken failed", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}

func GetTokenToRedis(uid int64) (token string, err error) {
	key := KeyUserToken + fmt.Sprintf("%d", uid)
	token, err = client.Get(ctx, key).Result()
	if err != nil {
		zap.L().Error("GetTokenToRedis failed", zap.String("key", key), zap.Error(err))
		return "", err
	}
	return token, err
}
