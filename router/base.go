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
