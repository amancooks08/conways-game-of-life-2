package grid

import (
	"reflect"
	"testing"

	"github.com/amancooks08/game-of-life/cell"
)

func TestNewCellAddresss(t *testing.T) {

	type args struct {
		row    int
		column int
	}
	type test struct {
		name        string
		args        args
		want        CellAddress
		expectError error
	}

	tests := []test{
		{
			name: "Cell address is created with positive values.",
			args: args{
				row:    1,
				column: 1,
			},
			want: CellAddress{
				Row:    1,
				Column: 1,
			},
			expectError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CellAddress{tc.args.row, tc.args.column}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestNeighboursFor(t *testing.T) {
	type args struct {
		grid Grid
	}
	type test struct {
		name string
		cell CellAddress
		args args
		want []CellAddress
	}
	tests := []test{
		{
			name: "Test neighbours for center cell to be eight.",
			cell: CellAddress{1, 1},
			args: args{
				grid: Grid{
					cells: [][]cell.Cell{
						{cell.NewLiveCell(), cell.NewLiveCell(), cell.NewLiveCell()},
						{cell.NewLiveCell(), cell.NewLiveCell(), cell.NewLiveCell()},
						{cell.NewLiveCell(), cell.NewLiveCell(), cell.NewLiveCell()},
					},
				},
			},
			want: []CellAddress{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		},
		{
			name: "Test neighbours for corner cell to be three.",
			cell: CellAddress{0, 0},
			args: args{
				grid: Grid{
					cells: [][]cell.Cell{
						{cell.NewLiveCell(), cell.NewLiveCell(), cell.NewLiveCell()},
						{cell.NewLiveCell(), cell.NewLiveCell(), cell.NewLiveCell()},
						{cell.NewLiveCell(), cell.NewLiveCell(), cell.NewLiveCell()},
					},
				},
			},
			want: []CellAddress{{0, 1}, {1, 0}, {1, 1}},
		},
		{
			name: "Test neighbours for edge cell to be five.",
			cell: CellAddress{0, 1},
			args: args{
				grid: Grid{
					cells: [][]cell.Cell{
						{cell.NewLiveCell(), cell.NewLiveCell(), cell.NewLiveCell()},
						{cell.NewLiveCell(), cell.NewLiveCell(), cell.NewLiveCell()},
						{cell.NewLiveCell(), cell.NewLiveCell(), cell.NewLiveCell()},
					},
				},
			},
			want: []CellAddress{{0, 0}, {0, 2}, {1, 0}, {1, 1}, {1, 2}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.cell.NeighboursFor(tc.args.grid)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}
