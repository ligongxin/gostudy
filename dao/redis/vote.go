package redis

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"math"
	"time"
	"web-app/models"
)

const (
	oneWeekOnSecond = 7 * 24 * 3600
	PostScore       = 423
)

var (
	ErrorVoteTimeExPried = errors.New("投票时间已过")
	ErrorVoteRepeated    = errors.New("不能重复投票")
)

// 1、支持投赞成票（1）、反对票（-1）、取消投票（0）
/*
direction =1：
	1、原本没投票，现在投赞成票
	2、原本投反对，现在投赞成票
direction =0：
	1、原本投赞成，现在取消投票
	2、原本投反对，现在取消投票
direction =-1
	1、原本没投票，现在投反对票
	2、原本投赞成，现在投反对票

帖子时间超过7天，不能投票
*/
func PostVote(userId, postId string, value float64) error {
	// 获取帖子发布时间 当前帖子的时间是否超过7天
	postTime := client.ZScore(ctx, getRedisKey(KeyPostTime), postId).Val()
	if float64(time.Now().Unix())-postTime > oneWeekOnSecond {
		return ErrorVoteTimeExPried
	}
	// 获取用户的投票记录
	ov := client.ZScore(ctx, getRedisKey(KeyPostVotedZSetPrefix+postId), userId).Val()
	// 重复投票
	if ov == value {
		return ErrorVoteRepeated
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	//判断差值
	diff := math.Abs(value - ov)
	// 更新帖子的分数 更新
	pipe := client.TxPipeline()
	pipe.ZIncrBy(ctx, getRedisKey(KeyPostScore), PostScore*diff*op, postId)
	// 更新用户投票信息
	if value == 0 {
		pipe.ZRem(ctx, getRedisKey(KeyPostVotedZSetPrefix+postId), userId)
	} else {
		pipe.ZAdd(ctx, getRedisKey(KeyPostVotedZSetPrefix+postId), redis.Z{
			Score:  value,
			Member: userId,
		})
	}
	_, err := pipe.Exec(ctx)
	return err
}

// 插入redis
func CreatePostCache(p *models.Post) error {
	// 添加 帖子发布时间
	score := float64(time.Now().Unix())

	// 开启一个TxPipeline事务
	pipe := client.TxPipeline()
	pipe.ZAdd(ctx, getRedisKey(KeyPostTime), redis.Z{
		Score:  score,
		Member: p.PostId,
	})
	// 添加帖子的分数
	pipe.ZAdd(ctx, getRedisKey(KeyPostScore), redis.Z{
		Score:  score,
		Member: p.PostId,
	})
	//通过Exec函数提交redis事务
	_, err := pipe.Exec(ctx)
	return err
}
