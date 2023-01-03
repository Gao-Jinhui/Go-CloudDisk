package v1

import (
	"go-chat/internal/chatserver/store/mysql"
	"go-chat/internal/pkg/model/chatserver/v1"
	"go-chat/pkg/e"
	"go-chat/pkg/utils"
)

func CreateUser(request *v1.UserBasic) int {
	if _, err := mysql.FindUserByName(request.Name); err != utils.ErrRecordNotFound {
		return e.ErrorExistUser
	}
	if _, err := mysql.FindUserByPhone(request.Phone); err != utils.ErrRecordNotFound {
		return e.ErrorExistUser
	}
	if _, err := mysql.FindUserByEmail(request.Email); err != utils.ErrRecordNotFound {
		return e.ErrorExistUser
	}
	request.UpdateLoginInTime()
	request.UpdateHeartbeatTime()
	request.UpdateLoginOutTime()
	mysql.CreateUser(request)
	return e.Success
}
