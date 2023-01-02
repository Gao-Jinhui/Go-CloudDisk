package utils

import (
	"go-chat/pkg/e"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidateParams(c *gin.Context, request interface{}) int {
	if err := c.ShouldBindJSON(request); err != nil {
		return e.InvalidParams
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return e.ErrorValidator
	}
	return e.Success
}
