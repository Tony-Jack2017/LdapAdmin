package middleware

import (
	"LdapAdmin/common/constant"
	"LdapAdmin/common/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthorityMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, exist := context.Get("role")
		if !exist {
			context.JSON(http.StatusInternalServerError, model.ResponseErr{
				Code:   constant.ServerError,
				ErrMsg: "the server is happening some errors",
			})
		}
	}
}
