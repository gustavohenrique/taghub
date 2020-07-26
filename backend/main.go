package main

import (
	"fmt"

	"backend/libs/configuration"
	"backend/src/containers/repository"
	"backend/src/containers/service"
	"backend/src/http"
	"backend/src/sqlite"
)

func main() {
	config := configuration.Load()
	db := sqlite.New(sqlite.Config{
		URL: config.DatabaseURL,
	})
	db.Connect()

	repositories := repository.NewRepositoryContainer(db)
	services := service.NewServiceContainer(repositories)

	port := ":" + config.HTTPPort
	backend := http.NewServer(services)

	fmt.Println("Listening on", port)
	backend.Start(port)
}
