package cell

import (
	errors "github.com/amancooks08/game-of-life/errors"
)

type DeadCell struct {
}

func (c DeadCell) IsAlive() bool {
	return false
}

func NewDeadCell() Cell {
	return DeadCell{}
}

func (c DeadCell) NextGeneration(aliveNeighbours int) (Cell, error) {
	if aliveNeighbours < 0 {
		return nil, errors.ErrInvalidNeighbours
	}
	if aliveNeighbours == 3 {
		return NewLiveCell(), nil
	}
	return NewDeadCell(), nil
}
