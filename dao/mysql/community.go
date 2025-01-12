package mysql

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"web-app/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select * from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		// 判断是否查询出空的
		fmt.Printf("%#v\n", communityList)
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetailById(id int64) (communityDetail *models.Community, err error) {
	communityDetail = new(models.Community)
	sqlStr := `select * from community where community_id=?`
	if err = db.Get(communityDetail, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("where is no community in db")
			err = ErrorInvalidId
		}
	}
	return communityDetail, err
}
