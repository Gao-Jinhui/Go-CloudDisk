package core

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/pkg/errno"
	v1 "go-chat/internal/pkg/model/chatserver/v1"
	"net/http"
)

func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		code, msg := errno.DecodeErr(err)
		c.JSON(http.StatusOK, v1.BaseResponse{
			Code: code,
			Msg:  msg,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
