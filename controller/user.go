package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"web-app/logic"
	"web-app/models"
)

// SignupHandler 注册
func SignupHandler(c *gin.Context) {
	// 1、获取参数和参数校验
	req := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(req); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{"error_msg": "SignUp with invalid param"})
	}
	// 2、业务逻辑处理
	if err := logic.Signup(req); err != nil {
		return
	}
	// 3、返回响应
}
