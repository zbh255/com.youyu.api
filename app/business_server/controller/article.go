package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/errors"
	"com.youyu.api/lib/log"
	"context"
	errs "errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ArticleApi interface {
	AddArticle(c *gin.Context)
	GetArticles(c *gin.Context)
	GetArticle(c *gin.Context)
	DelArticle(c *gin.Context)
	SetArticle(c *gin.Context)
	// 文章的热度，评论数，点赞接口
	Options(c *gin.Context)
	ReduceArticleStatisticsFabulous(c *gin.Context)
	AddArticleStatisticsHot(c *gin.Context)
	AddArticleStatisticsFabulous(c *gin.Context)
	GetArticleStatistics(c *gin.Context)
}

type Article struct {
	articleJson       *rpc.Article
	articleStatistics *rpc.ArticleStatistics
}

// 向连接池取得连接
func GetRpcServer(meta interface{},err error) (rpc.MysqlApiClient, *grpc.ClientConn,error) {
	if err != nil {
		return nil,nil,err
	} else {
		m := meta.(*[2]interface{})
		return m[0].(rpc.MysqlApiClient),m[1].(*grpc.ClientConn),nil
	}
}


func (a *Article) AddArticle(c *gin.Context) {
	a.articleJson = &rpc.Article{}
	err := c.BindJSON(a.articleJson)
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
	// 退出归还连接
	defer ConnectAndConf.ConnPool.Put(lis)
	// 使用uuid作为文章的id
	a.articleJson.ArticleId = "0"
	// TODO:使用登录的用户uid作为文章编写者
	a.articleJson.Uid = 1
	_, err = client.AddArticle(context.Background(), a.articleJson)
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

func (a *Article) GetArticles(c *gin.Context) {

}

func (a *Article) GetArticle(c *gin.Context) {
	articleId := c.Query("article_id")
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
	result, err := client.GetArticle(context.Background(), &rpc.GetArticleRequest{ArticleId: articleId})
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
			"message": errors.ErrDatabase.Message,
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

func (a *Article) DelArticle(c *gin.Context) {
	articleId := c.Query("article_id")
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
	result, err := client.DelArticle(context.Background(), &rpc.GetArticleRequest{ArticleId: articleId})
	if err != nil {
		c.JSON(errors.ErrDatabase.HttpCode, gin.H{
			"code":    errors.ErrDatabase.Code,
			"message": errors.ErrDatabase.Message,
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

func (a *Article) SetArticle(c *gin.Context) {
	a.articleJson = &rpc.Article{}
	err := c.BindJSON(a.articleJson)
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
	// 退出归还连接
	defer ConnectAndConf.ConnPool.Put(lis)
	_, err = client.UpdateArticle(context.Background(), a.articleJson)
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

func (a *Article) ReduceArticleStatisticsFabulous(c *gin.Context) {
	articleId := c.Query("article_id")
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
	_, err = client.DelArticleStatisticsFabulous(context.Background(), &rpc.GetArticleRequest{ArticleId: articleId})
	if err != nil {
		c.JSON(errors.ErrDatabase.HttpCode, gin.H{
			"code":    errors.ErrDatabase.Code,
			"message": err.Error(),
			"data":    nil,
		})
		log.Logger.Err(err).Timestamp()
		return
	} else {
		c.JSON(errors.OK.HttpCode, gin.H{
			"code":    errors.OK.Code,
			"message": errors.OK.Message,
			"data":    nil,
		})
	}
}

func (a *Article) AddArticleStatisticsHot(c *gin.Context) {
	articleId := c.Query("article_id")
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
	_, err = client.AddArticleStatisticsHot(context.Background(), &rpc.GetArticleRequest{ArticleId: articleId})
	if err != nil {
		c.JSON(errors.ErrDatabase.HttpCode, gin.H{
			"code":    errors.ErrDatabase.Code,
			"message": err.Error(),
			"data":    nil,
		})
		log.Logger.Err(err).Timestamp()
		return
	} else {
		c.JSON(errors.OK.HttpCode, gin.H{
			"code":    errors.OK.Code,
			"message": errors.OK.Message,
			"data":    nil,
		})
	}
}

func (a *Article) AddArticleStatisticsFabulous(c *gin.Context) {
	articleId := c.Query("article_id")
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
	_, err = client.AddArticleStatisticsFabulous(context.Background(), &rpc.GetArticleRequest{ArticleId: articleId})
	if err != nil {
		c.JSON(errors.ErrDatabase.HttpCode, gin.H{
			"code":    errors.ErrDatabase.Code,
			"message": err.Error(),
			"data":    nil,
		})
		log.Logger.Err(err).Timestamp()
	} else {
		c.JSON(errors.OK.HttpCode, gin.H{
			"code":    errors.OK.Code,
			"message": errors.OK.Message,
			"data":    nil,
		})
	}
}

func (a *Article) GetArticleStatistics(c *gin.Context) {
	articleId := c.Query("article_id")
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
	result, err := client.GetArticleStatistics(context.Background(), &rpc.GetArticleRequest{ArticleId: articleId})
	// 查看结果是否为0
	// TODO: 使用Casus功能进行修改
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
			"message": errors.ErrDatabase.Message,
			"data":    result,
		})
	} else {
		c.JSON(errors.OK.HttpCode, gin.H{
			"code":    errors.OK.Code,
			"message": errors.OK.Message,
			"data":    result,
		})
	}
}
