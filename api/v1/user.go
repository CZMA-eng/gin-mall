package v1

import (
	"gin_mall_tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context){
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err == nil{
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}