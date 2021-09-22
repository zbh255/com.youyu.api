// 评论模型
package model

import (
	"com.youyu.api/app/rpc/proto_files"
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
	if DB.Model(&Article{}).Where("article_id = ?", cm.ArticleId).First(&Article{}).RowsAffected == 0 {
		return errors.WithStack(ArticleIdNotExists)
	}
	return DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(cm).Error
	})
}

// 查询主评论返回自定义错误和gorm原生错误
// 默认置顶的数据排在第一名
func (m *CommentMaster) GetMasterComments(articleId string, options *proto_files.OrderOptions) ([]*CommentMaster, error) {
	cms := make([]*CommentMaster, 0)
	if result := DB.Limit(int(options.PageNum)).Offset(int(options.PageNum*options.Page-options.PageNum)).
		Where("article_id = ?", articleId).Order("is_top desc").Order(options.Order+" "+options.Type).
		Find(&cms); result.RowsAffected == 0 {
		return nil, errors.WithStack(ArticleIdNotExists)
	} else {
		return cms, errors.WithStack(result.Error)
	}
}

// 删除主评论，返回自定义错误和gorm原生错误
// 防止uid越界，需要检查uid
func (m *CommentMaster) DeleteComment(mid int64, uid int32) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Where("comment_mid = ? AND uid = ?", mid, uid).Delete(&CommentMaster{CommentMid: mid, Uid: int(uid)}); result.RowsAffected == 0 {
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
		result := tx.Where("comment_mid = ? AND is_top = ?", mid, false).First(m)
		if result.RowsAffected == 0 {
			return errors.WithStack(CommentMasterIdNotExists)
		} else if result.Error != nil {
			return errors.WithStack(result.Error)
		}
		// 成功则修改置顶状态
		return errors.WithStack(tx.Model(m).Update("is_top", true).Error)
	})
}

// 删除评论置顶
// 返回自定义错误和gorm原生错误
func (m *CommentMaster) DeleteCommentTop(mid int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("comment_mid = ? AND is_top = ?", mid, true).First(m)
		if result.RowsAffected == 0 {
			return errors.WithStack(CommentMasterIdNotExists)
		} else if result.Error != nil {
			return errors.WithStack(result.Error)
		}
		// 成功则修改置顶状态
		return errors.WithStack(tx.Model(m).Update("is_top", false).Error)
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
		if tx.Model(&CommentMaster{}).Where("comment_mid = ?", cs.CommentMid).First(&CommentMaster{}).RowsAffected == 0 {
			return errors.WithStack(CommentMasterIdNotExists)
		}
		if tx.Model(&Article{}).Where("article_id = ?", cs.ArticleId).First(&Article{}).RowsAffected == 0 {
			return errors.WithStack(ArticleIdNotExists)
		}
		return errors.WithStack(tx.Model(s).Create(cs).Error)
	})
}

// 只能指定页数和一页包含多少数据
// 子评论默认按创建时间正序排序
func (s *CommentSlave) GetSlaveComments(cms []*CommentMaster, page int, prePage int) ([]*proto_files.CommentMasterShow, error) {
	if cms == nil || len(cms) == 0 {
		return nil, errors.WithStack(CommentMasterIdNotExists)
	}
	css := make([]*proto_files.CommentMasterShow, 0, len(cms))
	var result *gorm.DB
	for i := 0; i < len(cms); i++ {
		tmp := make([]*CommentSlave, 0)
		result = DB.Model(s).Limit(prePage).Offset(page*prePage-prePage).Order("create_time asc").
			Where("comment_mid = ?", cms[i].CommentMid).Find(&tmp)
		tmpCSS := &proto_files.CommentMasterShow{
			CommentMid: cms[i].CommentMid,
			Type:       proto_files.CommentType(cms[i].Type),
			Text:       cms[i].Text,
			Uid:        int32(cms[i].Uid),
			ArticleId:  cms[i].ArticleId,
			Fabulous:   cms[i].Fabulous,
			CreateTime: cms[i].CreateTime.String(),
			IsTop:      cms[i].IsTop,
		}
		if result.RowsAffected != 0 {
			tmpSlave := make([]*proto_files.CommentSlave, len(tmp))
			// 初始化指针数据
			for k := range tmpSlave {
				tmpSlave[k] = &proto_files.CommentSlave{}
			}
			for k, v := range tmp {
				tmpSlave[k].CommentSid = v.CommentSid
				tmpSlave[k].CommentMid = v.CommentMid
				tmpSlave[k].Text = v.Text
				tmpSlave[k].Uid = int32(v.Uid)
				tmpSlave[k].Fabulous = v.Fabulous
				tmpSlave[k].ArticleId = v.ArticleId
				tmpSlave[k].Type = proto_files.CommentType(v.Type)
				tmpSlave[k].ReplyId = v.ReplyId
				tmpSlave[k].CreateTime = v.CreateTime.String()
			}
			// 有子评论则赋值
			tmpCSS.SlaveComment = tmpSlave
		} else {
			tmpCSS.SlaveComment = nil
		}
		css = append(css, tmpCSS)
	}

	return css, errors.WithStack(result.Error)
}

func (s *CommentSlave) DeleteComment(mid, sid int64, uid int32) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Where("comment_mid = ? AND comment_sid = ? AND uid = ?", mid, sid, uid).
			Delete(&CommentSlave{CommentSid: sid}); result.RowsAffected == 0 {
			return errors.WithStack(CommentSlaveIdNotExists)
		} else {
			return errors.WithStack(result.Error)
		}
	})
}
