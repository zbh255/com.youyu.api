package model

import (
	"com.youyu.api/common/database"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

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

type Advertisement struct {
	Id     int32  `gorm:"primaryKey;column:advertisement_id"`
	Type   int32  `gorm:"column:advertisement_type"`
	Link   string `gorm:"column:advertisement_link"`
	Weight int32  `gorm:"column:advertisement_weight"`
	Body   string `gorm:"column:advertisement_body"`
	Owner  string `gorm:"column:advertisement_owner"`
}

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

// 数据库接口
var DB = database.DataBase(&database.Mysql{})

func (self *ArticleStatistics) AddFabulous(id string) error {
	db, err := DB.GetConnect()
	if err != nil {
		return errors.Wrap(err, "")
	}
	if db == nil {
		return errors.New("mysql server connection failed")
	}
	result := db.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return errors.New("the query record is zero")
	}
	result = db.Model(self).Where("article_id = ?", id).Update("article_fabulous", self.Fabulous+1)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

// 点赞-1
func (self *ArticleStatistics) ReduceFabulous(id string) error {
	db, err := DB.GetConnect()
	if err != nil {
		return errors.Wrap(err, "")
	}
	if db == nil {
		return errors.New("mysql server connection failed")
	}
	result := db.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return errors.New("the query record is zero")
	}
	result = db.Model(self).Where("article_id = ?", id).Update("article_fabulous", self.Fabulous-1)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

// 获得存储文章点赞等数据的表
func (self *ArticleStatistics) GetArticleStatistics(id string) (*ArticleStatistics, error) {
	db, err := DB.GetConnect()
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, errors.New("mysql server connection failed")
	}
	result := db.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return nil, errors.New("the query record is zero")
	} else {
		return self, nil
	}
}

// 错误类型为数据库连接错误和查询错误
func (self *ArticleStatistics) AddHot(id string) error {
	db, err := DB.GetConnect()
	if err != nil {
		return errors.Wrap(err, "")
	}
	result := db.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return errors.New("the query record is zero")
	}
	result = db.Model(self).Where("article_id = ?", id).Update("article_hot", self.Hot+1)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (self *ArticleStatistics) AddCommentNum(id string) error {
	db, err := DB.GetConnect()
	if err != nil {
		return errors.Wrap(err, "")
	}
	result := db.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return errors.New("the query record is zero")
	}
	result = db.Model(self).Where("article_id = ?", id).Update("article_comment_num", self.CommentNum+1)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (self *ArticleStatistics) AddSelf(statistics *ArticleStatistics) error {
	db, err := DB.GetConnect()
	if err != nil {
		return errors.Wrap(err, "")
	}
	result := db.Create(statistics)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (self *Article) AddArticle(article *Article) error {
	db, err := DB.GetConnect()
	if err != nil {
		return errors.Wrap(err, "")
	}
	// 查找文章是否存在
	result := db.Where("article_id = ?", article.Id).First(self)
	if result.RowsAffected > 0 {
		return errors.New("article already exists")
	}
	result = db.Create(article)
	if result.Error != nil {
		return result.Error
	}
	result = db.Create(&ArticleStatistics{
		Id:         article.Id,
		Fabulous:   0,
		Hot:        0,
		CommentNum: 0,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (self *Article) DelArticle(id string) error {
	db, err := DB.GetConnect()
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		// 根据主键删除数据
		// 因为设置的外键，所以存储点赞等信息的也会自动删除
		self.Id = id
		if err := tx.Delete(self).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// 获取文章的数据
func (self *Article) GetArticle(id string) (*Article, error) {
	db, err := DB.GetConnect()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	result := db.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return nil, errors.New("the query record is zero")
	} else {
		return self, nil
	}
}

func (self *Article) SetArticle(article *Article) error {
	db, err := DB.GetConnect()
	if err != nil {
		return errors.Wrap(err, "")
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("article_id = ?", article.Id).First(self).Error; err != nil {
			return err
		}
		// 创建时间不变化
		article.CreateTime = self.CreateTime
		if err := tx.Model(self).Where("article_id = ?", self.Id).Updates(article).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	} else {
		return nil
	}
}

// 返回连接查询的列表
func (self *Article) GetArticles(op *SelectOptions) ([]*ArticleDataLinkTable, error) {
	// 分页选项不能为0
	if op.Page == 0 || op.PageNum == 0 {
		return nil, errors.New("page and pageNum cannot be zero")
	}
	articleLinkS := make([]*ArticleDataLinkTable, 0)
	db, err := DB.GetConnect()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	// 排序规则,默认为热度
	order := ""
	switch op.Order {
	case "hot":
		order = "article_hot"
	case "fabulous":
		order = "article_fabulous"
	case "comment_num":
		order = "article_comment_num"
	default:
		order = "article_hot"
	}
	switch op.Type {
	case "asc":
		break
	default:
		op.Type = "desc"
	}
	articleLinkField := "article.article_id,article.article_abstract,article.article_content,article.article_tag,article.uid,article.article_create_time,article.article_update_time"
	articleStatisticsLinkField := "article_statistics.article_fabulous,article_statistics.article_hot,article_statistics.article_comment_num"
	if err := db.Model(self).Limit(int(op.PageNum)).Offset(int(op.PageNum*op.Page - op.PageNum)).Select(articleLinkField + "," + articleStatisticsLinkField).Joins("join article_statistics on article.article_id = article_statistics.article_id").Order(order + " " + op.Type).Scan(&articleLinkS).Error; err != nil {
		return nil, err
	}
	return articleLinkS, nil
}

func (self *Advertisement) AddAdvertisement(advertisement *Advertisement) error {
	db, err := DB.GetConnect()
	if err != nil {
		return errors.Wrap(err, "")
	}
	result := db.Create(advertisement)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

// 获得广告数据
func (self *Advertisement) GetAdvertisement(id int32) (*Advertisement, error) {
	db, err := DB.GetConnect()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	result := db.Where("advertisement_id = ?", id).First(self)
	// 找到结果
	if result.RowsAffected > 0 {
		return self, nil
	} else {
		return nil, errors.New("the query record is zero")
	}
}

// 自定义查询广告列表
func (self *Advertisement) GetAdvertisements(op *SelectOptions) ([]*Advertisement, error) {
	// 分页选项不能为0
	if op.Page == 0 || op.PageNum == 0 {
		return nil, errors.New("page and pageNum cannot be zero")
	}
	advertisements := make([]*Advertisement, 0)
	db, err := DB.GetConnect()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	switch op.Type {
	case "asc":
		break
	default:
		op.Type = "desc"
	}
	err = db.Model(self).Limit(int(op.PageNum)).Offset(int(op.PageNum*op.Page - op.PageNum)).Order("advertisement_weight" + " " + op.Type).Find(&advertisements).Error
	if err != nil {
		return nil, err
	}
	return advertisements, nil
}

func (self *Advertisement) DelAdvertisement(id int32) error {
	//"mysql server connection failed"
	db, err := DB.GetConnect()
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		// 根据主键删除数据
		self.Id = id
		if err := tx.Delete(self).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (self *Advertisement) SetAdvertisement(advertisement *Advertisement) error {
	db, err := DB.GetConnect()
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(self).Where("advertisement_id = ?", advertisement.Id).Updates(advertisement).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	} else {
		return nil
	}
}
