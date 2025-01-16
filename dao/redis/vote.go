package redis

import (
	"errors"
	"time"
	"web-app/models"
)

const (
	oneWeekOnSecond = 7 * 24 * 3600
)

var ErrorVoteTimeExPried = errors.New("投票时间已过")

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
	// 获取帖子的分数
	ol := client.ZScore(ctx, getRedisKey(KeyPostScore), postId).Val()

}
