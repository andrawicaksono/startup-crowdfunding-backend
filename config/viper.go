package config

import (
	"log"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigType("json")
	config.AddConfigPath(".")
	config.SetConfigName("config")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	return config
}
