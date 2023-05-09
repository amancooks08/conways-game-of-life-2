package errors

import "errors"

var (
	ErrInvalidNeighbours	= errors.New("invalid number of neighbours")
	ErrInvalidRows			= errors.New("invalid number of rows")
	ErrInvalidColumns		= errors.New("invalid number of columns")
	ErrInvalidCellAddress	= errors.New("invalid cell address")
)