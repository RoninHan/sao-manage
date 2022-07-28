package Response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK, // 自定义code
		"msg":  msg,  // message
		"data": data, // 数据
	})
	return
}

// 返回失败
func Err(c *gin.Context, httpCode int, msg string, jsonStr interface{}) {
	c.JSON(httpCode, map[string]interface{}{
		"code": httpCode,
		"msg":  msg,
		"data": jsonStr,
	})
	return
}