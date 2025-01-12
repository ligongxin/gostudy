package logic

import (
	"web-app/dao/mysql"
	"web-app/models"
)

// GetCommunityList 查询所有的数据
func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}
