package grid

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/amancooks08/game-of-life/cell"
	errors "github.com/amancooks08/game-of-life/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewGrid(t *testing.T) {
	type args struct {
		rows    int
		columns int
	}

	type test struct {
		name        string
		args        args
		expectError error
	}

	tests := []test{
		{
			name: "Grid is created with positive rows and columns.",
			args: args{
				rows:    3,
				columns: 3,
			},
			expectError: nil,
		},
		{
			name: "Grid is not created with negative rows.",
			args: args{
				rows:    -1,
				columns: 3,
			},
			expectError: errors.ErrInvalidRows,
		},
		{
			name: "Grid is not created with negative columns.",
			args: args{
				rows:    3,
				columns: -1,
			},
			expectError: errors.ErrInvalidColumns,
		},
		{
			name: "Grid is not created with zero rows.",
			args: args{
				rows:    0,
				columns: 3,
			},
			expectError: errors.ErrInvalidRows,
		},
		{
			name: "Grid is not created with zero columns.",
			args: args{
				rows:    3,
				columns: 0,
			},
			expectError: errors.ErrInvalidColumns,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewGrid(tc.args.rows, tc.args.columns)
			if err != tc.expectError {
				t.Fatalf("Expected error %v, got %v", tc.expectError, err)
			}
		})
	}

}

func TestGenerateSeed(t *testing.T) {
	t.Run("Grid is generated with seed.", func(t *testing.T) {
		deadGrid, err := NewGrid(3, 3)
		grid, err := NewGrid(3, 3)
		if err != nil {
			t.Fatalf("Expected error %v, got %v", nil, err)
		}
		grid.GenerateSeed()
		if reflect.DeepEqual(deadGrid, grid) {
			t.Fatalf("Expected grid %v, got %v", deadGrid, grid)
		}
	})
}

func TestTickNewPopulation(t *testing.T) {

	t.Run("New Population is created.", func(t *testing.T) {
		grid, err := NewGrid(3, 3)
		if err != nil {
			t.Fatalf("Expected error %v, got %v", nil, err)
		}
		grid.GenerateSeed()
		newPopulation, err := grid.TickNewPopulation()
		if err != nil {
			t.Fatalf("Expected error %v, got %v", nil, err)
		}
		if reflect.DeepEqual(grid, newPopulation) {
			t.Fatalf("Expected grid %v, got %v", grid, newPopulation)
		}
	})
}

func TestDisplayPopulation(t *testing.T) {
	grid, err := NewGrid(3, 3)
	if err != nil {
		t.Fatalf("Expected error %v, got %v", nil, err)
	}

	grid.put(cell.NewLiveCell(), 0, 0)
	grid.put(cell.NewLiveCell(), 0, 1)
	grid.put(cell.NewLiveCell(), 1, 0)

	fmt.Println(grid.DisplayPopulation())
	newGrid, err := grid.TickNewPopulation()
	if err != nil {
		t.Fatalf("Expected error %v, got %v", nil, err)
	}
	expectedGrid, err := NewGrid(3, 3)
	if err != nil {
		t.Fatalf("Expected error %v, got %v", nil, err)
	}
	fmt.Println(expectedGrid.DisplayPopulation())
	assert.Equal(t, expectedGrid, newGrid)
}
