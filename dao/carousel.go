package dao

import (
	"context"
	"gin_mall_tmp/model"

	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

// 可以复用 db 的链接
func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao{
	return &CarouselDao{db}
}

func (dao *CarouselDao) ListCarousel()(carousel []model.Carousel, err error){
	err = dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return
}