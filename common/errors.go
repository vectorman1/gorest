package common

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

const DUPLICATE_ENTRY = 1062
const DATA_TOO_LONG = 1406
const CANNOT_BE_NULL = 1048

var (
	InvalidModelError     = errors.New("invalid model")
	EntityNotFoundError   = errors.New("entity not found")
	DuplicateEntityError  = errors.New("duplicate entity by unique column")
	PasswordTooShortError = errors.New("password is too short")
	DataTooLongError      = errors.New("a field exceeds it's max length")
	CannotBeNullError     = errors.New("a field is null when it cannot be")
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
		me, ok := err.(*mysql.MySQLError)
		if ok {
			domainErr := ParseMySQLError(me)
			switch domainErr {
			case DuplicateEntityError:
				return http.StatusBadRequest, HttpResponse{
					Code: http.StatusBadRequest,
					Body: domainErr.Error(),
				}
			case DataTooLongError:
				return http.StatusBadRequest, HttpResponse{
					Code: http.StatusBadRequest,
					Body: domainErr.Error(),
				}
			case CannotBeNullError:
				return http.StatusBadRequest, HttpResponse{
					Code: http.StatusBadRequest,
					Body: domainErr.Error(),
				}
			default:
				return http.StatusInternalServerError, HttpResponse{
					Code: http.StatusInternalServerError,
					Body: domainErr.Error(),
				}
			}
		}
		return http.StatusInternalServerError, HttpResponse{
			Code: http.StatusInternalServerError,
			Body: err.Error(),
		}
	}
}

func ParseMySQLError(sqlError *mysql.MySQLError) error {
	switch sqlError.Number {
	case DUPLICATE_ENTRY:
		return DuplicateEntityError
	case DATA_TOO_LONG:
		return DataTooLongError
	case CANNOT_BE_NULL:
		return CannotBeNullError
	default:
		return sqlError
	}
}
