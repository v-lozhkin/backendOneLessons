package app

import (
	imageStore "backendOneLessons/lesson4/internal/pkg/image/storage/fs"
	itemDelivery "backendOneLessons/lesson4/internal/pkg/item/delivery"
	"backendOneLessons/lesson4/internal/pkg/item/usecase"
	"fmt"
	"log"
	"net/http"
	"time"
)

const apiPort = 8000
const filePort = 8080

func App() {
	go startFileServer()
	startApiServer()
}

func startApiServer() {
	items := usecase.NewInmemory()
	images := imageStore.New("/Users/v.lozhkin/go/src/backendOneLessons/lesson4/storage/")

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
