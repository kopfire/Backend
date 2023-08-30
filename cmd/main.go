package main

import (
	"AvitoBackend/internal/server"
	"AvitoBackend/internal/storage"
	"log"
	"os"
)

func main() {
	storagePostgres, err := storage.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	if err := storagePostgres.Init(); err != nil {
		log.Fatal(err)
	}

	apiServer := server.NewAPIServer(":"+os.Getenv("APP_PORT"), storagePostgres)
	apiServer.Run()
}
