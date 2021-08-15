package controller

import (
	"com.youyu.api/common/config"
	"github.com/silenceper/pool"
)

// 该文件将grpc的连接和business用到的配置存放统一

type ConnectAndConfig struct {
	Config      *config.AutoGenerated
	ConnPool        pool.Pool
}

var ConnectAndConf *ConnectAndConfig