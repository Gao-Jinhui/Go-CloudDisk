package main

import (
	"go-chat/internal/chatserver"
	"go-chat/pkg/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQl()
	r := chatserver.Router()
	err := r.Run()
	if err != nil {
		return
	}
}
