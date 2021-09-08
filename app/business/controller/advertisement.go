package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/ecode"
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
	advertisementJson *rpc.Advertisement
	Logger            log.Logger
}

func (a *Advertisement) GetAdvertisement(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("advertisement_id"))
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
	// TODO:清理errors.Is
	//if errs.Is(err, errs.New("the query record is zero")) {
	//	c.JSON(errors.ErrDataBaseResultIsZero.HttpCode, gin.H{
	//		"code":    errors.ErrDataBaseResultIsZero.Code,
	//		"message": errors.ErrDataBaseResultIsZero.Message,
	//		"data":    result,
	//	})
	//	return
	//}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.GetAdvertisementErr.Code(),
			"message": ecode.GetAdvertisementErr.Message(),
			"data":    result,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.OK.Code(),
			"message": ecode.OK.Message(),
			"data":    result,
		})
	}
}

func (a *Advertisement) AddAdvertisement(c *gin.Context) {
	a.advertisementJson = &rpc.Advertisement{}
	err := c.BindJSON(a.advertisementJson)
	a.advertisementJson.AdvertisementId = 0
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
	_, err = client.AddAdvertisement(context.Background(), a.advertisementJson)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.AddAdvertisementErr.Code(),
			"message": ecode.AddAdvertisementErr.Message(),
			"data":    nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.OK.Code(),
			"message": ecode.OK.Message(),
			"data":    nil,
		})
	}
}

func (a *Advertisement) UpdateAdvertisement(c *gin.Context) {
	a.advertisementJson = &rpc.Advertisement{}
	err := c.BindJSON(a.advertisementJson)
	a.advertisementJson.AdvertisementId = 0
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
	_, err = client.UpdateAdvertisement(context.Background(), a.advertisementJson)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.UpdAdvertisementErr.Code(),
			"message": ecode.UpdAdvertisementErr.Message(),
			"data":    nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.OK.Code(),
			"message": ecode.OK.Message(),
			"data":    nil,
		})
	}
}

func (a *Advertisement) DelAdvertisement(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("advertisement_id"))
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
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.DelAdvertisementErr.Code(),
			"message": ecode.DelAdvertisementErr.Message(),
			"data":    nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.OK.Code(),
			"message": ecode.OK.Message(),
			"data":    nil,
		})
	}
}
