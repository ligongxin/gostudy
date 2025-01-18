package redis

import (
	"go.uber.org/zap"
)

func SaveToken(uid, token string) (err error) {
	err = client.Set(ctx, getRedisKey(KeyUserToken)+uid, token, 0).Err()
	if err != nil {
		zap.L().Error("SaveToken failed", zap.String("key", getRedisKey(KeyUserToken)+uid), zap.Error(err))
		return err
	}
	return nil
}

func GetTokenToRedis(uid string) (token string, err error) {
	token, err = client.Get(ctx, getRedisKey(KeyUserToken)+uid).Result()
	if err != nil {
		zap.L().Error("GetTokenToRedis failed", zap.String("key", getRedisKey(KeyUserToken)+uid), zap.Error(err))
		return "", err
	}
	return token, err
}
