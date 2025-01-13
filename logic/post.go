package logic

import (
	"go.uber.org/zap"
	"web-app/dao/mysql"
	"web-app/models"
	"web-app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//生成post_id
	p.PostId = snowflake.GenID()
	//创建帖子
	return mysql.CreatePost(p)
}

func GetPostDetail(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询帖子
	post, err := mysql.GetPostDetailById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailById failed", zap.Int64("post_id", pid), zap.Error(err))
		return
	}
	//查询帖子的作者
	author, err := mysql.GetUserById(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserById failed", zap.Int64("post.AuthorId", post.AuthorId), zap.Error(err))
		return
	}
	// 查询帖子所属社区详情
	community, err := mysql.GetCommunityDetailById(post.CommunityId)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("post.CommunityId", post.CommunityId), zap.Error(err))
		return
	}
	// 拼接数据
	data = &models.ApiPostDetail{
		Username:  author.Username,
		Post:      post,
		Community: community,
	}
	return
}
