package logic

import (
	"fmt"
	"web-app/dao/mysql"
	"web-app/dao/redis"
	"web-app/models"
	"web-app/pkg/jwt"
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

func Login(p *models.ParamLogin) (data *models.ResponseLogin, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err = mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成token
	var token string
	token, err = jwt.GenToken(user.UserId)
	if err != nil {
		return
	}
	data = &models.ResponseLogin{
		UserId: user.UserId,
		Token:  token,
	}

	// 存token到redis
	if err = redis.SaveToken(user.UserId, token); err != nil {
		return nil, err
	}
	return data, nil
}
