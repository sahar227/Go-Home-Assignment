package server

import (
	"io/ioutil"
	"net/http"
	"sync"
	"validator/internal/dto"

	"github.com/mailru/easyjson"
)

// Handles /validate-url path according to the request method
func (server Server) validateURL(w http.ResponseWriter, r *http.Request) {
	// reject non POST requests because we do not handle them
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}

	server.validateURLPost(w, r)
}

func getRequestDataFromBody(w http.ResponseWriter, r *http.Request) (request dto.URLValidationRequest, ok bool) {
	// read request data from body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return request, false
	}

	// unmarshall request data into an object
	err = easyjson.Unmarshal(bodyBytes, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return request, false
	}

	return request, true
}

// Handles /validate-url path on POST requests
func (server Server) validateURLPost(w http.ResponseWriter, r *http.Request) {
	request, ok := getRequestDataFromBody(w, r)
	if !ok {
		return
	}

	// initialize entry in map if not already initialized
	server.mutex.Lock()
	if _, ok := server.requestToResponseMap[request]; !ok {
		server.requestToResponseMap[request] = responseCache{isValidated: false, once: &sync.Once{}}
	}
	server.mutex.Unlock()

	// fetch data from data store only once per combination of domain and path
	server.requestToResponseMap[request].once.Do(func() {
		var tempCache = server.requestToResponseMap[request]
		tempCache.isValidated = server.DataAccess.DoesURLExist(request)
		server.requestToResponseMap[request] = tempCache
	})

	// create response
	location := ""
	if server.requestToResponseMap[request].isValidated {
		location = "https://" + request.Domain + request.Path
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
