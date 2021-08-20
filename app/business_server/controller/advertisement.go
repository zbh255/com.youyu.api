package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/errors"
	"com.youyu.api/lib/log"
	"context"
	errs "errors"
	"github.com/gin-gonic/gin"
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
}

func (a *Advertisement) GetAdvertisement(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("advertisement_id"))
	lis,err := ConnectAndConf.ConnPool.Get()
	client, _,err := GetRpcServer(lis,err)
	if err != nil {
		c.JSON(errors.ErrInternalServer.HttpCode,gin.H{
			"code":    errors.ErrInternalServer.Code,
			"message": errors.ErrInternalServer.Message,
		})
		log.Logger.Err(err).Timestamp()
		return
	}
	result, err := client.GetAdvertisement(context.Background(), &rpc.AdvertisementRequest{AdvertisementId: int32(id)})
	// 查看结果是否为0
	if errs.Is(err, errs.New("the query record is zero")) {
		c.JSON(errors.ErrDataBaseResultIsZero.HttpCode, gin.H{
			"code":    errors.ErrDataBaseResultIsZero.Code,
			"message": errors.ErrDataBaseResultIsZero.Message,
			"data":    result,
		})
		return
	}
	if err != nil {
		c.JSON(errors.ErrDatabase.HttpCode, gin.H{
			"code":    errors.ErrDatabase.Code,
			"message": err.Error(),
			"data":    result,
		})
		return
	} else {
		c.JSON(errors.OK.HttpCode, gin.H{
			"code":    errors.OK.Code,
			"message": errors.OK.Message,
			"data":    result,
		})
	}
}

func (a *Advertisement) AddAdvertisement(c *gin.Context) {
	a.advertisementJson = &rpc.Advertisement{}
	err := c.BindJSON(a.advertisementJson)
	a.advertisementJson.AdvertisementId = 0
	if err != nil {
		c.JSON(errors.ErrParamConvert.HttpCode, gin.H{
			"code":    errors.ErrParamConvert.Code,
			"message": errors.ErrParamConvert.Message,
			"data":    nil,
		})
		return
	}
	lis,err := ConnectAndConf.ConnPool.Get()
	client, _,err := GetRpcServer(lis,err)
	if err != nil {
		c.JSON(errors.ErrInternalServer.HttpCode,gin.H{
			"code":    errors.ErrInternalServer.Code,
			"message": errors.ErrInternalServer.Message,
		})
		log.Logger.Err(err).Timestamp()
		return
	}
	_, err = client.AddAdvertisement(context.Background(), a.advertisementJson)
	if err != nil {
		c.JSON(errors.ErrDatabase.HttpCode, gin.H{
			"code":    errors.ErrDatabase.Code,
			"message": err.Error(),
			"data":    nil,
		})
		return
	} else {
		c.JSON(errors.OK.HttpCode, gin.H{
			"code":    errors.OK.Code,
			"message": errors.OK.Message,
			"data":    nil,
		})
	}
}

func (a *Advertisement) UpdateAdvertisement(c *gin.Context) {
	a.advertisementJson = &rpc.Advertisement{}
	err := c.BindJSON(a.advertisementJson)
	a.advertisementJson.AdvertisementId = 0
	if err != nil {
		c.JSON(errors.ErrParamConvert.HttpCode, gin.H{
			"code":    errors.ErrParamConvert.Code,
			"message": errors.ErrParamConvert.Message,
			"data":    nil,
		})
		return
	}
	lis,err := ConnectAndConf.ConnPool.Get()
	client, _,err := GetRpcServer(lis,err)
	if err != nil {
		c.JSON(errors.ErrInternalServer.HttpCode,gin.H{
			"code":    errors.ErrInternalServer.Code,
			"message": errors.ErrInternalServer.Message,
		})
		log.Logger.Err(err).Timestamp()
		return
	}
	_, err = client.UpdateAdvertisement(context.Background(), a.advertisementJson)
	if err != nil {
		c.JSON(errors.ErrDatabase.HttpCode, gin.H{
			"code":    errors.ErrDatabase.Code,
			"message": err.Error(),
			"data":    nil,
		})
		return
	} else {
		c.JSON(errors.OK.HttpCode, gin.H{
			"code":    errors.OK.Code,
			"message": errors.OK.Message,
			"data":    nil,
		})
	}
}

func (a *Advertisement) DelAdvertisement(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("advertisement_id"))
	lis,err := ConnectAndConf.ConnPool.Get()
	client, _,err := GetRpcServer(lis,err)
	if err != nil {
		c.JSON(errors.ErrInternalServer.HttpCode,gin.H{
			"code":    errors.ErrInternalServer.Code,
			"message": errors.ErrInternalServer.Message,
		})
		log.Logger.Err(err).Timestamp()
		return
	}
	_, err = client.DelAdvertisement(context.Background(), &rpc.AdvertisementRequest{AdvertisementId: int32(id)})
	if err != nil {
		c.JSON(errors.ErrDatabase.HttpCode, gin.H{
			"code":    errors.ErrDatabase.Code,
			"message": errors.ErrDatabase.Message,
			"data":    nil,
		})
		return
	} else {
		c.JSON(errors.OK.HttpCode, gin.H{
			"code":    errors.OK.Code,
			"message": errors.OK.Message,
			"data":    nil,
		})
	}
}
