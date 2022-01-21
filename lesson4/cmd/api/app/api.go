package app

import (
	"backendOneLessons/lesson4/cmd/api/app/config"
	"backendOneLessons/lesson4/internal/pkg/api/middlewares"
	imageStore "backendOneLessons/lesson4/internal/pkg/image/storage/fs"
	echoDelivery "backendOneLessons/lesson4/internal/pkg/item/delivery/echo"
	itemPostgresRepo "backendOneLessons/lesson4/internal/pkg/item/repository/postgres"
	itemUsecase "backendOneLessons/lesson4/internal/pkg/item/usecase"
	userDelivery "backendOneLessons/lesson4/internal/pkg/user/delivery"
	userRepo "backendOneLessons/lesson4/internal/pkg/user/repository/inmemory"
	user "backendOneLessons/lesson4/internal/pkg/user/usecase"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/labstack/echo/v4"
	echoMiddlewares "github.com/labstack/echo/v4/middleware"
	echolog "github.com/labstack/gommon/log"

	// postgres driver
	_ "github.com/lib/pq"

	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func App() {
	server := echo.New()
	server.Logger.SetLevel(echolog.INFO)

	cfg := config.Config{}
	cfg.ReadFromFile(server.Logger)

	server.Use(echoMiddlewares.Recover())
	server.Use(echoMiddlewares.Logger())
	server.Use(middlewares.RequestIDMiddleware())
	server.Use(echoMiddlewares.TimeoutWithConfig(echoMiddlewares.TimeoutConfig{
		Timeout: time.Second * 10,
	}))

	loglevel, ok := loglevelMap[cfg.Loglevel]
	if !ok {
		loglevel = echolog.INFO
	}
	server.Logger.SetLevel(loglevel)

	stat := promauto.Factory{}

	authMiddleware := middlewares.JWTAuthMiddleware(cfg.AuthConfig.JWTSecret)

	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("user=%s password=%s port=%d dbname=%s sslmode=disable host=%s",
			cfg.DBConfig.User,
			cfg.DBConfig.Password,
			cfg.DBConfig.Port,
			cfg.DBConfig.DBName,
			cfg.DBConfig.Host,
		),
	)
	if err != nil {
		server.Logger.Fatalf("failed to open db connection %v", err)
	}

	itemsRepository := itemPostgresRepo.New(db, stat)

	userRepository := userRepo.New()

	itemsUsecase := itemUsecase.New(itemsRepository, stat)
	images := imageStore.New(cfg.StoragePath, stat)
	usersUsecase := user.New(userRepository)

	itemsDelivery := echoDelivery.New(itemsUsecase, images, stat)
	usersDelivery := userDelivery.New(
		usersUsecase,
		time.Now().Add(cfg.AuthConfig.JWTTTL).Unix(),
		cfg.AuthConfig.JWTSecret,
	)

	v1Group := server.Group("/v1")
	itemsGroup := v1Group.Group("/items")
	itemsGroup.GET("", itemsDelivery.List)
	itemsGroup.POST("", itemsDelivery.Create, authMiddleware)
	itemsGroup.POST("/:id/upload", itemsDelivery.Upload, authMiddleware)
	itemsGroup.GET("/:id", itemsDelivery.List)
	itemsGroup.PUT("/:id", itemsDelivery.Update, authMiddleware)
	itemsGroup.DELETE("/:id", itemsDelivery.Delete, authMiddleware)

	v1Group.POST("/user/login", usersDelivery.Login)

	v1Group.Static("/static", "storage")

	server.Any("/metrics", func(ectx echo.Context) error {
		promhttp.Handler().ServeHTTP(ectx.Response().Writer, ectx.Request())
		return nil
	})

	go func() {
		if err := server.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil && err != http.ErrServerClosed {
			server.Logger.Fatal(err)
		}
	}()

	quite := make(chan os.Signal, 1)
	signal.Notify(quite, syscall.SIGINT, syscall.SIGTERM)
	<-quite
	server.Logger.Info("shutdown inited")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}

var loglevelMap = map[string]echolog.Lvl{
	"debug": echolog.DEBUG,
	"info":  echolog.INFO,
	"error": echolog.ERROR,
	"warn":  echolog.WARN,
	"off":   echolog.OFF,
}
