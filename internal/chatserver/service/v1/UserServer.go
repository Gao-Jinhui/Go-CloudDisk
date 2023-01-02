package v1

import (
	"go-chat/internal/chatserver/store/mysql"
	"go-chat/internal/pkg/model/chatserver/v1"
	"go-chat/pkg/e"
	"go-chat/pkg/utils"
	"gorm.io/gorm"
)

func CreateUser(request *v1.CreateUserRequest) int {
	if _, err := mysql.FindUserByName(request.Name); err != gorm.ErrRecordNotFound {
		return e.ErrorExistUser
	}
	mysql.CreateUser(request.Name, request.Password)
	return e.Success
}

func DeleteUser(request *v1.DeleteUserRequest) int {
	err := mysql.DeleteUserByNameAndPassword(request.Name, request.Password)
	if err == utils.ErrWrongPassword {
		return e.ErrorWrongPassword
	} else if err == gorm.ErrRecordNotFound {
		return e.ErrorNotExistUser
	}
	return e.Success
}
