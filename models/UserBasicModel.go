package models

import (
	"go-chat/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name     string
	PassWord string
	Phone    string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email    string `valid:"email"`
	//Avatar        string //头像
	Identity   string
	ClientIp   string
	ClientPort string
	//Salt          string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func FindUserByName(name string) (*UserBasic, error) {
	var user UserBasic
	err := utils.DB.Where("name = ?", name).First(&user).Error
	return &user, err
}

func CreateUser(name, password string) {
	//utils.DB.AutoMigrate(&UserBasic{})
	user := UserBasic{Name: name, PassWord: password, LoginTime: time.Now(), HeartbeatTime: time.Now(), LoginOutTime: time.Now()}
	utils.DB.Create(&user)
	return
}
