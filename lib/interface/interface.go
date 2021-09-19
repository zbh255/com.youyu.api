package _interface

import "github.com/gin-gonic/gin"

/*
	此包用与统一项目内的一些IO接口
*/

// 从Rpc配置中心获取配置文件的统一接口

type CentConfig interface {
	GetConfig() ([]byte, error)
	UpdateConfig([]byte) error
}

// Cloud 连接云端的统一接口
type Cloud interface {}

// 评论 controller统一接口
type CommentApi interface {
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Get(c *gin.Context)
}
