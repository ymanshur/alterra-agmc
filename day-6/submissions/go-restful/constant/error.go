package constant

import "errors"

var (
	ErrInvalidUrlParam = errors.New("invalid id url parameter")
	ErrInvalidInput    = errors.New("invalid input")
	ErrRecordNotFound  = errors.New("record not found")
)
