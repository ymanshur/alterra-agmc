package constant

import "errors"

var (
	ErrInvalidUrlParam    = errors.New("invalid id url parameter")
	ErrCredentialNotMatch = errors.New("these credentials do not match our records")
	ErrReqEmailPassword   = errors.New("required both email and password")
)
