package logic

import (
	"fmt"
	"web-app/dao/mysql"
	"web-app/models"
	"web-app/pkg/snowflake"
)

// Signup 注册逻辑
func Signup(req *models.ParamSignUp) (err error) {
	// 1、判断 用户是否存在
	if err := mysql.CheckUserExist(req.Username); err != nil {
		fmt.Println(err)
		return err
	}

	// 2、生成userid
	userId := snowflake.GenID()
	// 3、保存进数据库
	u := models.User{
		UserId:   userId,
		Username: req.Username,
		Password: req.Password,
	}
	return mysql.InsertUser(&u)
}

func Login(p *models.ParamLogin) (err error) {
	u := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(u); err != nil {
		return err
	}
	return
}
