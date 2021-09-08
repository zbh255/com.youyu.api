// 用户接口的控制器
package controller

type UserApi interface {
	GetUserInfo(user string)
	UpdateUserInfo(user string)
}
