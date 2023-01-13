package userModel

import (
	"go-chat/internal/pkg/errno"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string `validate:"omitempty,max=30"`
	PassWord      string `validate:"required,max=30"`
	Phone         string `validate:"omitempty,len=11"`
	Email         string `validate:"omitempty,email"`
	Avatar        string //头像
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginInTime   time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (user *UserBasic) UpdateLoginInTime() {
	user.LoginInTime = time.Now()
}

func (user *UserBasic) UpdateHeartbeatTime() {
	user.HeartbeatTime = time.Now()
}

func (user *UserBasic) UpdateLoginOutTime() {
	user.LoginOutTime = time.Now()
}

func (user *UserBasic) IsSufficientToLogin() error {
	if user.Name == "" && user.Phone == "" && user.Email == "" {
		return errno.ErrInsufficientParameters
	}
	return nil
}
