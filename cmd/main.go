package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	srv := new(server.Server)
	srv.Run(8080)
	if err := initConfig(); err != nil {
		log.Fatalf("error initialising configs: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
