package v1

import (
	"github.com/gin-gonic/gin"
	userServer "go-chat/internal/chatserver/service/v1"
	"go-chat/internal/pkg/model/chatserver/v1"
	e2 "go-chat/pkg/e"
	"go-chat/pkg/utils"
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
	request := new(v1.CreateUserRequest)
	response := new(v1.CreateUserResponse)
	response.Code = utils.BindAndValidateParams(c, request)
	response.Msg = e2.GetMsg(response.Code)
	if response.Code != e2.SUCCESS {
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data = userServer.CreateUser(request)
	c.JSON(http.StatusOK, response)
	return
}
