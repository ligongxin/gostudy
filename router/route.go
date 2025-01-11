package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-app/controller"
	"web-app/logger"
	"web-app/middlewares"
	"web-app/pkg/snowflake"
	"web-app/task"
)

func SetupRoute(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置为生产模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})
	// 手动触发刷新和奖励结算任务
	r.GET("/task", func(c *gin.Context) {
		task.TriggerManualRefreshAndSettle()
		c.JSON(http.StatusOK, gin.H{"message": "任务已手动触发"})
	})

	r.GET("/token", controller.GetTokenHandler)

	r.POST("/signup", controller.SignupHandler)
	r.POST("/login", controller.LoginHandler)
	r.GET("/snow", middlewares.JwtAuthDiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": snowflake.GenID()})
	})

	return r
}
