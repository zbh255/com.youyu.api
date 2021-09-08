// Package controller 存储一些Business会用到的全局值
package controller

var (
	// TokenSigningKey Jwt签名密钥
	TokenSigningKey []byte
	// TokenExpTime Token 的过期时间,分钟计时
	TokenExpTime int
)
