package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-chat/servers"
)

// TestController
// @Summary 获取用户信息
// @Tags 用户模块
// @Success 200 {string} json{"code","message"} 成功后返回值
// @Router /test [get]
func TestController(c *gin.Context) {
	users := servers.GetUser()
	fmt.Println(users)
	c.JSON(200, gin.H{
		"message": users,
	})
}
