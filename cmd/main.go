package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/onemgvv/go-api-server"
	"github.com/onemgvv/go-api-server/pkg/handler"
	"github.com/onemgvv/go-api-server/pkg/repository"
	"github.com/onemgvv/go-api-server/pkg/service"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error config initialization: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading environment variables %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(goApiServer.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured white running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
