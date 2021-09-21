// 评论模型
package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

// 添加主评论，返回自定义错误和gorm原生错误
func (m *CommentMaster) AddComment(cm *CommentMaster) error {
	// 初始化数据
	cm.IsTop = false
	cm.CommentMid = 0
	cm.Fabulous = 0
	cm.CreateTime = time.Now()
	// crud
	if DB.Model(&Article{}).Where("article_id = ?",cm.ArticleId).First(&Article{}).RowsAffected == 0 {
		return errors.WithStack(ArticleIdNotExists)
	}
	return DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(cm).Error
	})
}

// 查询主评论返回自定义错误和gorm原生错误
func (m *CommentMaster) GetMasterComments(articleId string) ([]*CommentMaster,error) {
	cms := make([]*CommentMaster,0)
	if result := DB.Where("article_id = ?", articleId).Find(&cms);result.RowsAffected == 0 {
		return nil,errors.WithStack(ArticleIdNotExists)
	} else {
		return cms,errors.WithStack(result.Error)
	}
}

// 删除主评论，返回自定义错误和gorm原生错误
// 防止uid越界，需要检查uid
func (m *CommentMaster) DeleteComment(mid int64,uid int32) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Where("comment_mid = ? AND uid = ?",mid,uid).Delete(&CommentMaster{CommentMid: mid,Uid: int(uid)}); result.RowsAffected == 0 {
			return errors.WithStack(CommentMasterIdNotExists)
		} else {
			return errors.WithStack(result.Error)
		}
	})
}

// 添加评论置顶
// 返回自定义错误和gorm原生错误
func (m *CommentMaster) AddCommentTop(mid int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("comment_mid = ? AND is_top = ?",mid,false).First(m)
		if result.RowsAffected == 0 {
			return errors.WithStack(CommentMasterIdNotExists)
		}  else if result.Error != nil {
			return errors.WithStack(result.Error)
		}
		// 成功则修改置顶状态
		return errors.WithStack(tx.Model(m).Update("is_top",true).Error)
	})
}

// 删除评论置顶
// 返回自定义错误和gorm原生错误
func (m *CommentMaster) DeleteCommentTop(mid int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("comment_mid = ? AND is_top = ?",mid,true).First(m)
		if result.RowsAffected == 0 {
			return errors.WithStack(CommentMasterIdNotExists)
		}  else if result.Error != nil {
			return errors.WithStack(result.Error)
		}
		// 成功则修改置顶状态
		return errors.WithStack(tx.Model(m).Update("is_top",false).Error)
	})
}

/*
	以下为评论从表操作模型
*/

// 添加子评论
// 返回自定义错误和gorm原生错误
func (s *CommentSlave) AddComment(cs *CommentSlave) error {
	// 初始化一些数据
	cs.CommentSid = 0
	cs.CreateTime = time.Now()
	cs.Fabulous = 0
	return DB.Transaction(func(tx *gorm.DB) error {
		if tx.Model(&CommentMaster{}).Where("comment_mid = ?",cs.CommentMid).First(&CommentMaster{}).RowsAffected == 0 {
			return errors.WithStack(CommentMasterIdNotExists)
		}
		if tx.Model(&Article{}).Where("article_id = ?",cs.ArticleId).First(&Article{}).RowsAffected == 0 {
			return errors.WithStack(ArticleIdNotExists)
		}
		return errors.WithStack(tx.Model(s).Create(cs).Error)
	})
}

func (s *CommentSlave) GetSlaveComments(cms []*CommentMaster) (map[*CommentMaster][]*CommentSlave,error) {
	if cms == nil || len(cms) == 0 {
		return nil,errors.WithStack(CommentMasterIdNotExists)
	}
	// 构建哨兵条件
	CmsTmp := make(map[int64]*CommentMaster)
	for _,v := range cms {
		CmsTmp[v.CommentMid] = v
	}
	midList := make([]int64,len(cms))
	for i, c := range cms {
		midList[i] = c.CommentMid
	}
	css := make(map[*CommentMaster][]*CommentSlave)
	tmp := make([]*CommentSlave,0)

	var result *gorm.DB
	for _,v := range midList {
		if result = DB.Model(s).Where("comment_mid = ?",v).Find(&tmp); result.RowsAffected != 0 {
			css[CmsTmp[v]] = tmp
		} else {
			css[CmsTmp[v]] = nil
		}
	}

	return css,errors.WithStack(result.Error)
}

func (s *CommentSlave) DeleteComment(mid,sid int64,uid int32) error  {
	return DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Where("comment_mid = ? AND comment_sid = ? AND uid = ?",mid,sid,uid).
			Delete(&CommentSlave{CommentSid: sid}); result.RowsAffected == 0 {
			return errors.WithStack(CommentSlaveIdNotExists)
		} else {
			return errors.WithStack(result.Error)
		}
	})
}