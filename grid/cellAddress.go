package grid

type CellAddress struct {
	Row    int
	Column int
}

func (add CellAddress) NeighboursFor(grid Grid) []CellAddress {
	lastRowIndex, lastColumnIndex := grid.Rows()-1, grid.Columns()-1
	var neighbours []CellAddress
	for row := add.Row - 1; row <= add.Row+1; row++ {
		for column := add.Column - 1; column <= add.Column+1; column++ {
			if row < 0 || column < 0 || row > lastRowIndex || column > lastColumnIndex || (row == add.Row && column == add.Column) {
				continue
			}
			neighbour := CellAddress{row, column}
			neighbours = append(neighbours, neighbour)
		}
	}
	return neighbours
}
