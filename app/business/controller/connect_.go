package controller

import (
	"com.youyu.api/app/rpc/data/option"
	"com.youyu.api/lib/config"
)

// 该文件将grpc的连接和business用到的配置存放统一

type ConnectAndConfig struct {
	Config        *config.AutoGenerated
	DataBaseLink  *option.MysqlApiServer
	SecretKeyLink *option.SecretKeyApiServer
}

var ConnectAndConf *ConnectAndConfig

// 拿一个初始化好的数据库连接
func TakeDataBaseLink() *option.MysqlApiServer {
	return ConnectAndConf.DataBaseLink
}

// 拿一个token和密钥rpc的连接
func TakeSecretKeyLink() *option.SecretKeyApiServer {
	return ConnectAndConf.SecretKeyLink
}