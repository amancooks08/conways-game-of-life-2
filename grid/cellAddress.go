package grid

type CellAddress struct {
	Row    uint32
	Column uint32
}

func (add CellAddress) NeighboursFor(grid Grid) []CellAddress {
	lastRowIndex, lastColumnIndex := grid.Rows()-1, grid.Columns()-1
	var neighbours []CellAddress
	for row := add.Row; row <= add.Row+1; row++ {
		for column := add.Column; column <= add.Column+1; column++ {
			if row < 0 || column < 0 || row > lastRowIndex || column > lastColumnIndex || (row == add.Row && column == add.Column) {
				continue
			}
			neighbour := CellAddress{row, column}

			neighbours = append(neighbours, neighbour)
		}
	}
	return neighbours
}
