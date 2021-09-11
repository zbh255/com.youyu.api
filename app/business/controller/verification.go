// 验证相关控制器
package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/ecode/status"
	"com.youyu.api/lib/log"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type VerificationApi interface {
	WeChatLogin(c *gin.Context)
	WeiboLogin(c *gin.Context)
	OtherLogin(c *gin.Context)
	SendVerificationCode(c *gin.Context)
}

type Verification struct {
	Logger log.Logger
}

// OtherLogin 分配第三方登录的类型
func (v *Verification) OtherLogin(c *gin.Context) {
	switch c.Query("type") {
	case "wechat_login":
		v.WeChatLogin(c)
		break
	case "weibo_login":
		v.WeiboLogin(c)
		break
	}
}

func (v *Verification) WeChatLogin(c *gin.Context) {
	wxCode := c.Query("code")
	wxUrl := fmt.Sprintf(WechatLoginRawUrl, ConnectAndConf.Config.Project.Auth.WechatLogin.AppID,
		ConnectAndConf.Config.Project.Auth.WechatLogin.AppSercret, wxCode)
	response, err := http.Get(wxUrl)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		return
	}
	bytes := make([]byte,response.ContentLength)
	_, err = response.Body.Read(bytes)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		return
	}
	jsons := rpc.WechatTokenInfo{}
	err = json.Unmarshal(bytes, &jsons)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.JsonParseError.Code(),
			"message": ecode.JsonParseError.Message(),
		})
		return
	}
	// 判断微信服务器的返回值是否正确
	if jsons.Errcode != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    jsons.Errcode,
			"message": jsons.Errmsg,
		})
		return
	}
	// 成功则签钥
	// 连接secretKey_rpc
	secretKeyLis, err := ConnectAndConf.SecretKeyRpcConnPool.Get()
	defer ConnectAndConf.SecretKeyRpcConnPool.Put(secretKeyLis)
	secretKeyClient, _, err := GetSecretKeyRpcServer(secretKeyLis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		v.Logger.Error(err)
		return
	}
	wechatToken, err := secretKeyClient.BindWechatToken(context.Background(), &jsons)
	st, _ := status.FromError(err)
	c.JSON(http.StatusOK, gin.H{
		"code":         st.Code,
		"message":      st.Message,
		"wechat_token": wechatToken.Token,
	})
}

func (v *Verification) WeiboLogin(c *gin.Context) {
	panic("implement me")
}

func (v *Verification) SendVerificationCode(c *gin.Context) {
	panic("implement me")
}