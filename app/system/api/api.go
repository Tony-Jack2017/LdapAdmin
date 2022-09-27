package api

import (
	"LdapAdmin/app/system/model"
	"LdapAdmin/app/system/service"
	"LdapAdmin/common/constant"
	model2 "LdapAdmin/common/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

/* $ Api */

func AddApi(c *gin.Context) {
	var req model.AddApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model2.ResponseErr{
			Code:   constant.ParamsError,
			ErrMsg: err.Error(),
		})
		return
	}
	if _, err := service.AddApiService(&req); err != nil {
		c.JSON(http.StatusInternalServerError, model2.ResponseErr{
			Code:   constant.SqlError,
			ErrMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model2.Response{
		Code: 0,
		Msg:  "Add api success",
	})
}

func DeleteApi(c *gin.Context) {
	var req model.DeleteApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model2.ResponseErr{
			Code:   constant.ParamsError,
			ErrMsg: err.Error(),
		})
		return
	}
	if err := service.DeleteApiService(&req); err != nil {
		c.JSON(http.StatusInternalServerError, model2.ResponseErr{
			Code:   constant.SqlError,
			ErrMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model2.Response{
		Code: 0,
		Msg:  "Delete api success",
	})
}

func GetApiList(c *gin.Context) {
	var req model.GetApiListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model2.ResponseErr{
			Code:   constant.ParamsError,
			ErrMsg: err.Error(),
		})
		return
	}
	resp, total, err := service.GetApiListService(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model2.ResponseErr{
			Code:   constant.SqlError,
			ErrMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model2.ResponseData{
		Code: 0,
		Msg:  "Get api list success",
		Data: model2.Data{
			List:  resp,
			Total: total,
		},
	})
}

func ModifyApi(c *gin.Context) {
	var req model.ModifyApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model2.ResponseErr{
			Code:   constant.ParamsError,
			ErrMsg: err.Error(),
		})
		return
	}
	if err := service.ModifyApiService(&req); err != nil {
		c.JSON(http.StatusInternalServerError, model2.ResponseErr{
			Code:   constant.SqlError,
			ErrMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model2.Response{
		Code: 0,
		Msg:  "Modify api success",
	})
}

/* $ ApiGroup */

func AddApiGroup(c *gin.Context) {
	var req model.AddApiGroupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model2.ResponseErr{
			Code:   constant.ParamsError,
			ErrMsg: err.Error(),
		})
		return
	}
	if _, err := service.AddApiGroupService(&req); err != nil {
		c.JSON(http.StatusInternalServerError, model2.ResponseErr{
			Code:   constant.SqlError,
			ErrMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model2.Response{
		Code: 0,
		Msg:  "Add api group success",
	})
}

func DeleteApiGroup(c *gin.Context) {
	var req model.DeleteApiGroupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model2.ResponseErr{
			Code:   constant.ParamsError,
			ErrMsg: err.Error(),
		})
		return
	}
	if err := service.DeleteApiGroupService(&req); err != nil {
		c.JSON(http.StatusInternalServerError, model2.ResponseErr{
			Code:   constant.SqlError,
			ErrMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model2.Response{
		Code: 0,
		Msg:  "Delete api group success",
	})
}

func GetApiGroupList(c *gin.Context) {
	var req model.GetApiGroupListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model2.ResponseErr{
			Code:   constant.ParamsError,
			ErrMsg: err.Error(),
		})
		return
	}
	resp, total, err := service.GetApiGroupListService(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model2.ResponseErr{
			Code:   constant.SqlError,
			ErrMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model2.ResponseData{
		Code: 0,
		Msg:  "Get api group list success",
		Data: model2.Data{
			List:  resp,
			Total: total,
		},
	})
}

func ModifyApiGroup(c *gin.Context) {
	var req model.ModifyApiGroupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model2.ResponseErr{
			Code:   constant.ParamsError,
			ErrMsg: err.Error(),
		})
		return
	}
	if err := service.ModifyApiGroupService(&req); err != nil {
		c.JSON(http.StatusInternalServerError, model2.ResponseErr{
			Code:   constant.SqlError,
			ErrMsg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model2.Response{
		Code: 0,
		Msg:  "Modify api group success",
	})
}
