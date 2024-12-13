package repository

import (
	"gin-api/global"
	"gin-api/model"
)

type UserRepository struct{}

var UserRepo = new(UserRepository)

func (dao *UserRepository) GetUserByMobile(mobile string) (user *model.User, err error) {
	err = global.DB.First(&user, "phone = ?", mobile).Error
	return
}
