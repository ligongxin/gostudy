package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web-app/logic"
)

// CommunityHandler 社区列表
func CommunityHandler(c *gin.Context) {
	//// 查询到所有的社区 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
