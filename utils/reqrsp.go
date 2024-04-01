package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, token string, data interface{}, msg string) {
	c.JSON(http.StatusOK,
		gin.H{
			"code":  200,
			"token": token,
			"data":  data,
			"msg":   msg,
		})
}

func Failed(c *gin.Context, token string, msg string) {
	c.JSON(http.StatusOK,
		gin.H{
			"code":  100,
			"token": token,
			"msg":   msg,
		})
}

func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest,
		gin.H{
			"code": 400,
			"msg":  msg,
		})
}
