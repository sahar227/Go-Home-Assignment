package data

import "validator/internal/dto"

// IDataAccess defines the method we support for accessing the data store
type IDataAccess interface {
	DoesURLExist(urlValidationRequest dto.URLValidationRequest) bool
}
