package mysql

import (
	"fmt"
	"go-chat/internal/pkg/model/chatserver/v1"
	"go-chat/pkg/utils"
	"gorm.io/gorm"
)

func FindUserByName(name string) (*v1.UserBasic, error) {
	fmt.Println(name)
	var user v1.UserBasic
	if name == "" {
		return &user, utils.ErrRecordNotFound
	}
	err := utils.DB.Where("name = ?", name).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return &user, utils.ErrRecordNotFound
	}
	return &user, nil
}

func FindUserByPhone(phone string) (*v1.UserBasic, error) {
	fmt.Println(phone)
	var user v1.UserBasic
	if phone == "" {
		return &user, utils.ErrRecordNotFound
	}
	err := utils.DB.Where("phone = ?", phone).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return &user, utils.ErrRecordNotFound
	}
	return &user, nil
}

func FindUserByEmail(email string) (*v1.UserBasic, error) {
	fmt.Println(email)
	var user v1.UserBasic
	if email == "" {
		return &user, utils.ErrRecordNotFound
	}
	err := utils.DB.Where("email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return &user, utils.ErrRecordNotFound
	}
	return &user, nil
}

func CreateUser(request *v1.UserBasic) {
	//utils.DB.AutoMigrate(&UserBasic{})
	utils.DB.Create(request)
	return
}
