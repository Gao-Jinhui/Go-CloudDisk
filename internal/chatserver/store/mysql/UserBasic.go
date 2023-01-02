package mysql

import (
	"go-chat/internal/pkg/model/chatserver/v1"
	"go-chat/pkg/utils"
	"gorm.io/gorm"
	"time"
)

func FindUserByName(name string) (*v1.UserBasic, error) {
	var user v1.UserBasic
	err := utils.DB.Where("name = ?", name).First(&user).Error
	return &user, err
}

func CreateUser(name, password string) {
	//utils.DB.AutoMigrate(&UserBasic{})
	user := v1.UserBasic{Name: name, PassWord: password, LoginTime: time.Now(), HeartbeatTime: time.Now(), LoginOutTime: time.Now()}
	utils.DB.Create(&user)
	return
}

func DeleteUserByNameAndPassword(name, password string) error {
	user, err := FindUserByName(name)
	if err != gorm.ErrRecordNotFound {
		if user.PassWord != password {
			return utils.ErrWrongPassword
		}
		utils.DB.Delete(user)
		return nil
	} else {
		return err
	}
}
