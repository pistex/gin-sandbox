package helpers

import (
	"kwanjai/interfaces"
	"log"
	"net/http"
)

type httpError struct {
	status  int
	Message string `json:"message,"`
}

func (e *httpError) GetStatus() int {
	return e.status
}

func (e *httpError) GetJSON() interface{} {
	return e
}

func NewBadRequestError(err error) interfaces.IHTTPError {
	return &httpError{
		status:  http.StatusBadRequest,
		Message: err.Error(),
	}
}

func NewInternalServerError(err error) interfaces.IHTTPError {
	return &httpError{
		status:  http.StatusInternalServerError,
		Message: err.Error(),
	}
}

func CheckErrorAndPanic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
