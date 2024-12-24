package dao

import (
	"context"
	"gin_mall_tmp/model"

	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// 可以复用 db 的链接
func NewNoticeDaoByDB(db *gorm.DB) *UserDao{
	return &UserDao{db}
}

func (dao *UserDao) GetNoticeById(id uint)(notice *model.Notice, err error){
	err = dao.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
}