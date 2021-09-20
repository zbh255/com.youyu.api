// 用户接口的控制器
package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/ecode/status"
	"com.youyu.api/lib/log"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type UserApi interface {
	UserOption(c *gin.Context)
	GetUserInfo(c *gin.Context)
	UpdateUserInfo(c *gin.Context)
	AddUserCheckInfo(c *gin.Context)
	AddUserHeadPortrait(c *gin.Context)
}

type UserInfo struct {
	Logger log.Logger
}

func (u UserInfo) UserOption(c *gin.Context)  {

}

func (u *UserInfo) GetUserInfo(c *gin.Context) {
	self := false
	uid := ""
	if _,bl := c.Get("Result"); bl {
		uid = c.Param("uid")
	} else {
		// 获取Token
		tokenHead := c.Request.Header.Get("Authorization")
		tokenHeadInfo := strings.SplitN(tokenHead, " ", 2)
		uid = c.DefaultQuery("uid","nouid")
		if uid == c.Writer.Header().Get(tokenHeadInfo[1]) {
			self = true
		}
	}
	// 连接data_rpc
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	dataClient, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		u.Logger.Error(err)
		return
	}
	if self {
		uis,err := dataClient.GetUserInfoInSelf(context.Background(),&rpc.UserAuth{Uid: uid})
		st,_ := status.FromError(err)
		c.JSON(http.StatusOK,gin.H{
			"code":st.Code,
			"message":st.Message,
			"data":uis,
		})
	} else {
		uios,err := dataClient.GetUserInfoInOther(context.Background(),&rpc.UserAuth{Uid: uid})
		st,_ := status.FromError(err)
		c.JSON(http.StatusOK,gin.H{
			"code":st.Code,
			"message":st.Message,
			"data":uios,
		})
	}
}

func (u *UserInfo) UpdateUserInfo(c *gin.Context) {
	jsons := rpc.UserInfoSet{}
	err := c.BindJSON(&jsons)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.JsonParseError.Code(),
			"message":ecode.JsonParseError.Message(),
		})
		return
	}
	// 参数校验
	err = jsons.Validate()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.ParaMeterErr.Code(),
			"message":err.Error(),
		})
		return
	}
	// 获取Token
	tokenHead := c.Request.Header.Get("Authorization")
	tokenHeadInfo := strings.SplitN(tokenHead, " ", 2)
	// 获取token对应的uid
	uidString := c.Writer.Header().Get(tokenHeadInfo[1])
	uid,err := strconv.Atoi(uidString)
	if err != nil {
		ReturnServerErrJson(c)
		return
	}
	jsons.Uid = int32(uid)
	// 连接data_rpc
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	dataClient, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		u.Logger.Error(err)
		return
	}
	_, err = dataClient.UpdateUserInfo(context.Background(),&jsons)
	st,_ := status.FromError(err)
	c.JSON(http.StatusOK,gin.H{
		"code":st.Code,
		"message":st.Message,
	})
}

func (u *UserInfo) AddUserCheckInfo(c *gin.Context) {
	jsons := rpc.UserSign{}
	err := c.BindJSON(&jsons)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.JsonParseError.Code(),
			"message":ecode.JsonParseError.Message(),
		})
		return
	}
	// 参数校验
	err = jsons.Validate()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.ParaMeterErr.Code(),
			"message":ecode.ParaMeterErr.Message(),
		})
		return
	}
	// 根据不同的注册类型来添加用户验证信息
	// 连接secretKey_rpc
	secretKeyLis, err := ConnectAndConf.SecretKeyRpcConnPool.Get()
	defer ConnectAndConf.SecretKeyRpcConnPool.Put(secretKeyLis)
	client, _, err := GetSecretKeyRpcServer(secretKeyLis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		u.Logger.Error(err)
		return
	}
	// 连接data_rpc
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	dataClient, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		u.Logger.Error(err)
		return
	}
	// 获取Token
	tokenHead := c.Request.Header.Get("Authorization")
	tokenHeadInfo := strings.SplitN(tokenHead, " ", 2)
	// 获取token对应的uid
	uidString := c.Writer.Header().Get(tokenHeadInfo[1])
	// 根据不同的选项验证用户
	switch jsons.SignType {
	case rpc.LoginAndSignType_Phone,rpc.LoginAndSignType_Email:
		// 获取验证码并对比验证码
		code, err := client.GetUserVcCode(context.Background(),&rpc.UserVcCode{BindInfo: jsons.UserBindInfo})
		st,_ := status.FromError(err)
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
		// 添加绑定
		if jsons.SignType == rpc.LoginAndSignType_Phone {
			phone, err := strconv.ParseInt(jsons.UserBindInfo,10,64)
			if err != nil {
				c.JSON(http.StatusOK,gin.H{
					"code":ecode.ServerErr.Code(),
					"message":ecode.ServerErr.Message(),
				})
				return
			}
			_, err = dataClient.AddUserCheckInfoPhone(context.Background(),&rpc.UserCheckPhone{
				Phone: phone,
				Ua:    &rpc.UserAuth{Uid: uidString},
			})
			st,_ := status.FromError(err)
			c.JSON(http.StatusOK,gin.H{
				"code":st.Code,
				"message":st.Message,
			})
		} else if jsons.SignType == rpc.LoginAndSignType_Email {
			_, err := dataClient.AddUserCheckInfoEmail(context.Background(),&rpc.UserCheckEmail{
				Email: jsons.UserBindInfo,
				Ua:    &rpc.UserAuth{Uid: uidString},
			})
			st,_ := status.FromError(err)
			c.JSON(http.StatusOK,gin.H{
				"code":st.Code,
				"message":st.Message,
			})
		}
		break
	case rpc.LoginAndSignType_Wechat:
		// 校验v_auth_token状态
		info, err := client.ForWechatTokenGetInfo(context.Background(),&rpc.User{Token: jsons.VAuthToken})
		st,_ := status.FromError(err)
		if st.Code != 0 {
			c.JSON(http.StatusOK,gin.H{
				"code":st.Code,
				"message":st.Message,
			})
			return
		}
		_, err = dataClient.AddUserCheckInfoWechat(context.Background(),&rpc.UserCheckWechat{
			Openid: info.Openid,
			Ua:     &rpc.UserAuth{Uid: uidString},
		})
		st,_ = status.FromError(err)
		c.JSON(http.StatusOK,gin.H{
			"code":st.Code,
			"message":st.Message,
		})
	}
}

// NOTE: 废弃方法
func (u UserInfo) AddUserHeadPortrait(c *gin.Context)   {

}


