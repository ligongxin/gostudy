package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		// 判断错误是不是validator的错误
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusOK, gin.H{"error_msg": errs.Translate(trans)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"error_msg": "SignUp with invalid param"})
	}
	// 2、业务逻辑处理
	if err := logic.Signup(req); err != nil {
		zap.L().Error("SignUp failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	// 3、返回响应
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {
	// 请求参数机校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusOK, gin.H{"message": errs.Translate(trans)})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "参数错误"})
		return
	}
	// 登录业务逻辑处理
	err := logic.Login(p)
	if err != nil {
		zap.L().Error("login failed", zap.Error(err))
		// 判断是不是数据库错误
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	// 返回响应
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
