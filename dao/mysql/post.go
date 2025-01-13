package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"web-app/models"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := "insert into post (post_id,title,content,author_id,community_id) value(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, p.PostId, p.Title, p.Content, p.AuthorId, p.CommunityId)
	// 保存到数据库
	return
}

func GetPostDetailById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := "select * from post where post_id = ?"
	if err = db.Get(post, sqlStr, pid); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("where is no post in db")
			err = ErrorInvalidId
		}
	}
	return
}

func GetPostList() (postList []*models.Post, err error) {
	sqlStr := "select * from post"
	if err = db.Select(&postList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("where is no post no data db")
			err = nil
		}
	}
	return
}
