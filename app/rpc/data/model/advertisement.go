// Package model Table Adevertisement绑定的操作方法
package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// 操作数据时有错误则返回gorm 的原始错误
func (self *Advertisement) AddAdvertisement(advertisement *Advertisement) error {
	result := DB.Create(advertisement)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

// GetAdvertisement 获得广告数据
// 广告id不存在时会返回自定义错误
// 操作数据时有错误则返回gorm 的原始错误
func (self *Advertisement) GetAdvertisement(id int32) (*Advertisement, error) {
	result := DB.Where("advertisement_id = ?", id).First(self)
	// 找到结果
	if result.RowsAffected == 0 {
		return nil, errors.WithStack(AdvertisementIdNotExists)
	} else if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	} else {
		return self,nil
	}
}

// 自定义查询广告列表
// 操作数据时有错误则返回gorm 的原始错误
func (self *Advertisement) GetAdvertisements(op *SelectOptions) ([]*Advertisement, error) {
	// 分页选项不能为0
	if op.Page == 0 || op.PageNum == 0 {
		return nil, errors.New("page and pageNum cannot be zero")
	}
	advertisements := make([]*Advertisement, 0)
	switch op.Type {
	case "asc":
		break
	default:
		op.Type = "desc"
	}
	err := DB.Model(self).Limit(int(op.PageNum)).Offset(int(op.PageNum*op.Page - op.PageNum)).Order("advertisement_weight" + " " + op.Type).Find(&advertisements).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return advertisements, nil
}

// 删除指定的广告
// 操作数据时有错误则返回gorm 的原始错误
func (self *Advertisement) DelAdvertisement(id int32) error {
	//"mysql server connection failed"
	err := DB.Transaction(func(tx *gorm.DB) error {
		// 根据主键删除数据
		self.Id = id
		if err := tx.Delete(self).Error; err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	return err
}

// 根据id更新广告
// 操作数据时有错误则返回gorm 的原始错误
func (self *Advertisement) SetAdvertisement(advertisement *Advertisement) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(self).Where("advertisement_id = ?", advertisement.Id).Updates(advertisement).Error; err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	if err != nil {
		return err
	} else {
		return nil
	}
}
