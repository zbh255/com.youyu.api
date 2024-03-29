package middleware

import (
	"com.youyu.api/app/business/controller"
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/auth"
	"com.youyu.api/lib/ecode"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
	"strings"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Token
		tokenHead := c.Request.Header.Get("Authorization")
		tokenHeadInfo := strings.SplitN(tokenHead, " ", 2)
		// 如果Authorization Header为空则填充默认数据走完流程
		switch len(tokenHeadInfo) {
		case 0:
			tokenHeadInfo = []string{"Bearer", "len 0 token"}
			break
		case 1:
			tokenHeadInfo = []string{"Bearer", "len 1 token"}
			break
		}
		// 获取签名密钥
		jwtC := auth.New(controller.TokenSigningKey)
		myClaims, err := jwtC.ParseToken(tokenHeadInfo[1])
		// 构造正确的值使流程无需改变
		if err == nil {err = &jwt.ValidationError{Errors: 0}}
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			// token 过时
			c.JSON(http.StatusOK, gin.H{
				"code":    ecode.AccessTokenExpires.Code(),
				"message": ecode.AccessTokenExpires.Message(),
			})
			c.Abort()
			return
		case jwt.ValidationErrorSignatureInvalid:
			// token 签名验证失败
			c.JSON(http.StatusOK, gin.H{
				"code":    ecode.AccessTokenSignature.Code(),
				"message": ecode.AccessTokenSignature.Message(),
			})
			c.Abort()
			return
		case 0:
			// 无错误
			break
		default:
			c.JSON(http.StatusOK, gin.H{
				"code":    ecode.AccessTokenErr.Code(),
				"message": ecode.AccessTokenErr.Message(),
			})
			c.Abort()
			return
		}
		// 连接签钥服务器
		client := controller.TakeSecretKeyLink()
		if client == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    ecode.ServerErr.Code(),
				"message": ecode.ServerErr.Message(),
			})
			c.Abort()
			return
		}
		// 检查用户是否在登录状态
		_, err = client.ForTokenGetBindUser(context.Background(), &rpc.User{Token: tokenHeadInfo[1], ExpTime: myClaims.ExpiresAt})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    ecode.NoLogin.Code(),
				"message": ecode.NoLogin.Message(),
			})
			c.Abort()
			return
		}
		// 成功则下一步
		c.Header(tokenHeadInfo[1], strconv.FormatInt(myClaims.Uid, 10))
		c.Next()
	}
}
