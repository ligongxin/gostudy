package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
	"web-app/controller"
	_ "web-app/docs"
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

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
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

	v1 := r.Group("/api/v1")
	// 注册
	v1.POST("/signup", controller.SignupHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)

	r.GET("/snow", middlewares.JwtAuthDiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": snowflake.GenID()})
	})
	v1.Use(middlewares.JwtAuthDiddleWare())
	{
		v1.GET("/community", controller.CommunityHandler)                 // 社区
		v1.GET("/community/:id", controller.CommunityDetailHandler)       //社区详情
		v1.POST("/post", controller.CreatePostHandler)                    // 创建帖子
		v1.GET("/post/:id", controller.PostDetailHandler)                 // 帖子详情
		v1.GET("/post", controller.GetPostListHandler)                    // 帖子列表
		v1.GET("/post/v2", controller.GetPostListHandler2)                // 帖子列表
		v1.GET("/community/post", controller.GetCommunityPostListHandler) // 查询社区下的帖子
		v1.POST("/vote", controller.PostVoteController)                   // 投票
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": 404})
	})
	return r
}
