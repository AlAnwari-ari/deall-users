package error

import "net/http"

var (
	ErrDataNotFound      = NewErrorf(http.StatusNotFound, "Data not found")
	ErrDataAlreadyExist  = NewErrorf(http.StatusBadRequest, "Data already exist")
	ErrTimeout           = NewErrorf(http.StatusRequestTimeout, "Timeout")
	ErrRelationNotFound  = NewErrorf(http.StatusNotFound, "Relation Data Not Found")
	ErrInvalidToken      = NewErrorf(http.StatusUnauthorized, "Invalid token")
	ErrUnathorized       = NewErrorf(http.StatusUnauthorized, "Unauthorized")
	ErrTokenExipred      = NewErrorf(http.StatusUnauthorized, "Token expired")
	ErrPayloadValidation = NewErrorf(http.StatusBadRequest, "Your data is/are not valid as requirement needs. Please check carefully!")
	ErrInternalServErr   = NewErrorf(http.StatusInternalServerError, "Internal Server Error")
	ErrUnknownError      = NewErrorf(http.StatusInternalServerError, "Unknown error")
	ErrAuthLogin         = NewErrorf(http.StatusNotFound, "Please check again your username and/or password!")
)
