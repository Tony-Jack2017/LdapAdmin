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
		group.POST("/add", api.AddApi)
		group.POST("/delete", api.DeleteApi)
		group.GET("/list", api.GetApiList)
		group.POST("/modify", api.ModifyApi)
	}
	RegisterApiGroupRoutes(group)
}

func RegisterApiGroupRoutes(v *gin.RouterGroup) {
	group := v.Group("group")
	{
		group.POST("/add", api.AddApiGroup)
		group.POST("/delete", api.DeleteApiGroup)
		group.GET("/list", api.GetApiGroupList)
		group.POST("/modify", api.ModifyApiGroup)
	}
}

func RegisterMenuRoutes(v *gin.RouterGroup) {
	group := v.Group("menu")
	{
		group.GET("/list", api.GetMenuList)
		group.POST("/add", api.AddMenu)
		group.POST("/delete", api.DeleteMenu)
		group.POST("/modify", api.ModifyMenu)
	}
}
