package main

import (
	"SalimProject/pkg/handler"
	"SalimProject/pkg/repository"
	"SalimProject/pkg/service"
	"SalimProject/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error with initializing config: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error with initializing .env file: %s", err.Error())
	}
	db, err := repository.NewDataBase(repository.Config{
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		Port:     viper.GetString("db.port"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("Error with database config initializing: %s", err.Error())
	}
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRouter()); err != nil {
			logrus.Fatal("Error in cmd main - ", err)
		}
	}()
	logrus.Println("Project Shutting Down")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("Project Shutting Down")
	if err := srv.Shutdown(); err != nil {
		logrus.Errorf("Error occured on server shutting down: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
