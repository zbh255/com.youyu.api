package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/ecode/status"
	"com.youyu.api/lib/log"
	"com.youyu.api/lib/utils"
	"context"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TagsApi interface {
	AddTag(c *gin.Context)
	GetTagInt32Id(c *gin.Context)
	TagOpt(c *gin.Context)
	GetTagText(c *gin.Context)
}

type Tags struct {
	Logger log.Logger
}

type AddTagBase struct {
	Tags []string `json:"tags"`
}

// 遵守REST规范做出了资源通过Body传递的改动
func (t *Tags) AddTag(c *gin.Context) {
	jsons := AddTagBase{}
	if c.BindJSON(&jsons) != nil {
		ReturnJsonParseErrJson(c)
		return
	}
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	client, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		t.Logger.Error(err)
		return
	}
	// 退出归还连接
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	_, err = client.AddTag(context.Background(), &rpc.Tag{Text: jsons.Tags})
	st, _ := status.FromError(err)
	c.JSON(http.StatusOK,gin.H{
		"code":st.Code,
		"message":st.Message,
	})
}

func (t *Tags) TagOpt(c *gin.Context) {
	switch c.Param("type") {
	case "id":
		t.GetTagText(c)
		break
	case "text":
		t.GetTagInt32Id(c)
		break
	}
}

func (t *Tags) GetTagInt32Id(c *gin.Context) {
	// 标签名数据经过base64编码，所以需要解码
	tagTextString, err := base64.StdEncoding.DecodeString(c.Param("data"))
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.EncodeError.Code(),
			"message":ecode.EncodeError.Message(),
		})
		return
	}
	// 把按;分割的标签名还原为切片
	tagTextS := utils.SplitStringsToTagList(string(tagTextString))
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	client, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		t.Logger.Error(err)
		return
	}
	// 退出归还连接
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	tag, err := client.GetTagInt32Id(context.Background(), &rpc.Tag{Text: tagTextS})
	st, _ := status.FromError(err)
	if st.Code != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    st.Code,
		"message": st.Message,
		"data":    tag,
	})
}

func (t *Tags) GetTagText(c *gin.Context) {

	// 标签名数据经过base64编码，所以需要解码
	tagIdString, err := base64.StdEncoding.DecodeString(c.Param("data"))
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.EncodeError.Code(),
			"message":ecode.EncodeError.Message(),
		})
		return
	}
	tagIdS := make([]int32,0)
	for _,v := range utils.SplitStringsToTagList(string(tagIdString)) {
		tid,err := strconv.Atoi(v)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"code":ecode.ServerErr.Code(),
				"message":ecode.ServerErr.Message(),
			})
			return
		}
		tagIdS = append(tagIdS,int32(tid))
	}
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	client, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		t.Logger.Error(err)
		return
	}
	// 退出归还连接
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	tag, err := client.GetTagText(context.Background(), &rpc.Tag{Tid: tagIdS})
	st, _ := status.FromError(err)
	if st.Code != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    st.Code,
		"message": st.Message,
		"data":    tag,
	})
}
