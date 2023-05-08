package cellAddress

import (
	"reflect"
	"testing"

	errors "github.com/amancooks08/game-of-life/errors"
)

func TestNewCellAddresss(t *testing.T){

	type args struct {
		row int
		column int
	}
	type test struct {
		name string
		args args
		want CellAddress
		expectError error
	}

	tests := []test{
		{
			name: "Cell address is created with positive values.",
			args: args{
				row: 1,
				column: 1,
			},
			want: CellAddress{
				Row: 1,
				Column: 1,
			},
			expectError: nil,
		},
		{
			name: "Cell address is not created with negative row value.",
			args: args{
				row: -1,
				column: 1,
			},
			want: CellAddress{},
			expectError: errors.ErrInvalidRowAddress,
		},
		{
			name: "Cell address is not created with negative column value.",
			args: args{
				row: 1,
				column: -1,
			},
			want: CellAddress{},
			expectError: errors.ErrInvalidColumnAddress,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T){
			got, err := NewCellAddress(tc.args.row, tc.args.column)
			if err != tc.expectError {
				t.Fatalf("Expected error %v, got %v", tc.expectError, err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}