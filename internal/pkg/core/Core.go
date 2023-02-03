package core

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/pkg/errno"
	"go-chat/internal/pkg/model/v1/response"
	"net/http"
)

func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		code, msg := errno.DecodeErr(err)
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code: code,
			Msg:  msg,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
