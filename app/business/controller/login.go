// 处理注册和登录的一个控制器
package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/ecode/status"
	"com.youyu.api/lib/log"
	"context"
	"encoding/base64"
	go_encrypt "github.com/abingzo/go-encrypt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type SignAndLoginApi interface {
	CreateSign(c *gin.Context)
	DeleteSign(c *gin.Context)
	CreateLoginState(c *gin.Context)
	DeleteLoginState(c *gin.Context)
}

type SignAndLogin struct {
	Logger log.Logger
}

// 注册用户
func (l *SignAndLogin) CreateSign(c *gin.Context) {
	jsons := &rpc.UserLoginOrSign{}
	err := c.BindJSON(jsons)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.JsonParseError.Code(),
			"message": ecode.JsonParseError.Message(),
			"data":    nil,
		})
		return
	}
	// 连接secretKey_rpc
	secretKeyLis, err := ConnectAndConf.SecretKeyRpcConnPool.Get()
	defer ConnectAndConf.SecretKeyRpcConnPool.Put(secretKeyLis)
	secretKeyClient, _, err := GetSecretKeyRpcServer(secretKeyLis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(err)
		return
	}
	// 获取私钥解码密码
	clientId := c.DefaultQuery("client_id", "noClient_id")
	rsaKey, err := secretKeyClient.GetPrivateKey(context.Background(), &rpc.RsaKey{ClientId: clientId})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.SecretKeyTimeout.Code(),
			"message": ecode.SecretKeyTimeout.Message(),
		})
		return
	}
	// 密码base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(jsons.UserPassword)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": ecode.EncodeError.Code(),
			"message": ecode.EncodeError.Message(),
		})
		return
	}
	// 解密密码
	rsa := go_encrypt.NewCoder().GetEncrypted().RsaCoder(go_encrypt.BitSize1024, nil, nil).
		SetPrivateKeyPem([]byte(rsaKey.PrivateKey)).Decode(decodeBytes)
	if rsa.Err() != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.AccessKeyErr.Code(),
			"message": ecode.AccessKeyErr.Message(),
		})
		return
	}
	jsons.UserPassword = string(rsa.GetPlainText())

	// 连接data_rpc
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	dataClient, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(err)
		return
	}
	// TODO:校验验证码和第三方验证系统token
	_, grpcErr := dataClient.CreateUserSign(context.Background(), jsons)
	if st,bl := status.FromError(grpcErr); bl {
		c.JSON(http.StatusOK,gin.H{
			"code": st.Code,
			"message":st.Message,
			"data":nil,
		})
	}
}

// 删除用户
// TODO:预计增加删除用户验证
func (l *SignAndLogin) DeleteSign(c *gin.Context) {
	// 获取Token
	tokenHead := c.Request.Header.Get("Authorization")
	tokenHeadInfo := strings.SplitN(tokenHead, " ", 2)
	// 获取token对应的uid
	uidString := c.GetHeader(tokenHeadInfo[1])
	// 连接data_rpc
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	dataClient, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(err)
		return
	}
	_, grpcErr := dataClient.DeleteUserSign(context.Background(), &rpc.UserAuth{Uid: uidString})
	if st,_ := status.FromError(grpcErr); st.Code != int32(ecode.OK) {
		c.JSON(http.StatusOK,gin.H{
			"code": st.Code,
			"message": st.Message,
		})
		return
	}
	// 删除登录状态Token
	// 连接secretKey_rpc
	secretKeyLis, err := ConnectAndConf.SecretKeyRpcConnPool.Get()
	defer ConnectAndConf.SecretKeyRpcConnPool.Put(secretKeyLis)
	client, _, err := GetSecretKeyRpcServer(secretKeyLis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(err)
		return
	}
	// 获取客户端id
	uid, err := strconv.Atoi(uidString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(err)
		return
	}
	_, err = client.DeleteBindUser(context.Background(), &rpc.User{
		Uid:   int32(uid),
		Token: tokenHeadInfo[1],
	})
	c.JSON(http.StatusOK, gin.H{
		"code":    ecode.OK.Code(),
		"message": ecode.OK.Message(),
	})
}

// 登录
func (l *SignAndLogin) CreateLoginState(c *gin.Context) {
	jsons := &rpc.UserLoginOrSign{}
	err := c.BindJSON(jsons)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.JsonParseError.Code(),
			"message": ecode.JsonParseError.Message(),
			"data":    nil,
		})
		return
	}
	// 连接secretKey_rpc
	secretKeyLis, err := ConnectAndConf.SecretKeyRpcConnPool.Get()
	defer ConnectAndConf.SecretKeyRpcConnPool.Put(secretKeyLis)
	secretKeyClient, _, err := GetSecretKeyRpcServer(secretKeyLis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(err)
		return
	}
	// 获取私钥解码密码
	clientId := c.DefaultQuery("client_id", "noClient_id")
	rsaKey, err := secretKeyClient.GetPrivateKey(context.Background(), &rpc.RsaKey{ClientId: clientId})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.SecretKeyTimeout.Code(),
			"message": ecode.SecretKeyTimeout.Message(),
		})
		return
	}
	// 密码base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(jsons.UserPassword)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": ecode.EncodeError.Code(),
			"message": ecode.EncodeError.Message(),
		})
		return
	}
	// 解密密码
	rsa := go_encrypt.NewCoder().GetEncrypted().RsaCoder(go_encrypt.BitSize1024, nil, nil).
		SetPrivateKeyPem([]byte(rsaKey.PrivateKey)).Decode(decodeBytes)
	if rsa.Err() != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.AccessKeyErr.Code(),
			"message": ecode.AccessKeyErr.Message(),
		})
		return
	}
	jsons.UserPassword = string(rsa.GetPlainText())

	// 连接data_rpc
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	dataClient, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(err)
		return
	}
	// TODO:校验验证码和第三方验证系统token
	baseData, grpcErr := dataClient.CheckUserStatus(context.Background(), jsons)
	if st,_ := status.FromError(grpcErr); st.Code != int32(ecode.OK) {
		c.JSON(http.StatusOK,gin.H{
			"code": st.Code,
			"message": st.Message,
		})
		return
	}
	// 校验成功则给用户签钥
	uid, err := strconv.Atoi(baseData.Data["uid"])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(err)
		return
	}
	// 保存token的时间，也就是保持登录状态的时间
	// Save为0 代表不保持，默认一天,1为7天,2为14天
	exp := 24 * time.Hour
	if jsons.Save != 0 && jsons.Save <= 2 {
		exp = exp * time.Duration(int(jsons.Save)*7)
	}
	// 如果测试时间有值，则按照测试时间
	if TokenExpTime > 0 {
		exp = time.Duration(TokenExpTime) * time.Minute
	}
	user, err := secretKeyClient.BindTokenToUser(context.Background(), &rpc.User{
		Uid:     int32(uid),
		ExpTime: int64(exp),
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(err)
		return
	}
	// 记录登录的信息
	l.Logger.Info("sign token :" + user.Token)
	l.Logger.Info("user sign :" + strconv.Itoa(int(user.Uid)))
	c.JSON(http.StatusOK, gin.H{
		"code":    ecode.OK.Code(),
		"message": ecode.OK.Message(),
		"token":   user.Token,
	})
}

// 注销
func (l *SignAndLogin) DeleteLoginState(c *gin.Context) {
	// 获取Token
	tokenHead := c.Request.Header.Get("Authorization")
	tokenHeadInfo := strings.SplitN(tokenHead, " ", 2)
	// 获取token对应的uid
	uidString := c.GetHeader(tokenHeadInfo[1])
	// 连接secretKey_rpc
	secretKeyLis, err := ConnectAndConf.SecretKeyRpcConnPool.Get()
	defer ConnectAndConf.SecretKeyRpcConnPool.Put(secretKeyLis)
	client, _, err := GetSecretKeyRpcServer(secretKeyLis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(err)
		return
	}
	// 获取客户端id
	uid, err := strconv.Atoi(uidString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(err)
		return
	}
	_, err = client.DeleteBindUser(context.Background(), &rpc.User{
		Uid:   int32(uid),
		Token: tokenHeadInfo[1],
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(err)
		return
	}
	// 删除成功
	c.JSON(http.StatusOK, gin.H{
		"code":    ecode.OK.Code(),
		"message": ecode.OK.Message(),
	})
}
