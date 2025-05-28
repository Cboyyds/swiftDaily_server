package service

import (
	"swiftDaily_myself/global"
	"swiftDaily_myself/model/database"
)

// 注册时创建用户

func (b *BaseService) EmailRegister(user database.User) error {
	user.RoleID = 1
	user.Avatar = ""
	user.Signature = "这个人很懒，什么也没有留下"
	if err := global.DB.Where("email = ?", user.Email).Find(&database.User{}).Error; err != nil {
		return err
	}
	if err := global.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
