package mysql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strings"
	"web-app/models"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post (post_id,title,content,author_id,community_id) value(?,?,?,?,?)`
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

func GetPostList(page, size int64) (postList []*models.Post, err error) {
	sqlStr := "select * from post limit ?,?"
	if err = db.Select(&postList, sqlStr, page-1, size); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("where is no post no data db")
			err = nil
		}
	}
	return
}

func GetPostListByIds(ids []string) (postList []*models.Post, err error) {
	// 使用sqlx in 查询
	sqlStr := ` select * from post
		where post_id in (?)
		order by FIND_IN_SET(post_id,?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
