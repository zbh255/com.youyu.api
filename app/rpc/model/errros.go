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
	// UserPasswordORUserNameErr 用户密码或用户名错误
	UserPasswordORUserNameErr = errors.New("user password or user name error")
	// ArticleIdAlreadyExists 文章id已存在"
	ArticleIdAlreadyExists = errors.New("the article id already exists")
	// ArticleIdNotExists 文章id不存在
	ArticleIdNotExists = errors.New("the article id not exists")
	// AdvertisementIdNotExists 广告id不存在
	AdvertisementIdNotExists = errors.New("the advertisement id does not exist")
	// TagNameAlreadyExists 标签名已经存在
	TagNameAlreadyExists = errors.New("the tag name already exists")
	// TagNameNotExists 标签名不存在
	TagNameNotExists = errors.New("the tag name not exists")
	// TagIdNotExists 标签id不存在
	TagIdNotExists = errors.New("the tag id does not exist")
	// 找不到wechat openid
	WechatOpenIdNotExists = errors.New("the wechat openid not exists")
	// 手机号码不存在
	PhoneNumberNotExists = errors.New("the phone number not exists")
	// 邮箱不存在
	EmailNotExists = errors.New("the email not exists")
	// 评论主id不存在
	CommentMasterIdNotExists = errors.New("the comment master id does not exist")
	// 评论从id不存在
	CommentSlaveIdNotExists = errors.New("the comment slave id does not exist")
)
