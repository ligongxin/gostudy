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
