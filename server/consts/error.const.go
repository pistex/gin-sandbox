package consts

import "errors"

var CredentialMismatch = errors.New("password or email is not correct")
var DuplicatedEmail = errors.New("this email already register")
