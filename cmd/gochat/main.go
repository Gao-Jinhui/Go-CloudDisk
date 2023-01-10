package main

import (
	"go-chat/internal/chatserver"
	"go-chat/internal/pkg/system"
)

func main() {
	system.InitLogger()
	system.InitConfig()
	system.InitMySQl()
	system.InitRedis()
	r := chatserver.Router()
	err := r.Run()
	if err != nil {
		system.Logger.Error("failed to run server !")
		return
	}
}
