package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const UidKey = "userId"

var ErrorUserNotLogin = errors.New("用户未登录")

func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(UidKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
