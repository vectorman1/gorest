package common

import (
	"errors"
	"net/http"
)

var (
	InvalidModelError   = errors.New("invalid model")
	EntityNotFoundError = errors.New("entity not found")
)

type HttpResponse struct {
	Code int
	Body interface{}
}

func GetErrorResponse(err error) (int, HttpResponse) {
	switch err {
	case InvalidModelError:
		return http.StatusBadRequest, HttpResponse{
			Code: http.StatusBadRequest,
			Body: err.Error(),
		}
	case EntityNotFoundError:
		return http.StatusNotFound, HttpResponse{
			Code: http.StatusNotFound,
			Body: err.Error(),
		}
	default:
		return http.StatusInternalServerError, HttpResponse{
			Code: http.StatusInternalServerError,
			Body: err.Error(),
		}
	}
}
