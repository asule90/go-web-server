package config

import (
	"github.com/spf13/viper"
	"go-web-server/utils"
)

func InitConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	utils.InitDB(
		viper.GetString("application.resources.db.driver"),
		viper.GetString("application.resources.db.host"),
		viper.GetInt32("application.resources.db.port"),
		viper.GetString("application.resources.db.name"),
		viper.GetString("application.resources.db.user"),
		viper.GetString("application.resources.db.password"),
	)
}
