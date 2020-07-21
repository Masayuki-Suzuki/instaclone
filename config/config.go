package config

import (
	"log"
	
	"github.com/spf13/viper"
)

var c *viper.Viper

func Init(env string) {
	c = viper.New()
	c.SetConfigFile("yaml")
	c.SetConfigName(env)
	c.AddConfigPath("config/environments")
	
	if err := c.ReadInConfig(); err != nil {
		log.Fatalf("Cannot read config file: %v\n", err)
	}
}

func GetConfig() *viper.Viper {
	return c
}
