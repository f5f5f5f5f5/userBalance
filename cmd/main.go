package main

import (
	"log"

	"github.com/f5f5f5f5f5/userBalance_service/package/handler"
	"github.com/f5f5f5f5f5/userBalance_service/server"

	"github.com/spf13/viper"
)

func main() {
	handlers := handler.Handler{}
	srv := new(server.Server)
	if err := srv.Run(); err != nil {
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
