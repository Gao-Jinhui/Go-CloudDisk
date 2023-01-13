package chatserver

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-chat/docs"
	"go-chat/internal/chatserver/controller/v1/userController"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	userRouter := r.Group("/user")
	{
		userRouter.POST("/createUser", userController.CreateUser)
	}
	return r
}
