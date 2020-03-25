package main

import (
	"fmt"

	"server/libs/configuration"
	"server/src/containers/repository"
	"server/src/containers/service"
	"server/src/http"
	"server/src/sqlite"
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
	server := http.NewServer(services)

    fmt.Println("Listening on", port)
    server.Start(port)
}
