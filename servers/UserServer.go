package servers

import (
	"go-chat/models"
	"go-chat/utils"
)

func GetUser() models.UserBasic {
	users := models.UserBasic{}
	utils.DB.Find(&users)
	return users
}
