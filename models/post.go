package models

import "time"

type Post struct {
	ID          int64     `json:"id" db:"id"`
	PostId      int64     `json:"post_id" db:"post_id"`                              // 帖子id
	Title       string    `json:"title" db:"title" binding:"required"`               // 标题
	Content     string    `json:"content" db:"content" binding:"required"`           // 内容
	AuthorId    int64     `json:"author_id" db:"author_id"`                          // 作者的用户id
	CommunityId int64     `json:"community_id" db:"community_id" binding:"required"` // 所属社区
	Status      int8      `json:"status" db:"status"`                                // 帖子状态
	CreateTime  time.Time `json:"create_time" db:"create_time"`                      // 创建时间
	UpdateTime  time.Time `json:"update_time" db:"update_time"`                      // 更新时间
}
