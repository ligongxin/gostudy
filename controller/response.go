package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(c *gin.Context, code int) {
	c.JSON(http.StatusOK, &ResponseData{
		Code:    code,
		Message: nil,
		Data:    nil,
	})
}
