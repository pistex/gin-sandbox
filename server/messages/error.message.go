package messages

import (
	"errors"
	"kwanjai/interfaces"
	"net/http"
)

var ErrCredentialMismatch = errors.New("email or password is not correct")
var ErrDuplicatedEmail = errors.New("this email already register")
var ErrBadAuthorizationSession = errors.New("bad authorization session")
var ErrNonceUsedOrExpired = errors.New("nonce is used or expired")
var ErrLoadPrivateKey = errors.New("cannot load private key")

type httpError struct {
	status  int
	Message string `json:"message"`
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
