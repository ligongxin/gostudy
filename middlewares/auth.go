package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
	"web-app/controller"
	"web-app/pkg/jwt"
)

// JwtAuthDiddleWare 基于中间件认证
func JwtAuthDiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx
		// 这里的具体实现方式要依据你的实际业务情况决定
		token := c.Request.Header.Get("Authorization")
		// 判断是否存在
		if token == "" {
			//c.JSON(http.StatusOK, gin.H{"message": "用户未登录"})
			controller.ResponseError(c, controller.CodeUserNotLogin)
			c.Abort()
			return
		}
		// 判断token的格式
		parts := strings.SplitN(token, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.COdeInvalidToken)
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.COdeInvalidToken)
			c.Abort()
			return
		}
		c.Set(controller.UidKey, mc.UserID)
		//c.Set("username", mc.Username)
		c.Next()
	}
}
