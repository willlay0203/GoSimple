package server

import (
	"net/http"
)

type APIError struct {
	StatusCode int
	Msg        string
}

func (e APIError) Error() string {
	return e.Msg
}

func CreateAPIErrorResponse(statusCode int, msg string) APIError {
	return APIError{StatusCode: statusCode, Msg: msg}
}

func CreateDefaultAPIErrorResponse(statusCode int) APIError {
	switch statusCode {
	case http.StatusBadRequest:
		return APIError{StatusCode: http.StatusBadRequest, Msg: "400 - Bad Request"}
	case http.StatusForbidden:
		return APIError{StatusCode: http.StatusForbidden, Msg: "403 - Forbidden"}
	case http.StatusNotFound:
		return APIError{StatusCode: http.StatusNotFound, Msg: "404 - Not Found"}
	default:
		return APIError{StatusCode: http.StatusInternalServerError, Msg: "500 - Internal Server Error"}
	}
}
