package system

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("../../config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	Logger.Info("Configuration initialization succeeded ")
}
