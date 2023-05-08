package cell

import (
	"testing"
	assert "github.com/stretchr/testify/assert" 
)

func TestNewDeadCell(t *testing.T) {
	t.Run("Dead cell is created.", func(t *testing.T) {
		got := NewDeadCell()
		// assert false for IsAlive
		assert.False(t, got.IsAlive())
	})
}


func TestNextGeneration(t *testing.T){
	type args struct {
		neighbours int
	}

	type test struct {
		name string
		args args
		want Cell
		expected error
	}

	tests := []test{
		{
			name: "Dead cell with 3 alive neighbours becomes alive.",
			args: args{neighbours: 3},
			want: NewLiveCell(),
			expected: nil,
		},
		{
			name: "Dead cell with 2 alive neighbours stays dead.",
			args: args{neighbours: 2},
			want: NewDeadCell(),
			expected: nil,
		},
		{
			name: "Dead cell with 4 alive neighbours stays dead.",
			args: args{neighbours: 4},
			want: NewDeadCell(),
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDeadCell().NextGeneration(tt.args.neighbours)
			assert.Equal(t, tt.expected, err)
			assert.Equal(t, tt.want, got)
		})
	}
}