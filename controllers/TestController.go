package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-chat/servers"
)

func TestController(c *gin.Context) {
	users := servers.GetUser()
	fmt.Println(users)
	c.JSON(200, gin.H{
		"message": users,
	})
}
