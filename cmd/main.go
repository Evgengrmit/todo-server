package main

import (
	"log"
	"todo"
	"todo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.MyServer)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred when starting the server: %s", err.Error())
	}
}
