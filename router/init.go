package router

import (
	"LdapAdmin/config"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	api := r.Group("/api")
	RegisterBaseRoutes(api)
	RegisterUserRoutes(api)
	r.Run(config.Conf.System.Host + config.Conf.System.Port)
}
