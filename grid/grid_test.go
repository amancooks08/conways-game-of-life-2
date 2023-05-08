package grid

import (
	"testing"
	errors "github.com/amancooks08/game-of-life/errors"
)

func TestNewGrid(t *testing.T) {
	type args struct{
		rows int
		columns int
	}

	type test struct {
		name string
		args args
		expectError error
	}

	tests := []test{
		{
			name: "Grid is created with positive rows and columns.",
			args: args{
				rows: 3,
				columns: 3,
			},
			expectError: nil,
		},
		{
			name: "Grid is not created with negative rows.",
			args: args{
				rows: -1,
				columns: 3,
			},
			expectError: errors.ErrInvalidRows,
		},
		{
			name: "Grid is not created with negative columns.",
			args: args{
				rows: 3,
				columns: -1,
			},
			expectError: errors.ErrInvalidColumns,
		},
		{
			name: "Grid is not created with zero rows.",
			args: args{
				rows: 0,
				columns: 3,
			},
			expectError: errors.ErrInvalidRows,
		},
		{
			name: "Grid is not created with zero columns.",
			args: args{
				rows: 3,
				columns: 0,
			},
			expectError: errors.ErrInvalidColumns,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T){
			_, err := NewGrid(tc.args.rows, tc.args.columns)
			if err != tc.expectError {
				t.Fatalf("Expected error %v, got %v", tc.expectError, err)
			}
		})
	}

}