package utils

import (
	"log"

	"github.com/spf13/viper"
)

func Config(env string) string {
	viper.SetConfigFile("./config.toml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	return viper.GetString(env)
}
