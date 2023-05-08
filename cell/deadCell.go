package cell

import (
	errors "github.com/amancooks08/game-of-life/errors"
)

type DeadCell struct {
}

func (c *DeadCell) IsAlive() bool {
	return false
}

func NewDeadCell() Cell {
	return &DeadCell{}
}

func (c *DeadCell) NextGeneration(neighbours int) (Cell, error) {
	if neighbours <= 0 {
		return nil, errors.ErrInvalidNeighbours
	}
	if neighbours == 3 {
		return NewLiveCell(), nil
	}
	return NewDeadCell(), nil
}
