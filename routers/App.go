package routers

import (
	"github.com/gin-gonic/gin"
	"go-chat/controllers"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/test", controllers.TestController)
	return r
}
