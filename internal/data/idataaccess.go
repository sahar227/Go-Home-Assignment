package data

import "validator/internal/dto"

type IDataAccess interface {
	DoesURLExist(urlValidationRequest dto.URLValidationRequest) bool
}
