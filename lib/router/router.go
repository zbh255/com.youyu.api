package router

import (
	"com.youyu.api/app/business/controller"
	"com.youyu.api/lib/log"
	"com.youyu.api/lib/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter 传入Gin Engine和业务日志接口
func InitRouter(r *gin.Engine, logger log.Logger) {
	var ControllerArticle controller.ArticleApi = &controller.Article{Logger: logger}
	var ControllerAdvertisement controller.AdvertisementApi = &controller.Advertisement{Logger: logger}
	var ControllerBase controller.BaseApi = &controller.Base{Logger: logger}
	ControllerTags := controller.TagsApi(&controller.Tags{Logger: logger})
	v1 := r.Group("/v1")
	v1.Use(middleware.ClientIdAuth())
	// 普通用户请求的api
	// ?article_id=0
	//
	v1.GET("/article", ControllerArticle.GetArticle)
	// 返回一些用于渲染页面的基本数据
	// 返回列表的请求参数: ?position=index&type=list&page=3&page_num=4&order=hot&orderType=desc
	// ?position=client_data&type=key&client_id=678-890-444-777
	// ?position=client_data&type=client_id
	v1.GET("/base", ControllerBase.InitDirection)

	// ?article_id=0
	v1.GET("/article_statistics", ControllerArticle.GetArticleStatistics)

	// ?advertisement_id=0
	v1.GET("/advertisement", ControllerAdvertisement.GetAdvertisement)

	v1.POST("/advertisement", ControllerAdvertisement.AddAdvertisement)
	v1.PUT("/advertisement", ControllerAdvertisement.UpdateAdvertisement)
	// ?advertisement_id=0
	v1.DELETE("/advertisement", ControllerAdvertisement.DelAdvertisement)

	// ?type=text&text=haha;chacha;baba;lala
	// ?type=id&id=11;22;33;44;55
	// TODO:ClientId 放在http header参数里面
	v1.GET("/tag",ControllerTags.TagOpt)

	// v1鉴权
	v1.Use(middleware.JwtAuth())
	{
		// ?type=add&data=xxx
		v1.POST("/tag",ControllerTags.TagOpt)
		// ?article_id=0&type=hot
		v1.PUT("/article_statistics", ControllerArticle.Options)
		// ?article_id=0&type=fabulous
		v1.DELETE("/article_statistics", ControllerArticle.ReduceArticleStatisticsFabulous)
		// ?article_id=0
		v1.PUT("/article", ControllerArticle.SetArticle)
		// ?article_id=0
		v1.DELETE("/article", ControllerArticle.DelArticle)
		// json参数
		v1.POST("/article", ControllerArticle.AddArticle)
	}

	// 自动鉴权接口，有无鉴权时会返回不同的响应
	ControllerLogin := controller.SignAndLoginApi(&controller.SignAndLogin{Logger: logger})
	v1AutoAuth := r.Group("/v1")
	v1AutoAuth.Use(middleware.ClientIdAuth())
	v1AutoAuth.POST("/login", ControllerLogin.CreateLoginState)
	v1AutoAuth.POST("/sign", ControllerLogin.CreateSign)

	v1AutoAuth.Use(middleware.JwtAuth())
	{
		v1AutoAuth.DELETE("/login", ControllerLogin.DeleteLoginState)
		v1AutoAuth.DELETE("/sign", ControllerLogin.DeleteSign)
	}
}
