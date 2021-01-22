package data

import (
	"validator/internal/dto"
	"validator/internal/model"
)

// In memory data store for development. In production our data would be stored externally.
var dataStore = []model.URL{
	model.URL{Domain: "ynte.co.il", Path: "/home/0,7340,L-8,00.html"},
	model.URL{Domain: "walla.co.il", Path: "/"},
	model.URL{Domain: "google.com", Path: "/index.html"}}

// Implements IDataAccess interface
type InMemoryDataAccess struct {
}

func (InMemoryDataAccess) DoesURLExist(urlValidationRequest dto.URLValidationRequest) bool {
	for i := 0; i < len(dataStore); i++ {
		url := dataStore[i]
		if url.Domain == urlValidationRequest.Domain && url.Path == urlValidationRequest.Path {
			return true
		}
	}
	return false
}
