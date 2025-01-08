package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"web-app/models"
)

const secret = "ligongxin"

// CheckUserExist 查询用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) form user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

func InsertUser(user *models.User) (err error) {
	// 密码加密
	user.Password = encryptPassword(user.Password)
	// 保存到数据库
	sqlStr := "insert into user(user_id,username,password) value(?,?,?)"
	_, err = db.Exec(sqlStr, user.UserId, user.Username, user.Password)
	return
}

// 加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))

}
