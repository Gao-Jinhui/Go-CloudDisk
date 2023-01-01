package userServer

import (
	"go-chat/models"
	"gorm.io/gorm"
)

func CreateUser(request *models.CreateUserRequest) string {
	if request.Password != request.Repassword {
		return "两次输入的密码不同！"
	}
	if _, err := models.FindUserByName(request.Name); err != gorm.ErrRecordNotFound {
		return "用户名已存在！"
	}
	models.CreateUser(request.Name, request.Password)
	return "用户创建成功！"
}
