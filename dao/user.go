package dao

import (
	"context"
	"gin_mall_tmp/model"

	"gorm.io/gorm"
)


type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// 可以复用 db 的链接
func NewUserDaoByDB(db *gorm.DB) *UserDao{
	return &UserDao{db}
}

func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exits bool, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Find(&user).Error
	if user == nil || err==gorm.ErrRecordNotFound{
		return nil, true, err
	}
	return user, false, nil
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}