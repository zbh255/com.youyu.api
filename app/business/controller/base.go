package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/ecode/status"
	"com.youyu.api/lib/log"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
)

type Base struct {
	// 业务日志接口
	Logger log.Logger
}

type BaseApi interface {
	GetIndexData(c *gin.Context)
	InitDirection(c *gin.Context)
}

type BaseQuery struct {
	Position string `form:"position" binding:"required"`
	Type     string `form:"type" binding:"required"`
	ClientId string `form:"client_id" binding:"-"`
}

// GetIndexData 返回渲染首页需要的广告和文章的数据
func (b *Base) GetIndexData(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))
	op := &rpc.ArticleOptions{
		Type:    c.Query("order_type"),
		Order:   c.Query("order"),
		Page:    int32(page),
		PageNum: int32(pageNum),
	}
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	client, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		b.Logger.Error(err)
		return
	}
	// 退出归还连接
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	// 查询文章
	articleResults, err1 := client.GetArticleList(context.Background(), op)
	st1, _ := status.FromError(err1)
	// 查询广告
	advertisementResults, err2 := client.GetAdvertisementList(context.Background(), op)
	st2, _ := status.FromError(err2)
	if st1.Code != 0 || st2.Code != 0 {
		st := &status.Status{}
		if st1.Code != 0 {
			st = st1
		}
		if st2.Code != 0 {
			st = st2
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
			"data":    nil,
		})
		return
	} else {

		c.JSON(http.StatusOK, gin.H{
			"code":              st1.Code,
			"message":           st1.Message,
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
	case "client_data":
		b.GetClientData(c)
		break
	}
}

// GetClientData 返回客户端需要的数据
func (b *Base) GetClientData(c *gin.Context) {
	jsons := BaseQuery{}
	if c.ShouldBindQuery(&jsons) != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.UrlParseError.Code(),
			"message": ecode.UrlParseError.Message(),
		})
		return
	}
	if jsons.Type == "key" {
		// 返回一个公钥
		// 检验UUid
		UUid, err := uuid.FromString(jsons.ClientId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    ecode.ClientIdError.Code(),
				"message": ecode.ClientIdError.Message(),
			})
			return
		}
		// 验证成功以后签钥
		lis, err := ConnectAndConf.SecretKeyRpcConnPool.Get()
		client, _, err := GetSecretKeyRpcServer(lis, err)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    ecode.ServerErr.Code(),
				"message": ecode.ServerErr.Message(),
			})
			return
		}
		defer ConnectAndConf.SecretKeyRpcConnPool.Put(lis)
		Key, err := client.GetPublicKey(context.Background(), &rpc.RsaKey{ClientId: UUid.String()})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    ecode.ServerErr.Code(),
				"message": ecode.ServerErr.Message(),
			})
			return
		}
		// 流程完成返回公钥
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.OK.Code(),
			"message": ecode.OK.Message(),
			"key":     Key.PublicKey,
		})
	} else if jsons.Type == "client_id" {
		// 返回一个客户端id
		UUid := uuid.NewV4()
		c.JSON(http.StatusOK, gin.H{
			"code":      ecode.OK.Code(),
			"message":   ecode.OK.Message(),
			"client_id": UUid.String(),
		})
	}
}
