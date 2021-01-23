package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	"validator/internal/data"
	"validator/internal/dto"

	"github.com/mailru/easyjson"
)

type responseCache struct {
	value bool
	once  *sync.Once // used to allow us to fetch the value only once for each resource
}

// Server struct holds our server dependencies
type Server struct {
	DataAccess           data.IDataAccess
	mutex                *sync.Mutex
	requestToResponseMap map[dto.URLValidationRequest]responseCache
}

// Handles /validate-url path
func (server Server) validateURL(w http.ResponseWriter, r *http.Request) {
	// reject non POST requests because we do not handle them
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}

	// read request data from body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// unmarshall request data into an object
	var request dto.URLValidationRequest
	err = easyjson.Unmarshal(bodyBytes, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// initialize entry in map if not already initialized
	server.mutex.Lock()
	if _, ok := server.requestToResponseMap[request]; !ok {
		server.requestToResponseMap[request] = responseCache{value: false, once: &sync.Once{}}
	}
	server.mutex.Unlock()

	// fetch data from data store only once per combination of domain and path
	server.requestToResponseMap[request].once.Do(func() {
		var tempCache = server.requestToResponseMap[request]
		tempCache.value = server.DataAccess.DoesURLExist(request)
		server.requestToResponseMap[request] = tempCache
	})

	// create response
	location := ""
	if server.requestToResponseMap[request].value {
		location = request.Domain + request.Path
	}
	responseObject := dto.URLValidationResponse{Location: location}
	responseJSON, err := easyjson.Marshal(responseObject)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(responseJSON)
}

//StartServer starts a server that listens to our endpoint using the given port
func (server Server) StartServer(port int) {
	http.HandleFunc("/validate-url", server.validateURL)
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
