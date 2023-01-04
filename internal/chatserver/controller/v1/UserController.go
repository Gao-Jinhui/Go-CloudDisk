package v1

import (
	"github.com/gin-gonic/gin"
	userServer "go-chat/internal/chatserver/service/v1"
	"go-chat/internal/pkg/model/chatserver/v1"
	system "go-chat/pkg/System"
	"go-chat/pkg/core"
	"go-chat/pkg/utils"
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
	c.Set("logger", system.GetLoggerEntry(">>>>test<<<"))
	request := new(v1.UserBasic)
	response := new(v1.BaseResponse)
	if err := utils.BindAndValidateParams(c, request); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	if err := request.IsSufficientToLogin(); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	err := userServer.CreateUser(c, request)
	core.WriteResponse(c, err, response)
	return
}
