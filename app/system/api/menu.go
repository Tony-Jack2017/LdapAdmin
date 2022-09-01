package api

import (
	"LdapAdmin/app/system/model"
	"LdapAdmin/app/system/service"
	"LdapAdmin/common/constant"
	model2 "LdapAdmin/common/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddMenu(c *gin.Context) {
	var req model.AddMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model2.ResponseErr{
			Code:   constant.ParamsError,
			ErrMsg: err.Error(),
		})
		return
	}
	_, err := service.AddMenuService(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model2.ResponseErr{
			Code:   constant.SqlError,
			ErrMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model2.Response{
		Code: 0,
		Msg:  "Add menu success",
	})

}

func DeleteMenu(c *gin.Context) {
	var req model.DeleteMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model2.ResponseErr{
			Code:   constant.ParamsError,
			ErrMsg: err.Error(),
		})
		return
	}
	err := service.DeleteMenuService(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model2.ResponseErr{
			Code:   constant.SqlError,
			ErrMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model2.Response{
		Code: 0,
		Msg:  "Delete menus success",
	})
}

func GetMenuList(c *gin.Context) {
	var req model.GetMenuListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model2.ResponseErr{
			Code:   constant.ParamsError,
			ErrMsg: err.Error(),
		})
		return
	}
	resp, count, err := service.GetMenuListService(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model2.ResponseErr{
			Code:   constant.SqlError,
			ErrMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model2.ResponseData{
		Code: 0,
		Msg:  "Get menu list success",
		Data: model2.Data{
			Total: count,
			List:  resp,
		},
	})
}

func ModifyMenu(c *gin.Context) {
	var req model.ModifyMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model2.ResponseErr{
			Code:   constant.ParamsError,
			ErrMsg: err.Error(),
		})
		return
	}
	if err := service.ModifyMenuService(&req); err != nil {
		c.JSON(http.StatusInternalServerError, model2.ResponseErr{
			Code:   constant.SqlError,
			ErrMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model2.Response{
		Code: 0,
		Msg:  "Modify menu success",
	})

}
