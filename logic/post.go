package logic

import (
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

func GetPostDetail(pid int64) (*models.Post, error) {
	return mysql.GetPostDetailById(pid)
}
