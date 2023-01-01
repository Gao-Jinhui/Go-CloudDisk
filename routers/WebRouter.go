package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-chat/controllers"
	"go-chat/docs"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "" //TODO:?
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	userRouter := r.Group("/user")
	{
		userRouter.POST("/createUser", controllers.CreateUserController)
	}
	return r
}
