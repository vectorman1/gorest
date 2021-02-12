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
	code int
	body interface{}
}

func GetErrorResponse(err error) (int, HttpResponse) {
	switch err {
	case InvalidModelError:
		return http.StatusBadRequest, HttpResponse{
			code: http.StatusBadRequest,
			body: err.Error(),
		}
	case EntityNotFoundError:
		return http.StatusNotFound, HttpResponse{
			code: http.StatusNotFound,
			body: err.Error(),
		}
	default:
		return http.StatusInternalServerError, HttpResponse{
			code: http.StatusInternalServerError,
			body: err.Error(),
		}
	}
}
