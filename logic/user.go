package logic

import (
	"web-app/dao/mysql"
	"web-app/models"
)

// Signup 注册逻辑
func Signup(req *models.ParamSignUp) (err error) {
	// 1、判断 用户是否存在
	if err := mysql.CheckUserExist(req.Username); err != nil {
		return err
	}

	// 2、生成userid

	// 3、保存进数据库

	return
}
