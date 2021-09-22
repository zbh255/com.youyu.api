// 抽象出来的公共函数
package controller

import (
	"com.youyu.api/lib/ecode"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 返回服务器错误的响应
func ReturnServerErrJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    ecode.ServerErr.Code(),
		"message": ecode.ServerErr.Message(),
	})
}

// 返回Json序列化失败的错误
func ReturnJsonParseErrJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    ecode.JsonParseError.Code(),
		"message": ecode.JsonParseError.Message(),
	})
}

// 返回参数解析失败的错误
func ReturnParaMeterErrJsons(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"code":ecode.ParaMeterErr.Code(),
		"message":ecode.ParaMeterErr.Message(),
	})
}

func GetHeaderTokenBindTheUid(c *gin.Context) string  {
	// 获取Token
	tokenHead := c.Request.Header.Get("Authorization")
	tokenHeadInfo := strings.SplitN(tokenHead, " ", 2)
	// 获取token对应的uid
	return c.Writer.Header().Get(tokenHeadInfo[1])
}
