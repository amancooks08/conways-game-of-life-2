package errors

import "errors"

var (
	ErrInvalidRowAddress    = errors.New("invalid row address")
	ErrInvalidColumnAddress = errors.New("invalid column address")
	
)