// 负责处理评论的控制器
package controller

import (
	"com.youyu.api/lib/interface"
	"com.youyu.api/lib/log"
	"github.com/gin-gonic/gin"
)

type CommentApi _interface.CommentApi

type ArticleComment struct {
	Logger log.Logger
}

func (a *ArticleComment) Add(c *gin.Context) {
	panic("implement me")
}

func (a *ArticleComment) Update(c *gin.Context) {
	panic("implement me")
}

func (a *ArticleComment) Delete(c *gin.Context) {
	panic("implement me")
}

func (a *ArticleComment) Get(c *gin.Context) {
	panic("implement me")
}
