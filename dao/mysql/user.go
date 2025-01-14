package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"go.uber.org/zap"
	"web-app/models"
)

const secret = "gong"

// CheckUserExist 查询用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *models.User) (err error) {
	// 密码加密
	user.Password = encryptPassword(user.Password)
	// 保存到数据库
	sqlStr := "insert into user (user_id,username,password) value(?,?,?)"
	_, err = db.Exec(sqlStr, user.UserId, user.Username, user.Password)
	return
}

// 加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))

}

func Login(user *models.User) (err error) {
	// 存起旧密码
	oPassword := user.Password
	sqlStr := "select user_id,username,password from user where username=?"
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err //查找出错
	}
	// 判断密码
	if user.Password != encryptPassword(oPassword) {
		return ErrorInvalidPassword
	}
	return
}

func GetUserById(id int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select * from user where user_id = ?`
	if err = db.Get(user, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("where is no user in db")
			err = ErrorInvalidUserId
		}
	}
	return
}
