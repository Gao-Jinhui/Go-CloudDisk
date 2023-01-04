package main

import (
	"go-chat/internal/chatserver"
	"go-chat/pkg/System"
)

func main() {
	system.InitLogger()
	system.InitConfig()
	system.InitMySQl()
	r := chatserver.Router()
	err := r.Run()
	if err != nil {
		return
	}
}
