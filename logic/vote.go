package logic

import (
	"go.uber.org/zap"
	"strconv"
	"web-app/dao/redis"
	"web-app/models"
)

func VoteForPost(userId int64, p *models.ParamVoteData) (err error) {
	if err = redis.PostVote(strconv.Itoa(int(userId)), p.PostID, float64(p.Direction)); err != nil {
		zap.L().Error("VoteForPost|redis.PostVote err ", zap.Error(err))
		return err
	}
	return
}
