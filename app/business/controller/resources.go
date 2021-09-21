// 负责管理授予客户端访问资源的临时token的控制器
package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/ecode/status"
	"com.youyu.api/lib/log"
	"com.youyu.api/lib/utils"
	"context"
	"fmt"
	go_encrypt "github.com/abingzo/go-encrypt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 返回的json信息
type SignatureTmpInfo struct {
	TmpSecretID   string `json:"tmp_secret_id"`
	TmpSecretKey  string `json:"tmp_secret_key"`
	TmpSessionKey string `json:"tmp_session_key"`
	FileName      string `json:"file_name"`
	BucketPath    string `json:"bucket_path"`
	CosPath       string `json:"cos_path"`
}

type ReSourcesApi interface {
	GetUploadHeadPortraitToken(c *gin.Context)
	GetUploadArticleVideoToken(c *gin.Context)
	GetUploadArticleImageToken(c *gin.Context)
}

type ReSources struct {
	Logger log.Logger
}

func (r *ReSources) GetUploadHeadPortraitToken(c *gin.Context) {
	// 获取Token
	tokenHead := c.Request.Header.Get("Authorization")
	tokenHeadInfo := strings.SplitN(tokenHead, " ", 2)
	// 获取token对应的uid
	uidString := c.Writer.Header().Get(tokenHeadInfo[1])
	// 设置头像能访问的权限并返回给前端临时密钥
	conf := ConnectAndConf.Config.Project.Cos
	urlPre := fmt.Sprintf("https://%s-%d.cos.%s.myqcloud.com", conf.PublicSourceBucket.Name, conf.Appid, conf.PublicSourceBucket.Region)
	fileType := c.DefaultQuery("file_type", ".png")
	if !utils.EqualForSlice(fileType,ConnectAndConf.Config.Project.UploadImageType) {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.FileTypeErr.Code(),
			"message":ecode.FileTypeErr.Message(),
		})
		return
	}
	hash, err := go_encrypt.NewCoder().GetAbstract().Md5Coder(go_encrypt.BASE64).SumString(uidString).
		Append(time.Now().String()).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		r.Logger.Error(err)
		return
	}
	// 连接data_rpc
	dataClient := TakeDataBaseLink()
	if dataClient == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		r.Logger.Error(errors.New("nil ptr"))
		return
	}
	uid, err := strconv.Atoi(uidString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		return
	}
	fileUrl := fmt.Sprintf("%s/%s/%s", urlPre, ConnectAndConf.Config.Project.CosHeadPortraitDir, string(hash)+fileType)
	_, err = dataClient.AddUserHeadPortrait(context.Background(), &rpc.UserHeadPortraitSet{Uid: int32(uid), Url: fileUrl})
	st, _ := status.FromError(err)
	if st.Code != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
		})
		return
	}
	// 给客户端发放签名
	// 新建客户端
	client := sts.NewClient(conf.SecretID, conf.SecretKey, nil)
	opt := &sts.CredentialOptions{
		Policy: &sts.CredentialPolicy{
			Version: "2.0",
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
					},
					Effect: "allow",
					Resource: []string{
						"qcs::cos:" + conf.PublicSourceBucket.Region + ":uid/" +
							strconv.Itoa(conf.Appid) + ":" + conf.PublicSourceBucket.Name +
							fmt.Sprintf("/%s/%s", ConnectAndConf.Config.Project.CosHeadPortraitDir, string(hash)+fileType),
					},
				},
			},
		},
		Region:          conf.PublicSourceBucket.Region,
		DurationSeconds: int64(conf.DurationSeconds),
	}
	credential, err := client.GetCredential(opt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": err.Error(),
		})
		return
	}
	jsons := SignatureTmpInfo{
		TmpSecretID:   credential.Credentials.TmpSecretID,
		TmpSecretKey:  credential.Credentials.TmpSecretKey,
		TmpSessionKey: credential.Credentials.SessionToken,
		FileName:      string(hash) + fileType,
		BucketPath:    "/" + ConnectAndConf.Config.Project.CosHeadPortraitDir,
		CosPath:       urlPre,
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    ecode.OK.Code(),
		"message": ecode.OK.Message(),
		"data":    &jsons,
	})
}

func (r *ReSources) GetUploadArticleVideoToken(c *gin.Context) {
	// 获取Token
	tokenHead := c.Request.Header.Get("Authorization")
	tokenHeadInfo := strings.SplitN(tokenHead, " ", 2)
	// 获取token对应的uid
	uidString := c.Writer.Header().Get(tokenHeadInfo[1])
	// 设置视频能访问的权限并返回给前端临时密钥
	conf := ConnectAndConf.Config.Project.Cos
	urlPre := fmt.Sprintf("https://%s-%d.cos.%s.myqcloud.com", conf.PublicSourceBucket.Name, conf.Appid, conf.PublicSourceBucket.Region)
	fileName := c.Param("name")
	fileType := "." + strings.SplitN(fileName,".",2)[1]
	if !utils.EqualForSlice(fileType,ConnectAndConf.Config.Project.UploadVideoType) {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.FileTypeErr.Code(),
			"message":ecode.FileTypeErr.Message(),
		})
		return
	}
	hash, err := go_encrypt.NewCoder().GetAbstract().Md5Coder(go_encrypt.BASE64).SumString(uidString).
		Append(time.Now().String()).Append(fileName).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		r.Logger.Error(err)
		return
	}
	//fileUrl := fmt.Sprintf("%s/%s/%s", urlPre, ConnectAndConf.Config.Project.CosVideoDir, string(hash)+fileType)
	// 给客户端发放签名
	// 新建客户端
	client := sts.NewClient(conf.SecretID, conf.SecretKey, nil)
	opt := &sts.CredentialOptions{
		Policy: &sts.CredentialPolicy{
			Version: "2.0",
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
					},
					Effect: "allow",
					Resource: []string{
						"qcs::cos:" + conf.PublicSourceBucket.Region + ":uid/" +
							strconv.Itoa(conf.Appid) + ":" + conf.PublicSourceBucket.Name +
							fmt.Sprintf("/%s/%s", ConnectAndConf.Config.Project.CosVideoDir, string(hash)+fileType),
					},
				},
			},
		},
		Region:          conf.PublicSourceBucket.Region,
		DurationSeconds: int64(conf.DurationSeconds),
	}
	credential, err := client.GetCredential(opt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": err.Error(),
		})
		return
	}
	jsons := SignatureTmpInfo{
		TmpSecretID:   credential.Credentials.TmpSecretID,
		TmpSecretKey:  credential.Credentials.TmpSecretKey,
		TmpSessionKey: credential.Credentials.SessionToken,
		FileName:      string(hash) + fileType,
		BucketPath:    "/" + ConnectAndConf.Config.Project.CosHeadPortraitDir,
		CosPath:       urlPre,
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    ecode.OK.Code(),
		"message": ecode.OK.Message(),
		"data":    &jsons,
	})
}

func (r *ReSources) GetUploadArticleImageToken(c *gin.Context) {
	// 获取Token
	tokenHead := c.Request.Header.Get("Authorization")
	tokenHeadInfo := strings.SplitN(tokenHead, " ", 2)
	// 获取token对应的uid
	uidString := c.Writer.Header().Get(tokenHeadInfo[1])
	// 设置视频能访问的权限并返回给前端临时密钥
	conf := ConnectAndConf.Config.Project.Cos
	urlPre := fmt.Sprintf("https://%s-%d.cos.%s.myqcloud.com", conf.PublicSourceBucket.Name, conf.Appid, conf.PublicSourceBucket.Region)
	fileName := c.Param("name")
	fileType := "." + strings.SplitN(fileName,".",2)[1]
	// 校验文件类型
	if !utils.EqualForSlice(fileType,ConnectAndConf.Config.Project.UploadImageType) {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.FileTypeErr.Code(),
			"message":ecode.FileTypeErr.Message(),
		})
		return
	}
	hash, err := go_encrypt.NewCoder().GetAbstract().Md5Coder(go_encrypt.BASE64).SumString(uidString).
		Append(time.Now().String()).Append(fileName).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		r.Logger.Error(err)
		return
	}
	//fileUrl := fmt.Sprintf("%s/%s/%s", urlPre, ConnectAndConf.Config.Project.CosVideoDir, string(hash)+fileType)
	// 给客户端发放签名
	// 新建客户端
	client := sts.NewClient(conf.SecretID, conf.SecretKey, nil)
	opt := &sts.CredentialOptions{
		Policy: &sts.CredentialPolicy{
			Version: "2.0",
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
					},
					Effect: "allow",
					Resource: []string{
						"qcs::cos:" + conf.PublicSourceBucket.Region + ":uid/" +
							strconv.Itoa(conf.Appid) + ":" + conf.PublicSourceBucket.Name +
							fmt.Sprintf("/%s/%s", ConnectAndConf.Config.Project.CosImgDir, string(hash)+fileType),
					},
				},
			},
		},
		Region:          conf.PublicSourceBucket.Region,
		DurationSeconds: int64(conf.DurationSeconds),
	}
	credential, err := client.GetCredential(opt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": err.Error(),
		})
		return
	}
	jsons := SignatureTmpInfo{
		TmpSecretID:   credential.Credentials.TmpSecretID,
		TmpSecretKey:  credential.Credentials.TmpSecretKey,
		TmpSessionKey: credential.Credentials.SessionToken,
		FileName:      string(hash) + fileType,
		BucketPath:    "/" + ConnectAndConf.Config.Project.CosImgDir,
		CosPath:       urlPre,
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    ecode.OK.Code(),
		"message": ecode.OK.Message(),
		"data":    &jsons,
	})
}
