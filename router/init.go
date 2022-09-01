package router

import (
	"LdapAdmin/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

// @title       LdapAdmin Server API
// @version     1.0
// @description This is a document for frontend
// @host        localhost:8080
// @BasePath    /request

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
	r.Run(fmt.Sprintf("%s:%s", config.Conf.System.Host, config.Conf.System.Port))
}
