package mysql

import "errors"

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
