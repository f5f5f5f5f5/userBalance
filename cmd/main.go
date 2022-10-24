package main

import (
	"log"
	balance "userBalance"

	"github.com/f5f5f5f5f5/userBalance/package/handler"

	"github.com/spf13/viper"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(balance.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
	if err := initConfig(); err != nil {
		log.Fatalf("error initialising configs: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
