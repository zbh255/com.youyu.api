package middleware

import (
	"com.youyu.api/lib/ecode"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"reflect"
	"unsafe"
)

// ClientIdAuth 鉴别请求头中的ClientId
// 没有ClientId 则返回一个新的ClientId
// 兼容性调整
func ClientIdAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientId := c.Request.Header.Get("Client-Id")
		_, err := uuid.FromString(clientId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    ecode.ClientIdError.Code(),
				"message": ecode.ClientIdError.Message(),
				"uuid":    uuid.NewV4().String(),
			})
			c.Abort()
			return
		}

		url := c.Request.URL.Query()
		url.Set("client_id",clientId)
		// 修改gin.Context中的私有变量
		v := reflect.ValueOf(c).Elem().FieldByName("queryCache")
		vv := reflect.NewAt(v.Type(),unsafe.Pointer(v.UnsafeAddr())).Elem()
		rv := reflect.ValueOf(url)
		if vv.Kind() != vv.Kind() {
			panic(fmt.Errorf("invalid kind, expected kind: %v, got kind:%v", v.Kind(), rv.Kind()))
		}
		vv.Set(rv)

		c.Request.URL.RawQuery += fmt.Sprintf("&%s=%s","client_id",clientId)
		c.Next()
	}
}
