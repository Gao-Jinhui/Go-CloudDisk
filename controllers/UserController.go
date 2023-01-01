package controllers

import (
	"github.com/gin-gonic/gin"
	"go-chat/e"
	"go-chat/models"
	userServer "go-chat/servers"
	"go-chat/utils"
	"net/http"
)

// CreateUserController
// @Summary 创建用户
// @Tags 用户模块
// @Accept  json
// @Produce  json
// @Param data body models.CreateUserRequest true "请示参数data"
// @Success 200 {object} models.CreateUserResponse "请求成功"
// @Router /user/createUser [post]
func CreateUserController(c *gin.Context) {
	request := new(models.CreateUserRequest)
	response := new(models.CreateUserResponse)
	response.Code = utils.BindAndValidateParams(c, request)
	response.Msg = e.GetMsg(response.Code)
	if response.Code != e.SUCCESS {
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data = userServer.CreateUser(request)
	c.JSON(http.StatusOK, response)
	return
}
