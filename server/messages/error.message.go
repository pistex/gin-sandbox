package messages

import "errors"

var ErrCredentialMismatch = errors.New("password or email is not correct")
var ErrDuplicatedEmail = errors.New("this email already register")
