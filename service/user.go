package service

import (
	"context"
	"gin_mall_tmp/conf"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/e"
	util "gin_mall_tmp/pkg/utils"
	"gin_mall_tmp/serializer"
	"log"
	"mime/multipart"
	"strings"
	"time"

	"gopkg.in/mail.v2"
)

type ValidEmailService struct {

}

type SendEmailService struct {
	Email string `form:"email" json:"email"`
	Password string `json:"password" form:"password"`
	OperationType uint `json:"operation_type" form:"operation_type"`
	// 1. bind email 2. unbind email 3. change password
}

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key string `json:"key" form:"key"`  // key for encryption： frontend
}

func (service *ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint
	code := e.Success

	if token == "" {
		code = e.InvalidParams
	}else{
		claims, err := util.ParseEmailToken(token)
		log.Println(claims)
		if err != nil {
			code = e.ErrorAuthToken
		}else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		}else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}

	// get user info
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}
	if operationType == 1 {
		// binding email
		user.Email = email
	}else if operationType == 2 {
		// unbinding
		user.Email = ""
	}else if operationType == 3 {
		err = user.SetPassword(password)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg: e.GetMsg(code),
			}
		}
	}
	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		code = e.Error
			return serializer.Response{
				Status: code,
				Msg: e.GetMsg(code),
			}
	}

	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
		Data : serializer.BuildUser(user),
	}
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
	if ! user.CheckPassord(service.Password){
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
	user, _ = userDao.GetUserById(uId)
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

func (service *SendEmailService) Send(ctx context.Context, uId uint) serializer.Response{
	code := e.Success
	var address string
	var notice *model.Notice
	token , err := util.GenerateEmailToken(uId, service.OperationType, service.Email, service.Password)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg : e.GetMsg(code),
			Error: err.Error(),
		}
	}
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		code := e.Error
		return serializer.Response{
			Status: code,
			Msg : e.GetMsg(code),
			Error: err.Error(),
		}
	}
	address = conf.ValidEmail + token
	mailStr:=notice.Text +  "\n" + "Email"
	mailText := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "fanjingbo")
	m.SetBody("text/html", mailText)
	d:=mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err!=nil {
		code := e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg : e.GetMsg(code),
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
	}
}