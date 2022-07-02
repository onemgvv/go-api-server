package main

import (
	"context"
	"errors"
	"github.com/joho/godotenv"
	"github.com/onemgvv/go-api-server/internal/config"
	deliveryHttp "github.com/onemgvv/go-api-server/internal/delivery/http"
	"github.com/onemgvv/go-api-server/internal/repository"
	"github.com/onemgvv/go-api-server/internal/server"
	"github.com/onemgvv/go-api-server/internal/service"
	"github.com/onemgvv/go-api-server/pkg/database/postgres"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const configDir = "configs"

// @title Go API Service
// @version 0.1
// @description API Server Template for GO Apps

// @host localhost:8000
// @basePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("[ENV Load] || [Failed]: %s", err.Error())
	}

	cfg, err := config.Init(configDir)
	if err != nil {
		log.Fatalf("[Config Load] || [Failed]: %s", err.Error())
	}

	db, err := postgres.Init(cfg)
	if err != nil {
		log.Fatalf("[Database INIT] || [Failed]: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(&service.Deps{
		Repos: repositories,
	})

	handlers := deliveryHttp.NewHandler(services)

	app := server.NewServer(cfg, handlers.InitRoutes())

	go func() {
		if err = app.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("[SERVER START] || [FAILED]: %s", err.Error())
		}
	}()

	log.Println("Application started")

	/**
	 *	Graceful Shutdown
	 */
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err = app.Shutdown(ctx); err != nil {
		log.Fatalf("[SERVER STOP] || [FAILED]: %s", err.Error())
	}

	if err = postgres.Close(db); err != nil {
		log.Fatalf("[DATABASE CONN CLOSE] || [FAILED]: %s", err.Error())
	}
}
