package app

import (
	"backendOneLessons/lesson4/internal/pkg/image"
	imageStore "backendOneLessons/lesson4/internal/pkg/image/storage/fs"
	"backendOneLessons/lesson4/internal/pkg/item"
	itemDelivery "backendOneLessons/lesson4/internal/pkg/item/delivery/default_http"
	echoDelivery "backendOneLessons/lesson4/internal/pkg/item/delivery/echo"
	"backendOneLessons/lesson4/internal/pkg/item/usecase"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echolog "github.com/labstack/gommon/log"
	"log"
	"net/http"
	"time"
)

const apiPort = 8000
const filePort = 8080

func App() {
	items := usecase.NewInmemory()
	images := imageStore.New("/Users/v.lozhkin/go/src/backendOneLessons/lesson4/storage/")

	startEchoServer(items, images)
}

func startEchoServer(items item.ItemUsecase, images image.Storage) {
	delivery := echoDelivery.New(items, images)
	server := echo.New()

	server.Use(middleware.Recover())
	server.Use(middleware.Logger())
	server.Logger.SetLevel(echolog.DEBUG)

	v1Group := server.Group("/v1")
	itemsGroup := v1Group.Group("/items")
	itemsGroup.GET("", delivery.List)
	itemsGroup.POST("", delivery.Create)
	itemsGroup.POST("/:id/upload", delivery.Upload)
	itemsGroup.GET("/:id", delivery.List)
	itemsGroup.PUT("/:id", delivery.Update)
	itemsGroup.DELETE("/:id", delivery.Delete)

	v1Group.Static("/static", "storage")

	log.Fatal(server.Start(fmt.Sprintf(":%d", apiPort)))

}

func startApiServer(items item.ItemUsecase, images image.Storage) {
	itemHandler := itemDelivery.New(items, images)

	http.Handle("/items/", itemHandler)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", apiPort),
		Handler:      nil,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	log.Printf("api start listening on %d port\n", apiPort)
	log.Fatal(server.ListenAndServe())
}

func startFileServer() {
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", filePort),
		Handler:      http.FileServer(http.Dir("/Users/v.lozhkin/go/src/backendOneLessons/lesson4/storage/")),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Printf("fs starts listening on %d port\n", filePort)
	log.Fatal(server.ListenAndServe())
}
