package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-chat/internal/pkg/errno"
)

func BindAndValidateParams(c *gin.Context, request interface{}) error {
	if err := c.ShouldBindJSON(request); err != nil {
		return errno.ErrBind
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		fmt.Println(err.Error())
		return errno.ErrValidation
	}
	return nil
}
