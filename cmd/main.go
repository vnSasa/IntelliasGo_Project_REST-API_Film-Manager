package main

import (
	_ "github.com/lib/pq"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/service"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/repository"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/handler"
	"syscall"
	"os/signal"
	"os"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
)

func main()  {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	
	repos := repository.NewRepository(db)
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

func initConfig() error  {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}