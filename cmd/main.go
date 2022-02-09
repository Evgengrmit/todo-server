package main

import (
	"github.com/spf13/viper"
	"log"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error occurred when initialization configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	serv := service.NewServer(repos)
	handl := handler.NewHandler(serv)
	srv := new(todo.MyServer)
	if err := srv.Run(viper.GetString("port"), handl.InitRoutes()); err != nil {
		log.Fatalf("error occurred when starting the server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
