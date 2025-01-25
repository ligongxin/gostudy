package redis

import (
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
	"web-app/models"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := size*page - 1
	return client.ZRevRange(ctx, key, start, end).Result()
}

func GetPostIdsInOrder(p *models.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTime)
	// 如果order传的是score，就查score的key
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScore)
	}
	// 请求起止值
	//start := (p.Page - 1) * p.Size
	//end := p.Size*p.Page - 1
	//return client.ZRevRange(ctx, key, start, end).Result()
	return getIDsFormKey(key, p.Page, p.Size)
}

func GetPostVoteData(ids []string) (data []int64) {
	data = make([]int64, 0, len(ids))
	//统计每篇帖子的投票数量
	for _, post_id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + post_id)
		v := client.ZCount(ctx, key, "1", "1").Val()
		data = append(data, v)
	}
	return
}

// GetCommunityPostIDsInOrder 根据社区查帖子ids
func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	// 根据排序查找对于的redis的key
	orderKey := getRedisKey(KeyPostTime)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScore)
	}
	// 使用 zinterstore 把分区的帖子set与帖子分数的 zset 生成一个新的zset
	// 针对新的zset 按之前的逻辑取数据
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityId)))
	// 利用缓存key减少zinterstore执行的次数
	key := orderKey + strconv.Itoa(int(p.CommunityId))
	if client.Exists(ctx, key).Val() < 1 {
		//不存在，需要计算
		pipe := client.Pipeline()
		pipe.ZInterStore(ctx, key, &redis.ZStore{
			Keys:      []string{cKey, orderKey},
			Aggregate: "MAX",
		}) // zinterstore 计算
		pipe.Expire(ctx, key, 60*time.Second)
		_, err := pipe.Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	// 存在的话就直接根据key查询ids
	return getIDsFormKey(key, p.Page, p.Size)
}
