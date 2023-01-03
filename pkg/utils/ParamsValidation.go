package utils

import (
	"fmt"
	"go-chat/pkg/e"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidateParams(c *gin.Context, request interface{}) int {
	if err := c.ShouldBindJSON(request); err != nil {
		return e.InvalidRequest
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		fmt.Println(err.Error())
		return e.InvalidParams
	}
	return e.Success
}
