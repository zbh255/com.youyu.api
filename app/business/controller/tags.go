package controller

import (
	"com.youyu.api/lib/log"
	"github.com/gin-gonic/gin"
)

type TagsApi interface {
	AddTag(c *gin.Context)
	GetTagInt32Id(c *gin.Context)
	GetTagText(c *gin.Context)
}

type Tags struct {
	Logger log.Logger
}

func (t *Tags) AddTag(c *gin.Context) {
	panic("implement me")
}

func (t *Tags) GetTagInt32Id(c *gin.Context) {
	panic("implement me")
}

func (t *Tags) GetTagText(c *gin.Context) {
	panic("implement me")
}

