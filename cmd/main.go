package main

import (
	"log"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	serv := service.NewServer(repos)
	handl := handler.NewHandler(serv)
	srv := new(todo.MyServer)
	if err := srv.Run("8080", handl.InitRoutes()); err != nil {
		log.Fatalf("error occurred when starting the server: %s", err.Error())
	}
}
