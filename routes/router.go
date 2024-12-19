package routes

import (
	"gin_mall_tmp/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	api "gin_mall_tmp/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func (c *gin.Context)  {
			c.JSON(200, "success")
		})
		v1.POST("user/register", api.UserRegister)
	}
	return r
}