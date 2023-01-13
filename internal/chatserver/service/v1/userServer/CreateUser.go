package userServer

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go-chat/internal/chatserver/store/mysql"
	"go-chat/internal/pkg/errno"
	"go-chat/internal/pkg/model/v1/userModel"
	"go-chat/pkg/auth"
	"golang.org/x/net/context"
)

func CreateUser(ctx context.Context, request *userModel.UserBasic) error {
	ctx.Value("logger").(*logrus.Entry).Info("start to create user")
	isNewUser := false
	if _, err := mysql.FindUserByName(ctx, request.Name); errors.Is(err, errno.ErrUserNotFound) {
		isNewUser = true
	}
	if _, err := mysql.FindUserByPhone(ctx, request.Phone); errors.Is(err, errno.ErrUserNotFound) {
		isNewUser = true
	}
	if _, err := mysql.FindUserByName(ctx, request.Name); errors.Is(err, errno.ErrUserNotFound) {
		isNewUser = true
	}
	if !isNewUser {
		ctx.Value("logger").(*logrus.Entry).Error("user already exist")
		return errno.ErrUserAlreadyExist
	}
	request.UpdateLoginInTime()
	request.UpdateHeartbeatTime()
	request.UpdateLoginOutTime()
	if pwd, err := auth.Encrypt(request.PassWord); err != nil {
		ctx.Value("logger").(*logrus.Entry).Error("failed to get encrypted password")
		return errno.ErrEncrypt
	} else {
		request.PassWord = pwd
	}
	if err := mysql.CreateUser(ctx, request); err != nil {
		ctx.Value("logger").(*logrus.Entry).Error("failed to create user")
		return err
	}
	ctx.Value("logger").(*logrus.Entry).Info("create user successfully")
	return nil
}
