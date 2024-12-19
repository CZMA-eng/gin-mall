package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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

const (
	PasswordCost = 12 // 加密难度
	Active string = "active"
)

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest=string(bytes)
	return err
}