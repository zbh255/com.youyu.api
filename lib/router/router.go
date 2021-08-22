package router

import (
	"com.youyu.api/app/business/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	var ControllerArticle controller.ArticleApi = &controller.Article{}
	var ControllerAdvertisement controller.AdvertisementApi = &controller.Advertisement{}
	var ControllerBase controller.BaseApi = &controller.Base{}
	v1 := r.Group("/v1")
	// 普通用户请求的api
	// json参数
	v1.POST("/article", ControllerArticle.AddArticle)
	// ?article_id=0
	//
	v1.GET("/article", ControllerArticle.GetArticle)
	// 返回一些用于渲染页面的基本数据
	// 返回列表的请求参数: ?position=index&type=list&page=3&page_num=4&order=hot&orderType=desc
	v1.GET("/base", ControllerBase.InitDirection)
	// ?article_id=0
	v1.PUT("/article", ControllerArticle.SetArticle)
	// ?article_id=0
	v1.DELETE("/article", ControllerArticle.DelArticle)

	// ?article_id=0
	v1.GET("/article_statistics", ControllerArticle.GetArticleStatistics)
	// ?article_id=0&type=hot
	v1.PUT("/article_statistics", ControllerArticle.Options)
	// ?article_id=0&type=fabulous
	v1.DELETE("/article_statistics", ControllerArticle.ReduceArticleStatisticsFabulous)

	// ?advertisement_id=0
	v1.GET("/advertisement", ControllerAdvertisement.GetAdvertisement)

	v1.POST("/advertisement", ControllerAdvertisement.AddAdvertisement)
	v1.PUT("/advertisement", ControllerAdvertisement.UpdateAdvertisement)
	// ?advertisement_id=0
	v1.DELETE("/advertisement", ControllerAdvertisement.DelAdvertisement)
}
