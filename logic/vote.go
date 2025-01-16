package logic

import (
	"strconv"
	"web-app/dao/redis"
	"web-app/models"
)

func VoteForPost(userId int64, p *models.ParamVoteData) (err error) {
	if err = redis.PostVote(strconv.Itoa(int(userId)), p.PostID, float64(p.Direction)); err != nil {
		return err
	}
	return
}
