package router

import (
	"LdapAdmin/app/system/api"
	"github.com/gin-gonic/gin"
)

func RegisterBaseRoutes(v *gin.RouterGroup) {
	group := v.Group("base")
	{
		group.POST("/login", api.Login)
	}
}

func RegisterApiRoutes(v *gin.RouterGroup) {
	group := v.Group("api")
	{
		group.GET("/list", api.GetApiList)
	}
}

func RegisterMenuRoutes(v *gin.RouterGroup) {
	group := v.Group("menu")
	{
		group.GET("/list", api.GetMenuListApi)
	}
}
