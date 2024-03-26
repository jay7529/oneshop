package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK,
		gin.H{
			"code": 200,
			"data": data,
			"msg":  msg,
		})
}

func Failed(c *gin.Context, msg string) {
	c.JSON(http.StatusOK,
		gin.H{
			"code": 100,
			"msg":  msg,
		})
}

func Error(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest,
		gin.H{
			"code": 400,
			"err":  err,
		})
}

func CheckErr(err error) bool {
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	return true
}
