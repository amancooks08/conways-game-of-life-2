package grid

import (
	rand "math/rand"

	cell "github.com/amancooks08/game-of-life/cell"
	errors "github.com/amancooks08/game-of-life/errors"
)

type Grid struct {
	cells [][]cell.Cell
}

func NewGrid(rows int, columns int) (Grid, error) {
	if rows < 0 || rows == 0 {
		return Grid{}, errors.ErrInvalidRows
	}
	if columns < 0 || columns == 0 {
		return Grid{}, errors.ErrInvalidColumns
	}
	return Grid{
		cells: make([][]cell.Cell, rows),
	}, nil
}

func (g Grid) GenerateSeed() {
	for row := 0; row < len(g.cells); row++ {
		for column := 0; column < len(g.cells[0]); column++ {
			if rand.Intn(2) == 0 {
				g.cells[row][column] = cell.NewDeadCell()
			} else {
				g.cells[row][column] = cell.NewLiveCell()
			}
		}
	}
}
