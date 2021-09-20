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
)

type AdvertisementApi interface {
	GetAdvertisement(c *gin.Context)
	AddAdvertisement(c *gin.Context)
	UpdateAdvertisement(c *gin.Context)
	DelAdvertisement(c *gin.Context)
}

type Advertisement struct {
	Logger            log.Logger
}

func (a *Advertisement) GetAdvertisement(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("advertisement_id"))
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	client, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		a.Logger.Error(err)
		return
	}
	result, err := client.GetAdvertisement(context.Background(), &rpc.AdvertisementRequest{AdvertisementId: int32(id)})
	// 查看结果是否为0
	if st, bl := status.FromError(err); bl {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
			"data":    result,
		})
	}
}

func (a *Advertisement) AddAdvertisement(c *gin.Context) {
	ad := rpc.Advertisement{}
	err := c.BindJSON(&ad)
	ad.AdvertisementId = 0
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.JsonParseError.Code(),
			"message": ecode.JsonParseError.Message(),
			"data":    nil,
		})
		return
	}
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	client, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		a.Logger.Error(err)
		return
	}
	_, err = client.AddAdvertisement(context.Background(), &ad)
	if st, bl := status.FromError(err); bl {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
			"data":    nil,
		})
	}
}

func (a *Advertisement) UpdateAdvertisement(c *gin.Context) {
	ad := rpc.Advertisement{}
	err := c.BindJSON(&ad)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.JsonParseError.Code(),
			"message": ecode.JsonParseError.Message(),
			"data":    nil,
		})
		return
	}
	// 获得广告id
	adId := c.Param("advertisement_id")
	aid, err := strconv.Atoi(adId)
	if err != nil {
		ReturnServerErrJson(c)
		return
	}
	ad.AdvertisementId = int32(aid)
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	client, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		a.Logger.Error(err)
		return
	}
	_, err = client.UpdateAdvertisement(context.Background(), &ad)
	if st, bl := status.FromError(err); bl {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
			"data":    nil,
		})
	}
}

func (a *Advertisement) DelAdvertisement(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("advertisement_id"))
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	client, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		a.Logger.Error(err)
		return
	}
	_, err = client.DelAdvertisement(context.Background(), &rpc.AdvertisementRequest{AdvertisementId: int32(id)})
	if st, bl := status.FromError(err); bl {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
			"data":    nil,
		})
	}
}
