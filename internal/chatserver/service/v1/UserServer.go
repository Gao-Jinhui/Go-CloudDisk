package userServer

import (
	"go-chat/internal/pkg/model/chatserver/v1"
	"gorm.io/gorm"
)

func CreateUser(request *v1.CreateUserRequest) string {
	if request.Password != request.Repassword {
		return "两次输入的密码不同！"
	}
	if _, err := v1.FindUserByName(request.Name); err != gorm.ErrRecordNotFound {
		return "用户名已存在！"
	}
	v1.CreateUser(request.Name, request.Password)
	return "用户创建成功！"
}
