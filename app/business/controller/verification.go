// 验证相关控制器
package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/ecode/status"
	"com.youyu.api/lib/log"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type VerificationApi interface {
	OtherLogin(c *gin.Context)
	SendVerificationCode(c *gin.Context)
}

type Verification struct {
	Logger log.Logger
}

type OtherLoginBase struct {
	Protocol string `json:"protocol"`
	Type string `json:"type"`
	Code int `json:"code"`
}

type SendVcCodeBase struct {
	AddrType string `json:"addr_type"`
	Addr string `json:"addr"`
}

// OtherLogin 分配第三方登录的类型
func (v *Verification) OtherLogin(c *gin.Context) {
	jsons := OtherLoginBase{}
	if c.BindJSON(&jsons) != nil {
		ReturnJsonParseErrJson(c)
		return
	}
	switch jsons.Type {
	case "wechat_login":
		v.WeChatLogin(c,&jsons)
		break
	case "weibo_login":
		v.WeiboLogin(c,&jsons)
		break
	}
}

func (v *Verification) WeChatLogin(c *gin.Context,base *OtherLoginBase) {
	wxCode := strconv.Itoa(base.Code)
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

func (v *Verification) WeiboLogin(c *gin.Context,base *OtherLoginBase) {

}

func (v *Verification) SendVerificationCode(c *gin.Context) {
	jsons := SendVcCodeBase{}
	if c.BindJSON(&jsons) != nil {
		ReturnJsonParseErrJson(c)
		return
	}
	addr, err := base64.StdEncoding.DecodeString(jsons.Addr)
	rand.Seed(time.Now().UnixNano() - 2 << 8)
	vcCode := rand.Intn(875555) + 100555
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.EncodeError.Code(),
			"message":ecode.EncodeError.Message(),
		})
		return
	}
	switch jsons.AddrType {
	case "email":
		// TODO:邮箱发送验证码的逻辑
		// 存储验证码
		lis, client, err := LinkSecretKeyRpc()
		if err != nil {
			ReturnServerErrJson(c)
			return
		}
		defer ConnectAndConf.SecretKeyRpcConnPool.Put(lis)
		_, err = client.BindUserVcCode(context.Background(),&rpc.UserVcCode{BindInfo: string(addr),VcCode: strconv.Itoa(vcCode)})
		st,_ := status.FromError(err)
		c.JSON(http.StatusOK,gin.H{
			"code":st.Code,
			"message":st.Message,
		})
	case "phone":
		// TODO:发送手机验证码的逻辑
		// 存储验证码
		lis, client, err := LinkSecretKeyRpc()
		if err != nil {
			ReturnServerErrJson(c)
			return
		}
		defer ConnectAndConf.SecretKeyRpcConnPool.Put(lis)
		_, err = client.BindUserVcCode(context.Background(),&rpc.UserVcCode{BindInfo: string(addr),VcCode: strconv.Itoa(vcCode)})
		st,_ := status.FromError(err)
		c.JSON(http.StatusOK,gin.H{
			"code":st.Code,
			"message":st.Message,
		})
	default:
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.ParaMeterErr.Code(),
			"message":ecode.ParaMeterErr.Message(),
		})
	}
}
