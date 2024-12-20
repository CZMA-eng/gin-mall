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

func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Find(&user).Error
	if user == nil || err==gorm.ErrRecordNotFound{
		return nil, false, err
	}
	return user, true, nil
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

func (dao *UserDao) GetUserById(id uint)(user *model.User, err error){
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

func (dao *UserDao) UpdateUserById(uid uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", uid).Updates(&user).Error
}