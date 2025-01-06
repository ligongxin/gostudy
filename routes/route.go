package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-app/logger"
)

func SetupRoute() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})
	return r
}
