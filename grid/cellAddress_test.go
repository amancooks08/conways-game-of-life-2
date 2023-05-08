package grid

import (
	"reflect"
	"testing"
)

func TestNewCellAddresss(t *testing.T) {

	type args struct {
		row    uint32
		column uint32
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
