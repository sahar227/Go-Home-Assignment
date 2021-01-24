package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"validator/internal/data"
	"validator/internal/dto"
)

type responseCache struct {
	isValidated bool
	once        *sync.Once // used to allow us to fetch the value only once for each resource
}

// Server struct holds our server dependencies
type Server struct {
	DataAccess           data.IDataAccess
	mutex                *sync.Mutex
	requestToResponseMap map[dto.URLValidationRequest]responseCache
}

//StartServer starts a server that listens to our endpoint using the given port
func (server Server) StartServer(port int) {
	// set route handler
	http.HandleFunc("/validate-url", server.validateURL)

	// start listening for requests
	fmt.Printf("Listening on port %v...\n", port)
	address := ":" + strconv.Itoa(port)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}

// New Creates a new server instance and initializes it
func New(dataAccess data.IDataAccess) Server {
	server := Server{DataAccess: dataAccess}
	server.mutex = &sync.Mutex{}
	server.requestToResponseMap = make(map[dto.URLValidationRequest]responseCache)
	return server
}
