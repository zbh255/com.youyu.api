// Package controller 存储一些Business会用到的全局值
package controller

import rpc "com.youyu.api/app/rpc/proto_files"

var (
	// TokenSigningKey Jwt签名密钥
	TokenSigningKey []byte
	// TokenExpTime Token 的过期时间,分钟计时
	TokenExpTime int
	// WechatLoginRawUrl 微信oauth2服务端的请求url,参数按顺序为,appid,appsecret,js_code
	WechatLoginRawUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	// 数据库排序默认选项
	// 一页包含的数据
	MaxPrePage = 30
	// 页数
	Page = 1
	// 排序表驱动
	OrderTableDriver = map[string]rpc.OrderOptions{
		"created-desc": {Order: "create_time", Type: "desc"},
		"created-asc":{Order: "create_time", Type: "asc"},
		"fabulous-desc": {Order: "fabulous", Type: "desc"},
		"fabulous-asc":{Order: "fabulous",Type: "asc"},
	}
)
