package system

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQl() {
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{})
	//DB.AutoMigrate(v1.UserBasic{})
	Logger.Info("Database initialization succeeded ")
}
