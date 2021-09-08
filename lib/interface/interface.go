package _interface

/*
	此包用与统一项目内的一些IO接口
*/

// 从Rpc配置中心获取配置文件的统一接口

type CentConfig interface {
	GetConfig() ([]byte, error)
	UpdateConfig([]byte) error
}

// Cloud 连接云端的统一接口
type Cloud interface {
}
