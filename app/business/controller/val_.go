// Package controller 存储一些Business会用到的全局值
package controller

var (
	// TokenSigningKey Jwt签名密钥
	TokenSigningKey []byte
	// TokenExpTime Token 的过期时间,分钟计时
	TokenExpTime int
	// WechatLoginRawUrl 微信oauth2服务端的请求url,参数按顺序为,appid,appsecret,js_code
	WechatLoginRawUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)
