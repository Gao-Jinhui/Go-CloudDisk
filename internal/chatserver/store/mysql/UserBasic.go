package mysql

import (
	"fmt"
	"github.com/pkg/errors"
	"go-chat/internal/pkg/errno"
	"go-chat/internal/pkg/model/v1/userModel"
	"go-chat/internal/pkg/system"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func FindUserByName(ctx context.Context, name string) (*userModel.UserBasic, error) {
	var user userModel.UserBasic
	err := system.DB.Where("name = ?", name).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func FindUserByPhone(ctx context.Context, phone string) (*userModel.UserBasic, error) {
	fmt.Println(phone)
	var user userModel.UserBasic
	err := system.DB.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func FindUserByEmail(ctx context.Context, email string) (*userModel.UserBasic, error) {
	fmt.Println(email)
	var user userModel.UserBasic
	err := system.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(ctx context.Context, request *userModel.UserBasic) error {
	return system.DB.Create(request).Error
}
