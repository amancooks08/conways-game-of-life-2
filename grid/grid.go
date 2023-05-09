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
	grid := make([][]cell.Cell, rows)
	for row := 0; row < rows; row++ {
		grid[row] = make([]cell.Cell, columns)
	}
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			grid[row][column] = cell.NewDeadCell()
		}
	}
	return Grid{
		cells: grid,
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

func (g Grid) TickNewPopulation() (Grid, error) {
	grid, err := NewGrid(int(g.Rows()), int(g.Columns()))
	if err != nil {
		return Grid{}, err
	}
	var row, column int
	for row = 0; row < int(len(g.cells)); row++ {
		for column = 0; column < int(len(g.cells[0])); column++ {
			cell := g.cells[row][column]
			cellNeighbours := CellAddress{row, column}.NeighboursFor(g)
			aliveNeighbours := grid.countAliveNeighbours(cellNeighbours)
			cell, err := g.cells[row][column].NextGeneration(aliveNeighbours)
			if err != nil {
				return Grid{}, err
			}
			grid.cells[row][column] = cell
		}
	}
	return grid, nil
}

func (g Grid) Rows() int {
	return int(len(g.cells))
}

func (g Grid) Columns() int {
	return int(len(g.cells[0]))
}

func (g Grid) countAliveNeighbours(cellNeighbours []CellAddress) int {
	var aliveNeighbours int
	for _, neighbour := range cellNeighbours {
		if g.cells[neighbour.Row][neighbour.Column] != nil && g.cells[neighbour.Row][neighbour.Column].IsAlive() {
			aliveNeighbours++
		}
	}
	return aliveNeighbours
}

func (g Grid) put(cell cell.Cell, row int, column int) error {
	if row < 0 || column < 0 || row >= g.Rows() || column >= g.Columns() {
		return errors.ErrInvalidCellAddress
	}
	g.cells[row][column] = cell
	return nil
}

func (g Grid) DisplayPopulation() string {
	var display string
	for row := 0; row < int(len(g.cells)); row++ {
		for column := 0; column < int(len(g.cells[0])); column++ {
			if g.cells[row][column].IsAlive() {
				display += "O "
			} else {
				display += "* "
			}
		}
		display += "\n"
	}
	return display
}
