package test

import (
	"fmt"
	"go-chat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	dsn := "root:9811@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to open database")
	}

	db.AutoMigrate(&models.UserBasic{})

	user := &models.UserBasic{Name: "test1", LoginTime: time.Now(), HeartbeatTime: time.Now(), LoginOutTime: time.Now()}
	db.Create(user)
}
