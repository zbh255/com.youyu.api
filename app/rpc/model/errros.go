// 定义包级别的错误
package model

import "errors"

var (
	// CreateSameExistence 写入的数据不能跟数据表中已存在的内容一样
	CreateSameExistence = errors.New("the same content exists when it is created")
	// UserNameAlreadyExists 用户名已经存在
	UserNameAlreadyExists = errors.New("the user name already exists")
	// UserDoesNotExist 用户不存在
	UserDoesNotExist = errors.New("the user does not exist")
	// 用户密码或用户名错误
	UserPasswordORUserNameErr = errors.New("user password or user name error")
)
