package v1

import (
	util "gin_mall_tmp/pkg/utils"
	"gin_mall_tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context){
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err == nil{
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("register error", err)
	}
}

func UserLogin(c *gin.Context){
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil{
		res := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("login error", err)
	}
}

func UserUpdate(c *gin.Context){
	var userUpdate service.UserService
	claims,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdate); err == nil{
		res := userUpdate.Update(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user update error", err)
	}
}

func UploadAvatar(c *gin.Context){
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
        return
    }
	fileSize := fileHeader.Size
	var UploadAvatar service.UserService
	claims ,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&UploadAvatar); err == nil{
		res := UploadAvatar.Post(c.Request.Context(), claims.ID, file, fileSize)
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
		util.LogrusObj.Infoln("avatar upload error", err)
	}
}

func SendEmail(c *gin.Context){
	var sendEmail service.SendEmailService
	claims ,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&sendEmail); err == nil{
		res := sendEmail.Send(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
		util.LogrusObj.Infoln("email send error", err)
	}
}

func ValidEmail(c *gin.Context){
	var ValidEmail service.ValidEmailService
	_ ,err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	if err := c.ShouldBind(&ValidEmail); err == nil{
		res := ValidEmail.Valid(c.Request.Context(), c.GetHeader("Authorization"))
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
		util.LogrusObj.Infoln("valid email error", err)
	}
}

func ShowMoney(c *gin.Context){
	var ShowMoney service.ShowMoneyService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&ShowMoney); err == nil{
		res := ShowMoney.Show(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}else{
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
		util.LogrusObj.Infoln("show money error", err)
	}
}