package service

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/pkg/e"
	util "gin_mall_tmp/pkg/utils"
	"gin_mall_tmp/serializer"
)

type CarouselService struct {

}

func (service *CarouselService) List(ctx context.Context) serializer.Response {
	 carouselDao := dao.NewCarouselDao(ctx)
	 carousels, err := carouselDao.ListCarousel()
	 if err != nil {
		util.LogrusObj.Info("err", err)
		code := e.Error
		return serializer.Response{
			Status: code,
			Msg : e.GetMsg(code),
			Error: err.Error(),
		}
	}
	
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}