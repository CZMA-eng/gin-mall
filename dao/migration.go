package dao

import (
	"fmt"
	"gin_mall_tmp/model"
)

func migration(){
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.Address{}, 
		&model.Admin{}, 
		&model.Carousel{}, 
		&model.Cart{}, 
		&model.Category{}, 
		&model.Favorite{}, 
		&model.Notice{}, 
		&model.Order{}, 
		&model.Product{}, 
		&model.ProductImg{}, 
		&model.User{},
	)
	if err != nil {
		fmt.Println("fail to migrate: ", err)
	}

	return
}