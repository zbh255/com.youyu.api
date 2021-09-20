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
	// 文章的评论
	GetArticleComments(c *gin.Context)
	AddArticleComment(c *gin.Context)
	DeleteArticleComment(c *gin.Context)
	UpdateCommentStatus(c *gin.Context)
}

type Article struct {
	// 业务日志
	Logger log.Logger
}

func (a *Article) AddArticle(c *gin.Context) {
	// 解决数据竞争的bug
	article := rpc.Article{}
	err := c.BindJSON(&article)
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
	// 获得uid并根据类型添加评论
	uidString := GetHeaderTokenBindTheUid(c)
	uid,err := strconv.Atoi(uidString)
	if err != nil {
		ReturnServerErrJson(c)
		return
	}
	article.Uid = int64(uid)
	// 文章id数据库操作接口会自动生成
	_, err = client.AddArticle(context.Background(), &article)
	if st, bl := status.FromError(err); bl {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.ServerErr.Code(),
			"message": ecode.ServerErr.Message(),
		})
	}
}

func (a *Article) GetArticles(c *gin.Context) {

}

func (a *Article) GetArticle(c *gin.Context) {
	articleId := c.Param("article_id")
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
	if st, bl := status.FromError(err); bl {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
			"data":    result,
		})
	}
}

// 修复删除文章的漏洞，删除文章只校验了uid,没有校验改文章是否属于该uid
// 如果用户非法输入则会导致删除其他用户的文章
// NOTE:已修复
// 删除文章
func (a *Article) DelArticle(c *gin.Context) {
	articleId := c.Param("article_id")
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
	// 获得uid并根据类型添加评论
	uidString := GetHeaderTokenBindTheUid(c)
	uid,err := strconv.Atoi(uidString)
	if err != nil {
		ReturnServerErrJson(c)
		return
	}
	_, err = client.DelArticle(context.Background(), &rpc.GetArticleRequest{ArticleId: articleId})
	if st, bl := status.FromError(err); bl {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
			"data":    nil,
		})
	}
}

func (a *Article) SetArticle(c *gin.Context) {
	article := rpc.Article{}
	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    ecode.JsonParseError.Code(),
			"message": ecode.JsonParseError.Message(),
			"data":    nil,
		})
		return
	}
	// 获得文章id
	articleId := c.Param("article_id")
	article.ArticleId = articleId
	// 连接data_rpc
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
	// 获得uid并根据类型添加评论
	uidString := GetHeaderTokenBindTheUid(c)
	uid,err := strconv.Atoi(uidString)
	if err != nil {
		ReturnServerErrJson(c)
		return
	}
	article.Uid = int64(uid)
	_, err = client.UpdateArticle(context.Background(), &article)
	if st, bl := status.FromError(err); bl {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
			"data":    nil,
		})
	}
}

// 根据不同的type导向添加与删除点赞\热度
func (a *Article) Options(c *gin.Context) {
	Type := c.Param("type")
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
	articleId := c.Param("article_id")
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
	if st, bl := status.FromError(err); bl {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
			"data":    nil,
		})
	}
}

func (a *Article) AddArticleStatisticsHot(c *gin.Context) {
	articleId := c.Param("article_id")
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
	if st, bl := status.FromError(err); bl {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
			"data":    nil,
		})
	}
}

func (a *Article) AddArticleStatisticsFabulous(c *gin.Context) {
	articleId := c.Param("article_id")
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
	if st, bl := status.FromError(err); bl {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
			"data":    nil,
		})
	}
}

func (a *Article) GetArticleStatistics(c *gin.Context) {
	articleId := c.Param("article_id")
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
	if st, bl := status.FromError(err); bl {
		c.JSON(http.StatusOK, gin.H{
			"code":    st.Code,
			"message": st.Message,
			"data":    result,
		})
	}
}

func (a *Article) GetArticleComments(c *gin.Context) {
	articleId := c.Param("article_id")
	if articleId == "" {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.ParaMeterErr.Code(),
			"message":ecode.ParaMeterErr.Message(),
		})
		return
	}
	// 获得data_rpc
	lis,client,err := LinkDataRpc()
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	if err != nil {
		ReturnServerErrJson(c)
		return
	}
	result, err := client.GetComment(context.Background(),&rpc.CommentSlave{ArticleId: articleId})
	st,_ := status.FromError(err)
	c.JSON(http.StatusOK,gin.H{
		"code":st.Code,
		"message":st.Message,
		"data":result,
	})
}

func (a *Article) AddArticleComment(c *gin.Context) {
	// 绑定参数并校验
	jsons := rpc.CommentSlave{}
	if c.BindJSON(&jsons) != nil {
		ReturnJsonParseErrJson(c)
		return
	}
	jsons.ArticleId = c.Param("article_id")
	if err := jsons.Validate(); err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.ParaMeterErr,
			"message":err.Error(),
		})
		return
	}
	// 获得uid并根据类型添加评论
	uidString := GetHeaderTokenBindTheUid(c)
	uid,err := strconv.Atoi(uidString)
	if err != nil {
		ReturnServerErrJson(c)
		return
	}
	jsons.Uid = int32(uid)
	// 连接data_rpc
	lis,client,err := LinkDataRpc()
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	if err != nil {
		ReturnServerErrJson(c)
		return
	}
	_, err = client.AddComment(context.Background(),&jsons)
	st,_ := status.FromError(err)
	c.JSON(http.StatusOK,gin.H{
		"code":st.Code,
		"message":st.Message,
		"data":nil,
	})
}

func (a *Article) DeleteArticleComment(c *gin.Context) {
	// 绑定参数并校验
	jsons := rpc.CommentSlave{}
	if c.BindJSON(&jsons) != nil {
		ReturnJsonParseErrJson(c)
		return
	}
	jsons.ArticleId = c.Param("article_id")
	if err := jsons.Validate(); err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.ParaMeterErr,
			"message":err.Error(),
		})
		return
	}
	// 获得uid并根据类型添加评论
	uidString := GetHeaderTokenBindTheUid(c)
	uid,err := strconv.Atoi(uidString)
	if err != nil {
		ReturnServerErrJson(c)
		return
	}
	jsons.Uid = int32(uid)
	// 连接data_rpc
	lis,client,err := LinkDataRpc()
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	if err != nil {
		ReturnServerErrJson(c)
		return
	}
	_, err = client.DeleteComment(context.Background(),&jsons)
	st,_ := status.FromError(err)
	c.JSON(http.StatusOK,gin.H{
		"code":st.Code,
		"message":st.Message,
		"data":nil,
	})
}

func (a *Article) UpdateCommentStatus(c *gin.Context) {
	// 绑定参数并校验
	jsons := rpc.UpdateCommentOption{}
	if c.BindJSON(&jsons) != nil {
		ReturnJsonParseErrJson(c)
		return
	}
	jsons.ArticleId = c.Param("article_id")
	if err := jsons.Validate(); err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":ecode.ParaMeterErr,
			"message":err.Error(),
		})
		return
	}
	// 获得uid并根据类型添加评论
	uidString := GetHeaderTokenBindTheUid(c)
	uid,err := strconv.Atoi(uidString)
	if err != nil {
		ReturnServerErrJson(c)
		return
	}
	jsons.Uid = int32(uid)
	// 连接data_rpc
	lis,client,err := LinkDataRpc()
	defer ConnectAndConf.DataRpcConnPool.Put(lis)
	if err != nil {
		ReturnServerErrJson(c)
		return
	}
	_, err = client.UpdateCommentStatus(context.Background(),&jsons)
	st,_ := status.FromError(err)
	c.JSON(http.StatusOK,gin.H{
		"code":st.Code,
		"message":st.Message,
		"data":nil,
	})
}
