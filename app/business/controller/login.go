// 处理注册和登录的一个控制器
package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/ecode/status"
	"com.youyu.api/lib/log"
	"com.youyu.api/lib/path"
	"context"
	"encoding/base64"
	go_encrypt "github.com/abingzo/go-encrypt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"math/rand"
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
	jsons := &rpc.UserSign{}
	// 参数校验
	err := jsons.Validate()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.ParaMeterErr,
			"message":err.Error(),
		})
		return
	}
	err = c.BindJSON(jsons)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.JsonParseError.Code(),
			"message": ecode.JsonParseError.Message(),
			"data":    nil,
		})
		return
	}
	// 连接secretKey_rpc
	secretKeyClient := TakeSecretKeyLink()
	if secretKeyClient == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(errors.New("nil ptr"))
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

	// 不为原生登录方式则执行逻辑
	switch jsons.SignType {
	case rpc.LoginAndSignType_Phone:
		// 验证手机对应的验证码是否正确
		break
	case rpc.LoginAndSignType_Email:
		// 验证邮箱对应的验证码是否正确
		break
	case rpc.LoginAndSignType_Wechat:
		// 验证微信登录状态并随机生成账号密码
		info, err := secretKeyClient.ForWechatTokenGetInfo(context.Background(),&rpc.User{Token: jsons.VAuthToken})
		st,_ := status.FromError(err)
		if st.Code != 0 {
			c.JSON(http.StatusOK,gin.H{
				"code":st.Code,
				"message":st.Message,
			})
			return
		}
		node, err := snowflake.NewNode(int64(path.BusinessNodeNumber))
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"code":ecode.ServerErr.Code(),
				"message":ecode.ServerErr.Message(),
			})
			return
		}
		// 随机生成账号密码
		jsons.UserName = "WE" + node.Generate().Base32()
		rand.Seed(time.Now().UnixNano())
		jsons.UserPassword = strconv.FormatInt(rand.Int63() >> 2,10) + info.Openid
		jsons.WechatData.Openid = info.Openid
		break
	case rpc.LoginAndSignType_Native:
		break
	}

	// 连接data_rpc
	dataClient := TakeDataBaseLink()
	if dataClient == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(errors.New("nil ptr"))
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

// NOTE: 接口更新：增加用户验证方式
// 删除用户
// TODO:预计完善删除用户验证
func (l *SignAndLogin) DeleteSign(c *gin.Context) {
	// 绑定参数
	jsons := rpc.UserSign{}
	err := c.BindJSON(&jsons)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.JsonParseError.Code(),
			"message":ecode.JsonParseError.Message(),
		})
		return
	}
	// 验证参数
	err = jsons.Validate()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.ParaMeterErr.Code(),
			"message":err.Error(),
		})
		return
	}
	// TODO:第三方滑动验证和验证码
	// 连接secretKey_rpc
	client := TakeSecretKeyLink()
	if client == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(errors.New("nil ptr"))
		return
	}
	// 连接data_rpc
	dataClient := TakeDataBaseLink()
	if dataClient == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(errors.New("nil ptr"))
		return
	}
	// 获取Token
	tokenHead := c.Request.Header.Get("Authorization")
	tokenHeadInfo := strings.SplitN(tokenHead, " ", 2)
	// 获取token对应的uid
	uidString := c.Writer.Header().Get(tokenHeadInfo[1])
	// 根据不同的选项验证用户
	switch jsons.SignType {
	case rpc.LoginAndSignType_Native:
		uis, err := dataClient.GetUserInfoInSelf(context.Background(),&rpc.UserAuth{Uid: uidString})
		st,_ := status.FromError(err)
		if st.Code != 0 {
			c.JSON(http.StatusOK,gin.H{
				"code":st.Code,
				"message":st.Message,
			})
			return
		}
		// 获取私钥解码密码
		clientId := c.GetHeader("Client-Id")
		rsaKey, err := client.GetPrivateKey(context.Background(), &rpc.RsaKey{ClientId: clientId})
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
		// 验证用户是否存在
		jsons.UserName = uis.UserName
		jsons.UserPassword = string(rsa.GetPlainText())
		_, err = dataClient.CheckUserStatus(context.Background(),&rpc.UserLogin{
			UserName:     jsons.UserName,
			UserPassword: jsons.UserPassword,
			LoginType:    jsons.SignType,
			WechatData:   nil,
		})
		st,_ = status.FromError(err)
		if st.Code != 0 {
			c.JSON(http.StatusOK,gin.H{
				"code":st.Code,
				"message":st.Message,
			})
			return
		}
		break
	case rpc.LoginAndSignType_Phone,rpc.LoginAndSignType_Email:
		// 验证邮箱和手机号码的绑定状态
		uis, err := dataClient.GetUserInfoInSelf(context.Background(),&rpc.UserAuth{Uid: uidString})
		st,_ := status.FromError(err)
		if st.Code != 0 {
			c.JSON(http.StatusOK,gin.H{
				"code":st.Code,
				"message":st.Message,
			})
			return
		}
		if jsons.SignType == rpc.LoginAndSignType_Phone && uis.PhoneStatus == 0 {
			c.JSON(http.StatusOK,gin.H{
				"code":ecode.UserPhoneLoginNotExists.Code(),
				"message":ecode.UserPhoneLoginNotExists.Message(),
			})
			return
		} else if jsons.SignType == rpc.LoginAndSignType_Email && uis.EmailStatus == 0{
			c.JSON(http.StatusOK,gin.H{
				"code":ecode.UserEmailLoginNotExists.Code(),
				"message":ecode.UserEmailLoginNotExists.Message(),
			})
			return
		}
		// 获取验证码并对比验证码
		code, err := client.GetUserVcCode(context.Background(),&rpc.UserVcCode{BindInfo: jsons.UserBindInfo})
		st,_ = status.FromError(err)
		if st.Code != 0 {
			c.JSON(http.StatusOK,gin.H{
				"code":st.Code,
				"message":st.Message,
			})
			return
		}
		if jsons.VCode != code.VcCode {
			c.JSON(http.StatusOK,gin.H{
				"code":ecode.VcCodeError.Code(),
				"message":ecode.VcCodeError.Message(),
			})
			return
		}
		break
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
	jsons := &rpc.UserLogin{}
	err := c.BindJSON(jsons)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.JsonParseError.Code(),
			"message": ecode.JsonParseError.Message(),
			"data":    nil,
		})
		return
	}
	// TODO:校验验证码
	// 连接secretKey_rpc
	secretKeyClient := TakeSecretKeyLink()
	if secretKeyClient == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(errors.New("nil ptr"))
		return
	}
	// 连接data_rpc
	dataClient := TakeDataBaseLink()
	if dataClient == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(errors.New("nil ptr"))
		return
	}
	// 响应数据
	var baseData *rpc.BaseData
	// 根据登录方式响应
	switch jsons.LoginType {
	case rpc.LoginAndSignType_Native:
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
		grpcErr := err
		baseData, grpcErr = dataClient.CheckUserStatus(context.Background(), jsons)
		if st,_ := status.FromError(grpcErr); st.Code != int32(ecode.OK) {
			c.JSON(http.StatusOK,gin.H{
				"code": st.Code,
				"message": st.Message,
			})
			return
		}
		break
	case rpc.LoginAndSignType_Phone,rpc.LoginAndSignType_Email:
		// 校验验证码
		code, err := secretKeyClient.GetUserVcCode(context.Background(),&rpc.UserVcCode{BindInfo: jsons.UserBindInfo})
		if st,_ := status.FromError(err); st.Code != int32(ecode.OK) {
			c.JSON(http.StatusOK,gin.H{
				"code": st.Code,
				"message": st.Message,
			})
			return
		}
		if code.VcCode != jsons.VCode {
			c.JSON(http.StatusOK,gin.H{
				"code":ecode.VcCodeError.Code(),
				"message":ecode.VcCodeError.Message(),
			})
			return
		}
		// 数据清零
		jsons.UserName = ""
		jsons.UserPassword = ""
		grpcErr := error(nil)
		baseData, grpcErr = dataClient.CheckUserStatus(context.Background(),jsons)
		if st,_ := status.FromError(grpcErr); st.Code != int32(ecode.OK) {
			c.JSON(http.StatusOK,gin.H{
				"code": st.Code,
				"message": st.Message,
			})
			return
		}
		break
	case rpc.LoginAndSignType_Wechat:
		// 通过v_auth_token获取系统保存的用户第三方系统的登录信息
		info, err := secretKeyClient.ForWechatTokenGetInfo(context.Background(),&rpc.User{Token: jsons.VAuthToken})
		if st,_ := status.FromError(err); st.Code != int32(ecode.OK) {
			c.JSON(http.StatusOK,gin.H{
				"code": st.Code,
				"message": st.Message,
			})
			return
		}
		jsons.WechatData.Openid = info.Openid
		// 数据清零
		jsons.UserName = ""
		jsons.UserPassword = ""
		jsons.VCode = ""
		baseData, err = dataClient.CheckUserStatus(context.Background(),jsons)
		if st,_ := status.FromError(err); st.Code != int32(ecode.OK) {
			c.JSON(http.StatusOK,gin.H{
				"code": st.Code,
				"message": st.Message,
			})
			return
		}
		break
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
	uidString := c.Writer.Header().Get(tokenHeadInfo[1])
	// 连接secretKey_rpc
	client := TakeSecretKeyLink()
	if client == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		l.Logger.Error(errors.New("nil ptr"))
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
