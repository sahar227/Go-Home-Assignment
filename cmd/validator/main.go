package main

import (
	"validator/internal/data"
	"validator/internal/server"
)

func main() {
	// In a production environment we would create an instance of something like SQLDatastore here and just make sure it implemets IDataAccess interface
	dataStore := data.InMemoryDataAccess{}
	appServer := server.Server{DataAccess: dataStore}
	appServer.StartServer(8080)
}
