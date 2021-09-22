package router

import (
	"com.youyu.api/app/business/controller"
	"com.youyu.api/lib/log"
	"com.youyu.api/lib/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter 传入Gin Engine和业务日志接口
// TODO:2021/09/20/17-52-34 使用Rest Level 2 重构接口
// TODO:/x/x/x/x 为需要修改成的接口原型
func InitRouter(r *gin.Engine, logger log.Logger) {
	var ControllerArticle controller.ArticleApi = &controller.Article{Logger: logger}
	var ControllerAdvertisement controller.AdvertisementApi = &controller.Advertisement{Logger: logger}
	var ControllerBase controller.BaseApi = &controller.Base{Logger: logger}
	ControllerLogin := controller.SignAndLoginApi(&controller.SignAndLogin{Logger: logger})
	ControllerTags := controller.TagsApi(&controller.Tags{Logger: logger})
	ControllerVerification := controller.VerificationApi(&controller.Verification{Logger: logger})
	ControllerUser := controller.UserApi(&controller.UserInfo{Logger: logger})
	ControllerResources := controller.ReSourcesApi(&controller.ReSources{Logger: logger})
	v1 := r.Group("/v1")
	v1.Use(middleware.ClientIdAuth())
	// 普通用户请求的api
	// ?article_id=0
	//TODO /article/article_id
	v1.GET("/article/:article_id", ControllerArticle.GetArticle)
	// 文章评论接口,文章评论修改接口不对不同用户开放
	//TODO /article/article_id/comment
	v1.GET("/article/:article_id/comment",ControllerArticle.GetArticleComments)
	v1.GET("/article/:article_id/comment/:comment_mid", ControllerArticle.GetArticleSubComments)

	// 返回一些用于渲染页面的基本数据
	// 返回列表的请求参数: ?position=index&type=list&page=3&page_num=4&order=hot&orderType=desc
	// ?position=client_data&type=key&client_id=678-890-444-777
	// ?position=client_data&type=client_id
	// TODO application/json 代替url参数 + /base/type
	v1.GET("/base/:position/:type", ControllerBase.InitDirection)

	// ?article_id=0
	// TODO GET article/article_id/statistics
	v1.GET("/article/:article_id/statistics", ControllerArticle.GetArticleStatistics)

	// ?advertisement_id=0
	// TODO GET /advertisement/advertisement_id
	v1.GET("/advertisement/:advertisement_id", ControllerAdvertisement.GetAdvertisement)


	// ?type=text&text=haha;chacha;baba;lala
	// ?type=id&id=11;22;33;44;55
	// TODO:ClientId 放在http header参数里面
	// TODO: /tag/id/tag_id
	// TODO: /tag/text/tag_text
	v1.GET("/tag/:type/:data",ControllerTags.TagOpt)

	// 发送验证码的接口
	// ?type=auth_code&addr_type=email&addr=base64Str
	// TODO POST application/json 代替url参数
	v1.POST("/user/check/code",ControllerVerification.SendVerificationCode)

	// 自动鉴权接口，有无鉴权时会返回不同的响应
	// TODO: POST /user/login/
	v1.POST("/user/login", ControllerLogin.CreateLoginState)
	// TODO: POST /user/sign
	v1.POST("/user/sign", ControllerLogin.CreateSign)

	// 第三方登录接口
	// ?protocol=wxlogin&type=wechat_login&code=xxxx
	// TODO POST application/json 代替url参数
	v1.POST("/user/login/other",ControllerVerification.OtherLogin)

	// 操作用户的接口
	// TODO GET /user/uid/info
	v1.Use(middleware.AutoJwtAuth()).GET("/user/:uid/info",ControllerUser.GetUserInfo)

	// 获取用户旗下的资源
	// 获得用户撰写过的文章
	v1.GET("/user/:uid/articles")

	// v1鉴权
	v1.Use(middleware.JwtAuth())
	{
		// ?type=add&data=xxx
		// TODO POST application/json 代替url参数
		v1.POST("/tag",ControllerTags.AddTag)
		// ?article_id=0&type=hot
		// TODO PUT /article/:article_id/statistics/:type
		v1.PUT("/article/:article_id/statistics/:type", ControllerArticle.Options)
		// ?article_id=0&type=fabulous
		// TODO DELETE /article/:article_id/statistics/:type
		v1.DELETE("/article/:article_id/statistics/fabulous", ControllerArticle.ReduceArticleStatisticsFabulous)
		// ?article_id=0
		// TODO PUT /article/:article_id
		v1.PUT("/article/:article_id", ControllerArticle.SetArticle)
		// ?article_id=0
		// TODO DELETE /article/:article_id
		v1.DELETE("/article/:article_id", ControllerArticle.DelArticle)
		// json参数
		v1.POST("/article", ControllerArticle.AddArticle)
		// 获得操作上传数据权限的接口
		// 头像
		v1.GET("/resources/headPortrait",ControllerResources.GetUploadHeadPortraitToken)
		// 文章图片
		v1.GET("/resources/image/:name",ControllerResources.GetUploadArticleImageToken)
		// 文章视频
		v1.GET("/resources/video/:name",ControllerResources.GetUploadArticleVideoToken)
		// 更新用户信息
		// TODO PUT /user/info
		v1.PUT("/user/info",ControllerUser.UpdateUserInfo)
		// 添加用户验证信息
		// TODO PUT /user/check
		v1.PUT("/user/check",ControllerUser.AddUserCheckInfo)
		// 需要鉴权的文章评论接口
		// TODO POST /article/article_id/comment
		v1.POST("/article/:article_id/comment",ControllerArticle.AddArticleComment)
		// TODO PUT /article/article_id/comment
		v1.PUT("/article/:article_id/comment",ControllerArticle.UpdateCommentStatus)
		// TODO DELETE /article/article_id/comment
		v1.DELETE("/article/:article_id/comment",ControllerArticle.DeleteArticleComment)
		// TODO DELETE /user/login
		v1.DELETE("/user/login", ControllerLogin.DeleteLoginState)
		// TODO DELETE /user/sign
		v1.DELETE("/user/sign", ControllerLogin.DeleteSign)
	}

	// TODO: 高等级鉴权
	v1.Use(middleware.JwtAuth())
	{
		// TODO /advertisement
		v1.POST("/advertisement", ControllerAdvertisement.AddAdvertisement)
		// TODO /advertisement/advertisement_id
		v1.PUT("/advertisement/:advertisement_id", ControllerAdvertisement.UpdateAdvertisement)
		// ?advertisement_id=0
		// TODO /advertisement/advertisement_id
		v1.DELETE("/advertisement/:advertisement_id", ControllerAdvertisement.DelAdvertisement)
	}
}
