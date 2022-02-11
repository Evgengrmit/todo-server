package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error occurred when initialization configs: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("error occured when loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("error occurred when initialization database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	handl := handler.NewHandler(serv)

	srv := new(todo.MyServer)
	go func() {
		if err = srv.Run(viper.GetString("port"), handl.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred when starting the server: %s", err.Error())
		}
	}()
	logrus.Print("Application started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("Application shutting down")
	if err = srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred when shutting down the server: %s", err.Error())
	}
	if err = db.Close(); err != nil {
		logrus.Errorf("error occurred when db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
