package serializer

import (
	"gin_mall_tmp/conf"
	"gin_mall_tmp/model"
)

type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"username"`
	NickName  string `json:"nickname"`
	Type      int    `json:"type"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreateAt  int64  `json:"create_at"`
}

func BuildUser(user *model.User) *User {
	return &User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   conf.Host + conf.HttpPort+ conf.AvatarPath+user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}