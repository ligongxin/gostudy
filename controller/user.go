package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"web-app/dao/mysql"
	"web-app/logic"
	"web-app/models"
	"web-app/pkg/jwt"
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
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2、业务逻辑处理
	if err := logic.Signup(req); err != nil {
		zap.L().Error("SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3、返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {
	// 请求参数机校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 登录业务逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("login failed", zap.String("username", p.Username), zap.Error(err))
		// 判断是不是数据库错误
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 返回响应
	ResponseSuccess(c, token)
}

func GetTokenHandler(c *gin.Context) {
	uidStr, ok := c.GetQuery("uid")
	if !ok {
		zap.L().Error("GetTokenHandler query uid failed", zap.String("uid", uidStr))
		ResponseError(c, CodeInvalidParam)
		return
	}
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	aToken, rToken, err := jwt.GenTokenV1(uid)
	if err != nil {
		zap.L().Error("GetTokenHandler token生成错误", zap.String("uid", uidStr), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{"msg": "token生成错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"a_token": aToken, "r_token": rToken})
}
