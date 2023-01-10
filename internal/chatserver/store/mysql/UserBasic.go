package mysql

import (
	"fmt"
	"github.com/pkg/errors"
	"go-chat/internal/pkg/errno"
	"go-chat/internal/pkg/model/chatserver/v1"
	"go-chat/internal/pkg/system"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func FindUserByName(ctx context.Context, name string) (*v1.UserBasic, error) {
	var user v1.UserBasic
	err := system.DB.Where("name = ?", name).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func FindUserByPhone(ctx context.Context, phone string) (*v1.UserBasic, error) {
	fmt.Println(phone)
	var user v1.UserBasic
	err := system.DB.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func FindUserByEmail(ctx context.Context, email string) (*v1.UserBasic, error) {
	fmt.Println(email)
	var user v1.UserBasic
	err := system.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(ctx context.Context, request *v1.UserBasic) error {
	return system.DB.Create(request).Error
}
