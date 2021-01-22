package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"validator/internal/data"
	"validator/internal/dto"

	"github.com/mailru/easyjson"
)

// Server ...
type Server struct {
	DataAccess data.IDataAccess
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

	// create response
	location := ""
	if server.DataAccess.DoesURLExist(request) {
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
