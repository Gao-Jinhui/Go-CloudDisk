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
	request := new(v1.UserBasic)
	response := new(v1.CreateUserResponse)
	if code := utils.BindAndValidateParams(c, request); code != e.Success {
		response.Code = code
	} else {
		if !request.IsSufficientToLogin() {
			response.Code = e.InsufficientParameters
		} else {
			response.Code = userServer.CreateUser(request)
		}
	}
	response.Msg = e.GetMsg(response.Code)
	c.JSON(http.StatusOK, response)
	return
}
