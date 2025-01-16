package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web-app/logic"
	"web-app/models"
)

// 投票
func PostVoteController(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			zap.L().Error("PostVoteController| param error", zap.Error(err))
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
		zap.L().Error("PostVoteController| CodeInvalidParam", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 获取当前用户userid
	userId, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeUserNotLogin)
		return
	}
	if err := logic.VoteForPost(userId, p); err != nil {

		return
	}
	ResponseSuccess(c, nil)
}
