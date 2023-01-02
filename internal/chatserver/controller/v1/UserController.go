package v1

import (
	"github.com/gin-gonic/gin"
	userServer "go-chat/internal/chatserver/service/v1"
	"go-chat/internal/pkg/model/chatserver/v1"
	e "go-chat/pkg/e"
	"go-chat/pkg/utils"
	"net/http"
)

// CreateUserController
// @Summary 创建用户
// @Tags 用户模块
// @Accept  json
// @Produce  json
// @Param data body v1.CreateUserRequest true "请示参数data"
// @Success 200 {object} v1.CreateUserResponse "请求成功"
// @Router /user/createUser [post]
func CreateUserController(c *gin.Context) {
	request := new(v1.CreateUserRequest)
	response := new(v1.CreateUserResponse)
	if code := utils.BindAndValidateParams(c, request); code != e.Success {
		response.Code = code
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = userServer.CreateUser(request)
	response.Msg = e.GetMsg(response.Code)
	c.JSON(http.StatusOK, response)
	return
}

// DeleteUserController
// @Summary 删除用户
// @Tags 用户模块
// @Accept  json
// @Produce  json
// @Param data body v1.DeleteUserRequest true "请示参数data"
// @Success 200 {object} v1.DeleteUserResponse "请求成功"
// @Router /user/deleteUser [post]
func DeleteUserController(c *gin.Context) {
	request := new(v1.DeleteUserRequest)
	response := new(v1.DeleteUserResponse)
	if code := utils.BindAndValidateParams(c, request); code != e.Success {
		response.Code = code
		response.Msg = e.GetMsg(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = userServer.DeleteUser(request)
	response.Msg = e.GetMsg(response.Code)
	c.JSON(http.StatusOK, response)
	return
}
