package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/errors"
	"com.youyu.api/lib/log"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Base struct {
}

type BaseApi interface {
	GetIndexData(c *gin.Context)
	InitDirection(c *gin.Context)
}

// 返回渲染首页需要的广告和文章的数据
func (b *Base) GetIndexData(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))
	op := &rpc.ArticleOptions{
		Type:    c.Query("order_type"),
		Order:   c.Query("order"),
		Page:    int32(page),
		PageNum: int32(pageNum),
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
	// 退出归还连接
	defer ConnectAndConf.ConnPool.Put(lis)
	// 查询文章
	articleResults, err1 := client.GetArticleList(context.Background(), op)
	// 查询广告
	advertisementResults, err2 := client.GetAdvertisementList(context.Background(), op)
	if err1 != nil || err2 != nil {
		c.JSON(errors.ErrDatabase.HttpCode, gin.H{
			"code":    errors.ErrDatabase.Code,
			"message": errors.ErrDatabase.Message,
			"data":    nil,
		})
		log.Logger.Err(err1).Timestamp()
		return
	} else {
		c.JSON(errors.OK.HttpCode, gin.H{
			"code":    errors.OK.Code,
			"message": errors.OK.Message,
			"data": []interface{}{
				advertisementResults,
				articleResults,
			},
		})
	}
}

func (b *Base) InitDirection(c *gin.Context) {
	switch c.Query("position") {
	case "index":
		b.GetIndexData(c)
		break
	}
}
