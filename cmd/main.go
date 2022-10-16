package main

import (
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/service"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/repository"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/handler"
	"syscall"
	"os/signal"
	"os"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main()  {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(app.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoute()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

}