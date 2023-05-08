package cellAddress

import(
	errors "github.com/amancooks08/game-of-life/errors"
)

type CellAddress struct {
	Row    int
	Column int
}

func NewCellAddress(row int, column int) (CellAddress, error) {
	if row < 0 {
		return CellAddress{}, errors.ErrInvalidRowAddress
	}
	if column < 0 {
		return CellAddress{}, errors.ErrInvalidColumnAddress
	}
	return CellAddress{
		Row:    row,
		Column: column,
	}, nil
}
