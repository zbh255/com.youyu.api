package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/log"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleApi interface {
	AddArticle(c *gin.Context)
	GetArticles(c *gin.Context)
	GetArticle(c *gin.Context)
	DelArticle(c *gin.Context)
	SetArticle(c *gin.Context)
	// Options 文章的热度，评论数，点赞接口
	Options(c *gin.Context)
	ReduceArticleStatisticsFabulous(c *gin.Context)
	AddArticleStatisticsHot(c *gin.Context)
	AddArticleStatisticsFabulous(c *gin.Context)
	GetArticleStatistics(c *gin.Context)
}

type Article struct {
	articleJson       *rpc.Article
	articleStatistics *rpc.ArticleStatistics
	// 业务日志
	Logger log.Logger
}

func (a *Article) AddArticle(c *gin.Context) {
	a.articleJson = &rpc.Article{}
	err := c.BindJSON(a.articleJson)
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
	// 退出归还连接
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	// 使用uuid作为文章的id
	a.articleJson.ArticleId = "0"
	// TODO:使用登录的用户uid作为文章编写者
	a.articleJson.Uid = 1
	_, err = client.AddArticle(context.Background(), a.articleJson)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.AddArticleErr.Code(),
			"message": ecode.AddArticleErr.Message(),
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

func (a *Article) GetArticles(c *gin.Context) {

}

func (a *Article) GetArticle(c *gin.Context) {
	articleId := c.Query("article_id")
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
	// 退出归还连接
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	result, err := client.GetArticle(context.Background(), &rpc.GetArticleRequest{ArticleId: articleId})
	// 查看结果是否为0
	// TODO: 交给前端判断
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
			"code":    ecode.GetArticleErr.Code(),
			"message": ecode.GetArticleErr.Message(),
			"data":    nil,
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

// 删除文章
func (a *Article) DelArticle(c *gin.Context) {
	articleId := c.Query("article_id")
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
	// 退出归还连接
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	_, err = client.DelArticle(context.Background(), &rpc.GetArticleRequest{ArticleId: articleId})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.DelArticleErr.Code(),
			"message": ecode.DelArticleErr.Message(),
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

func (a *Article) SetArticle(c *gin.Context) {
	a.articleJson = &rpc.Article{}
	err := c.BindJSON(a.articleJson)
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
	// 退出归还连接
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	_, err = client.UpdateArticle(context.Background(), a.articleJson)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.UpdArticleErr.Code(),
			"message": ecode.UpdArticleErr.Message(),
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

// 根据不同的type导向添加与删除点赞\热度
func (a *Article) Options(c *gin.Context) {
	Type := c.Query("type")
	switch Type {
	case "hot":
		a.AddArticleStatisticsHot(c)
		break
	case "fabulous":
		a.AddArticleStatisticsFabulous(c)
		break
	}
}

// ReduceArticleStatisticsFabulous 文章的点赞数-1
func (a *Article) ReduceArticleStatisticsFabulous(c *gin.Context) {
	articleId := c.Query("article_id")
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
	// 退出归还连接
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	_, err = client.DelArticleStatisticsFabulous(context.Background(), &rpc.GetArticleRequest{ArticleId: articleId})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.DelArticleFabulousErr.Code(),
			"message": ecode.DelArticleFabulousErr.Message(),
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

func (a *Article) AddArticleStatisticsHot(c *gin.Context) {
	articleId := c.Query("article_id")
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
	// 退出归还连接
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	_, err = client.AddArticleStatisticsHot(context.Background(), &rpc.GetArticleRequest{ArticleId: articleId})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.AddArticleHotErr.Code(),
			"message": ecode.AddArticleHotErr.Message(),
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

func (a *Article) AddArticleStatisticsFabulous(c *gin.Context) {
	articleId := c.Query("article_id")
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
	// 退出归还连接
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	_, err = client.AddArticleStatisticsFabulous(context.Background(), &rpc.GetArticleRequest{ArticleId: articleId})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.AddArticleFabulousErr.Code(),
			"message": ecode.AddArticleFabulousErr.Code(),
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.OK.Code(),
			"message": ecode.OK.Message(),
			"data":    nil,
		})
	}
}

func (a *Article) GetArticleStatistics(c *gin.Context) {
	articleId := c.Query("article_id")
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
	// 退出归还连接
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	result, err := client.GetArticleStatistics(context.Background(), &rpc.GetArticleRequest{ArticleId: articleId})
	// 查看结果是否为0
	// TODO: 使用Casus功能进行修改
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
			"code":    ecode.GetArticleStatisticsErr.Code(),
			"message": ecode.GetArticleStatisticsErr.Message(),
			"data":    result,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.OK.Code(),
			"message": ecode.OK.Message(),
			"data":    result,
		})
	}
}
