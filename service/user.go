package service

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/e"
	util "gin_mall_tmp/pkg/utils"
	"gin_mall_tmp/serializer"
	"mime/multipart"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key string `json:"key" form:"key"`  // key for encryption： frontend
}

func (service *UserService) Register (ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success
	if service.Key==""||len(service.Key)!=16{
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: "need longer key",
		}
	}
	util.Encrypt.SetKey(service.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}

	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}

	user = model.User{
		UserName: service.UserName,
		NickName: service.NickName,
		Status: model.Active,
		Avatar: "avatar.JPG",
		Money: util.Encrypt.AesEncoding("10000"),
	}
	// password encryption
	if err = user.SetPassword(service.Password);err!=nil{
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}

	// create user
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error
		
	}

	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
	}
}

func (service *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	// get user
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if !exist || err != nil {
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}
	// compare password
	if user.CheckPassord(service.Password)==false{
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Data: "invalid password, login again",
		}
	}
	// http stateless -> generate token for client
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err!=nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Data: "authentication error",
		}
	}

	return serializer.Response{
		Status: code,
		Data: serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg: e.GetMsg(code),
	}
}

func (service *UserService) Update(ctx context.Context, uId uint) serializer.Response {
	var user *model.User
	var err error
	code := e.Success

	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	// modify nickname
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg : e.GetMsg(code),
				Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
		Data: serializer.BuildUser(user),
	}
}

func (service *UserService) Post(ctx context.Context, uId uint, 
					file multipart.File, fileSize int64)serializer.Response{
	code := e.Success
	var user *model.User
	var err error
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		code = e.Error
			return serializer.Response{
				Status: code,
				Msg : e.GetMsg(code),
				Error: err.Error(),
			}
	}
	// store file in local 
	path, err := UploadAvartarToLocalStatic(file, uId, user.UserName)
	if err != nil {
		code = e.ErrorUploadFail
		return serializer.Response{
			Status: code,
			Msg : e.GetMsg(code),
			Error: err.Error(),
		}
	}
	user.Avatar = path
	err=userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg : e.GetMsg(code),
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
		Data: serializer.BuildUser(user),
	}
}