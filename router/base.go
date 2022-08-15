package router

import (
	"LdapAdmin/app/base/api"
	"github.com/gin-gonic/gin"
)

func RegisterBaseRoutes(v *gin.RouterGroup) {
	base := v.Group("base")
	{
		base.POST("/login", api.Login)
	}
}

func RegisterApiRoutes(v *gin.RouterGroup) {
	_ = v.Group("api")
}

func RegisterMenuRoutes(v *gin.RouterGroup) {
	_ = v.Group("api")
}
