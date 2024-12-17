package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Email string
	PasswordDigest string  // hashcode of pwd
	NickName string
	Status string // ban or not
	Avatar string
	Money string // encrypted 
}