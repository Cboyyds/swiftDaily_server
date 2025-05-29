package service

import (
	"errors"
	"swiftDaily_myself/global"
	"swiftDaily_myself/model/database"
	"swiftDaily_myself/utils"
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

func (b *BaseService) EmailLogin(user *database.User) (database.User, error) {
	var u database.User
	if err := global.DB.Where("email = ?", user.Email).First(&u).Error; err != nil {
		return u, err
	}
	if err := global.DB.Model(&u).Update("status", 1).Error; err != nil {
		return u, err
	}
	if !utils.BcryptCheck(user.Password, u.Password) {
		return u, errors.New("密码错误")
	}
	return u, nil
}
