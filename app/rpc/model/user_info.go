package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

// 用户信息创建的默认模型
var DefaultUserInfoTemplate = UserInfo{
	Uid:         0,
	Phone:       0,
	Email:       "",
	PhoneStatus: 0,
	EmailStatus: 0,
	CreateTime:  time.Time{},
	UpdateTime:  time.Time{},
	Sex:         0,
	Age:         0,
	Name:        "",
	NickName:    "",
	Addr:        "",
	Explain:     "",
	Level:       9,
}

// CreateUserInfo 根据用户基本模型创建用户信息
// tx为userBase创建的事务
func (ui *UserInfo) CreateUserInfo(userInfo UserInfo, tx *gorm.DB) error {
	return errors.WithStack(tx.Create(&userInfo).Error)
}

// 获取用户信息
func (ui *UserInfo) GetUserInfo(uid int32) (*UserInfo, error) {
	if DB.Where("uid = ?", uid).First(ui).RowsAffected == 0 {
		return nil, errors.WithStack(UserDoesNotExist)
	} else {
		return ui, nil
	}
}

// UpdateUserInfo 更新用户信息
// UserInfo模型的CreateTime会与查询的结果同步
// 会更新UserInfo模型的UpdateTime
func (ui *UserInfo) UpdateUserInfo(userInfo *UserInfo) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("uid = ?", userInfo.Uid).First(ui).Error; err != nil {
			return errors.WithStack(err)
		}
		// 创建时间不变化
		userInfo.CreateTime = ui.CreateTime
		// 更新更新时间
		userInfo.UpdateTime = time.Now()
		if err := tx.Model(ui).Where("uid = ?", ui.Uid).Updates(userInfo).Error; err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	return err
}
