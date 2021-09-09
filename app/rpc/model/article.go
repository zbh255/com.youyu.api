// Package model Table Article绑定的操作方法
package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// AddArticle 获取文章的数据
// pkg/errors处理错误
// article_id 已存在的时候返回自包装错误
// 操作数据时有错误则返回gorm 的原始错误
func (self *Article) AddArticle(article *Article) (*Article, error) {
	// 查找文章是否存在
	result := DB.Where("article_id = ?", article.Id).First(self)
	if result.RowsAffected > 0 {
		return nil, errors.WithStack(ArticleIdAlreadyExists)
	}
	result = DB.Create(article)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	result = DB.Create(&ArticleStatistics{
		Id:         article.Id,
		Fabulous:   0,
		Hot:        0,
		CommentNum: 0,
	})
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	return article, nil
}

// DelArticle 操作数据时有错误则返回gorm 的原始错误
func (self *Article) DelArticle(id string) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		// 根据主键删除数据
		// 因为设置的外键，所以存储点赞等信息的也会自动删除
		self.Id = id
		if err := tx.Delete(self).Error; err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	return err
}

// GetArticle 获取文章的数据
// pkg/errors处理错误
// article_id 不存在的时候返回自包装错误
// 操作数据时有错误则返回gorm 的原始错误
func (self *Article) GetArticle(id string) (*Article, error) {
	result := DB.Where("article_id = ?", id).First(self)
	if result.RowsAffected == 0 {
		return nil,errors.WithStack(ArticleIdNotExists)
	} else if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	} else {
		return self, nil
	}
}

// SetArticle article_id 不存在的时候返回自包装错误
// 操作数据时有错误则返回gorm 的原始错误
// TODO:增加uid的判断
func (self *Article) SetArticle(article *Article) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Where("article_id = ?", article.Id).First(self); result.RowsAffected == 0 {
			return errors.WithStack(ArticleIdNotExists)
		} else if result.Error != nil {
			return errors.WithStack(result.Error)
		}
		// 创建时间不变化
		article.CreateTime = self.CreateTime
		if err := tx.Model(self).Where("article_id = ?", self.Id).Updates(article).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	} else {
		return nil
	}
}

// GetArticles 返回连接查询的列表
// 操作数据时有错误则返回gorm 的原始错误
func (self *Article) GetArticles(op *SelectOptions) ([]*ArticleDataLinkTable, error) {
	// 分页选项不能为0
	if op.Page == 0 || op.PageNum == 0 {
		return nil, errors.New("page and pageNum cannot be zero")
	}
	articleLinkS := make([]*ArticleDataLinkTable, 0)
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
	if err := DB.Model(self).Limit(int(op.PageNum)).Offset(int(op.PageNum*op.Page - op.PageNum)).Select(articleLinkField + "," + articleStatisticsLinkField).Joins("join article_statistics on article.article_id = article_statistics.article_id").Order(order + " " + op.Type).Scan(&articleLinkS).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return articleLinkS, nil
}
