package router

import (
	"LdapAdmin/config"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	request := r.Group("/request")
	{
		// system
		RegisterBaseRoutes(request)
		RegisterApiRoutes(request)
		RegisterMenuRoutes(request)

		// user
		RegisterUserRoutes(request)
	}
	r.Run(config.Conf.System.Host + config.Conf.System.Port)
}
