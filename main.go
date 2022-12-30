package main

import (
	"go-chat/routers"
	"go-chat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQl()
	r := routers.Router()
	r.Run()
}
