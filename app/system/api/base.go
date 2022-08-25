package api

import (
	"LdapAdmin/app/system/model"
	"LdapAdmin/app/system/service"
	"LdapAdmin/common/constant"
	model2 "LdapAdmin/common/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login
// @Summary      User Login
// @Description: user login by account and password
// @Tags         Login
// @Accept       application/json
// @Produce      application/json
// @Param        object body     model.LoginReq false "login data"
// @Success      200    {object} model.ResponseCommon
// @Failure      400    {object} model.ResponseErr
// @Failure      500    {object} model.ResponseErr
// @Route        "/request/base/login" [post]
func Login(c *gin.Context) {
	var req model.LoginReq
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, model2.ResponseErr{
			Code:   constant.ParamsError,
			ErrMsg: err.Error(),
		})
		return
	}
	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	resp, errCode, err := service.LoginService(&req, ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model2.ResponseErr{
			Code:   errCode,
			ErrMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model2.ResponseCommon{
		Code: 0,
		Msg:  "login success",
		Data: resp,
	})
}
