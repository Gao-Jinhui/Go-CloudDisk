package userController

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/chatserver/service/v1/userServer"
	"go-chat/internal/pkg/core"
	"go-chat/internal/pkg/model/v1/response"
	"go-chat/internal/pkg/model/v1/userModel"
	"go-chat/internal/pkg/system"
	"go-chat/pkg/utils"
)

// CreateUser
// @Summary 创建用户
// @Tags 用户模块
// @Accept  json
// @Produce  json
// @Param data body v1.CreateUserRequest true "请示参数data"
// @Success 200 {object} v1.CreateUserResponse "请求成功"
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	c.Set("logger", system.GetLoggerEntry("Create User"))
	req := new(userModel.UserBasic)
	resp := new(response.BaseResponse)
	if err := utils.BindAndValidateParams(c, req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	if err := req.IsSufficientToLogin(); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	err := userServer.CreateUser(c, req)
	core.WriteResponse(c, err, resp)
	return
}
