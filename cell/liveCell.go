package cell

import (
	errors "github.com/amancooks08/game-of-life/errors"
)
type LiveCell struct {
}

func (c *LiveCell) IsAlive() bool {
	return true
}

func NewLiveCell() Cell {
	return &LiveCell{}
}

func (c *LiveCell) NextGeneration(neighbours int) (Cell, error) {
	if neighbours <= 0 {
		return nil, errors.ErrInvalidNeighbours
	}
	if neighbours < 2 || neighbours > 3 {
		return NewDeadCell(), nil
	}
	return NewLiveCell(), nil
}