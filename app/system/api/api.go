package api

import "github.com/gin-gonic/gin"

// AddApi
// @Summary      Add Server Api
// @Description: That is for admin add the server api
// @Tags         Api
// @Accept       application/json
// @Produce      application/json
// @Param        object body     model.AddReq false "the data for adding api"
// @Success      200    {object} model.ResponseCommon
// @Failure      400    {object} model.ResponseErr
// @Failure      500    {object} model.ResponseErr
// @Route        "/request/base/api/add" [post]
func AddApi(c *gin.Context) {

}

func DeleteApi(c *gin.Context) {

}

func GetApiList(c *gin.Context) {

}

func ModifyApi(c *gin.Context) {

}
