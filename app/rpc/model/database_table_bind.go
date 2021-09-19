package model

import (
	"gorm.io/gorm"
	"time"
)

// 文章点赞表
type ArticleStatistics struct {
	Id         string `gorm:"primaryKey;column:article_id"`
	Fabulous   int32  `gorm:"column:article_fabulous"`
	Hot        int32  `gorm:"column:article_hot"`
	CommentNum int32  `gorm:"column:article_comment_num"`
}

type SelectOptions struct {
	Type    string
	Order   string
	Page    int32
	PageNum int32
}

// 文章表
type Article struct {
	Id         string    `gorm:"primaryKey;column:article_id"`
	Abstract   string    `gorm:"column:article_abstract"`
	Content    string    `gorm:"column:article_content"`
	Title      string    `gorm:"column:article_title"`
	Tag        string    `gorm:"column:article_tag"`
	Uid        int64     `gorm:"column:uid"`
	CreateTime time.Time `gorm:"column:article_create_time"`
	UpdateTime time.Time `gorm:"column:article_update_time"`
}

// 广告表
type Advertisement struct {
	Id     int32  `gorm:"primaryKey;column:advertisement_id"`
	Type   int32  `gorm:"column:advertisement_type"`
	Link   string `gorm:"column:advertisement_link"`
	Weight int32  `gorm:"column:advertisement_weight"`
	Body   string `gorm:"column:advertisement_body"`
	Owner  string `gorm:"column:advertisement_owner"`
}

// 首页展示数据连接表
type ArticleDataLinkTable struct {
	Id         string    `gorm:"primaryKey;column:article_id"`
	Abstract   string    `gorm:"column:article_abstract"`
	Title      string    `gorm:"column:article_title"`
	Tag        string    `gorm:"column:article_tag"`
	Uid        int64     `gorm:"column:uid"`
	CreateTime time.Time `gorm:"column:article_create_time"`
	// article_statistics
	UpdateTime time.Time `gorm:"column:article_update_time"`
	Fabulous   int32     `gorm:"column:article_fabulous"`
	Hot        int32     `gorm:"column:article_hot"`
	CommentNum int32     `gorm:"column:article_comment_num"`
}

// 标签表
type Tags struct {
	Tid  int32  `gorm:"primaryKey;column:tid"`
	Text string `gorm:"column:text"`
}

// 用户基本表
type UserBase struct {
	Uid          int32  `gorm:"primaryKey;column:uid"`
	UserPassword string `gorm:"column:password"`
	Salt         string `gorm:"column:salt"`
	Name         string `gorm:"column:name"`
}

// 用户信息表
type UserInfo struct {
	Uid                int32     `gorm:"column:uid"`
	Name               string    `gorm:"primaryKey;column:name"`
	Phone              int64     `gorm:"column:phone"`
	Email              string    `gorm:"column:email"`
	PhoneStatus        int       `gorm:"column:phone_status"`
	EmailStatus        int       `gorm:"column:email_status"`
	CreateTime         time.Time `gorm:"column:create_time"`
	UpdateTime         time.Time `gorm:"column:update_time"`
	Sex                int       `gorm:"column:sex"`
	Age                int       `gorm:"column:age"`
	NickName           string    `gorm:"column:nick_name"`
	Explain            string    `gorm:"column:explain"`
	Level              int32     `gorm:"column:level"`
	WechatOpenId       string    `gorm:"column:wechat_openid"`
	WechatOpenIdStatus int32     `gorm:"column:wechat_openid_status"`
	HeadPortrait       string    `gorm:"column:head_portrait"`
	Country            string    `gorm:"column:country"`
	Province           string    `gorm:"column:province"`
	City               string    `gorm:"column:city"`
	DetailAddr         string    `gorm:"column:detail_addr"`
	Language           string    `gorm:"column:language"`
}

// 评论主表
type CommentMaster struct {
	CommentMid int64     `gorm:"primaryKey;column:comment_mid"`
	Type       int       `gorm:"column:type"`
	Text       string    `gorm:"column:text"`
	Uid        int       `gorm:"column:uid"`
	ArticleId  string    `gorm:"column:article_id"`
	Fabulous   int32     `gorm:"column:fabulous"`
	CreateTime time.Time `gorm:"column:create_time"`
	IsTop      bool      `gorm:"column:is_top"`
}

// 评论从表
type CommentSlave struct {
	CommentSid int64     `gorm:"primaryKey;column:comment_sid"`
	CommentMid int64     `gorm:"column:comment_mid"`
	Type       int       `gorm:"column:type"`
	Text       string    `gorm:"column:text"`
	Uid        int       `gorm:"column:uid"`
	ReplyId    int64     `gorm:"column:reply_id"`
	ArticleId  string    `gorm:"column:article_id"`
	Fabulous   int32     `gorm:"column:fabulous"`
	CreateTime time.Time `gorm:"column:create_time"`
}

// DB 数据库接口
var DB *gorm.DB

