package cell

import (
	errors "github.com/amancooks08/game-of-life/errors"
)
type LiveCell struct {
}

func (c LiveCell) IsAlive() bool {
	return true
}

func NewLiveCell() Cell {
	return LiveCell{}
}

func (c LiveCell) NextGeneration(aliveNeighbours int) (Cell, error) {
	if aliveNeighbours < 0 {
		return nil, errors.ErrInvalidNeighbours
	}
	if aliveNeighbours < 2 || aliveNeighbours > 3 {
		return NewDeadCell(), nil
	}
	return NewLiveCell(), nil
}