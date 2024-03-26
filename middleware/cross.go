package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {
	method := c.Request.Method

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 支持的所有跨域请求的方法
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	// 允許跨域設置可以返回其他字段，可以自定義字段
	c.Header("Access-Control-Allow-Headers", "*")
	// 允許瀏覽器可以解析的Header
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	// 允許客户端傳遞效驗訊息
	c.Header("Access-Control-Allow-Credentials", "true")

	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 處理請求
	c.Next()
}
