package model

import "gorm.io/gorm"

type Carousel struct {
	gorm.Model
	ImgPath string
	ProductId uint `gor:"not null"`
}