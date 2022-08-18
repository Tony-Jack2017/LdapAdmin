package router

import (
	"LdapAdmin/config"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	request := r.Group("/request")
	{
		RegisterBaseRoutes(request)
		RegisterUserRoutes(request)
	}
	r.Run(config.Conf.System.Host + config.Conf.System.Port)
}
