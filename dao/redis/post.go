package redis

import (
	"web-app/models"
)

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

func GetPostVoteData(ids []string) (data []int64) {
	data = make([]int64, 0, len(ids))
	//统计每篇帖子的投票数量
	for _, postId := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + postId)
		v := client.ZCount(ctx, key, "1", "1").Val()
		data = append(data, v)
	}
	return
}

func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	return nil, nil
}
