package mysql

import "web-app/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := "insert into post (post_id,title,content,author_id,community_id) value(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, p.PostId, p.Title, p.Content, p.AuthorId, p.CommunityId)
	// 保存到数据库
	return
}
