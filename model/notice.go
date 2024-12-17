package model

import "gorm.io/gorm"

type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"` // 多用于大文本块
}