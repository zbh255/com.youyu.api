package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/ecode/status"
	"com.youyu.api/lib/log"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"net/http"
)

type Base struct {
	// 业务日志接口
	Logger log.Logger
}

type BaseApi interface {
	GetIndexData(c *gin.Context)
	InitDirection(c *gin.Context)
}


// GetIndexData 返回渲染首页需要的广告和文章的数据
func (b *Base) GetIndexData(c *gin.Context) {
	jsons := rpc.ArticleOptions{}
	err := c.BindJSON(&jsons)
	if err != nil {
		ReturnJsonParseErrJson(c)
		return
	}
	client := TakeDataBaseLink()
	if client == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
		b.Logger.Error(errors.New("nil ptr"))
		return
	}
	// 查询文章
	articleResults, err1 := client.GetArticleList(context.Background(), &jsons)
	st1, _ := status.FromError(err1)
	// 查询广告
	advertisementResults, err2 := client.GetAdvertisementList(context.Background(), &jsons)
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
	switch c.Param("position") {
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
	typeString := c.Param("type")
	if typeString == "key" {
		// 返回一个公钥
		// 检验UUid,已由中间件检验
		// 验证成功以后签钥
		client := TakeSecretKeyLink()
		if client == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    ecode.ServerErr.Code(),
				"message": ecode.ServerErr.Message(),
			})
			b.Logger.Error(errors.New("nil ptr"))
			return
		}
		// 获得UUID
		UUID := c.Request.Header.Get("Client-Id")
		Key, err := client.GetPublicKey(context.Background(), &rpc.RsaKey{ClientId: UUID})
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
	} else if typeString == "client_id" {
		// 返回一个客户端id
		UUid := uuid.NewV4()
		c.JSON(http.StatusOK, gin.H{
			"code":      ecode.OK.Code(),
			"message":   ecode.OK.Message(),
			"client_id": UUid.String(),
		})
	}
}
