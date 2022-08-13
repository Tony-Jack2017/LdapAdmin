package router

import (
	"LdapAdmin/app/user/api"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(v *gin.RouterGroup) {
	user := v.Group("user")
	{
		user.POST("/add", api.AddUser)
		user.POST("/delete", api.DeleteUser)
		user.GET("/list", api.GetUsers)
	}
}
