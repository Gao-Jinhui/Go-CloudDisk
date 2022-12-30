package main

import "go-chat/routers"

func main() {
	r := routers.Router()
	r.Run()
}
