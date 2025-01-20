package redis

import "web-app/models"

func GetPostIdsInOrder(p *models.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTime)
	// 如果order传的是score，就查score的key
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScore)
	}
	// 请求起止值
	start := (p.Page - 1) * p.Size
	end := p.Size*p.Page - 1
	return client.ZRevRange(ctx, key, start, end).Result()
}
