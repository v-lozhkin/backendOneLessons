package app

import (
	"backendOneLessons/lesson4/internal/pkg/api/middlewares"
	imageStore "backendOneLessons/lesson4/internal/pkg/image/storage/fs"
	echoDelivery "backendOneLessons/lesson4/internal/pkg/item/delivery/echo"
	inmemory2 "backendOneLessons/lesson4/internal/pkg/item/repository/inmemory"
	userDelivery "backendOneLessons/lesson4/internal/pkg/user/delivery"
	"backendOneLessons/lesson4/internal/pkg/user/repository/inmemory"
	user "backendOneLessons/lesson4/internal/pkg/user/usecase"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddlewares "github.com/labstack/echo/v4/middleware"
	echolog "github.com/labstack/gommon/log"
)

const apiPort = 8000
const secretKey = "superdupersecret"

func App() {
	items := inmemory2.New()
	images := imageStore.New("/Users/v.lozhkin/go/src/backendOneLessons/lesson4/storage/")

	userRepo := inmemory.New()
	users := user.New(userRepo)
	usersDel := userDelivery.New(users, time.Now().Add(time.Hour*72).Unix(), secretKey)

	delivery := echoDelivery.New(items, images)
	server := echo.New()

	server.Use(echoMiddlewares.Recover())
	server.Use(echoMiddlewares.Logger())
	server.Use(middlewares.RequestIDMiddleware())
	server.Logger.SetLevel(echolog.DEBUG)

	authMiddleware := middlewares.JWTAuthMiddleware(secretKey)

	v1Group := server.Group("/v1")
	itemsGroup := v1Group.Group("/items")
	itemsGroup.GET("", delivery.List)
	itemsGroup.POST("", delivery.Create, authMiddleware)
	itemsGroup.POST("/:id/upload", delivery.Upload, authMiddleware)
	itemsGroup.GET("/:id", delivery.List)
	itemsGroup.PUT("/:id", delivery.Update, authMiddleware)
	itemsGroup.DELETE("/:id", delivery.Delete, authMiddleware)

	v1Group.POST("/user/login", usersDel.Login)

	v1Group.Static("/static", "storage")

	go func() {
		if err := server.Start(fmt.Sprintf(":%d", apiPort)); err != nil && err != http.ErrServerClosed {
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

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDA5NzQ0MTgsIk5hbWUiOiJhZG1pbiJ9.bi1d5xTOAE5g8OecF5EsV6IgCEYtgq0O6dChMsBnBXM
