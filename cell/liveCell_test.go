package cell

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestNewLiveCell(t *testing.T) {

	t.Run("Live cell is created.", func(t *testing.T) {
		got := NewLiveCell()
		// assert true for IsAlive

		assert.True(t, got.IsAlive())
	})
}

func TestNextGenerationOfLiveCell(t *testing.T) {
	type args struct {
		neighbours int
	}

	type test struct {
		name     string
		args     args
		want     Cell
		expected error
	}

	tests := []test{
		{
			name:     "Live cell with 2 alive neighbours stays alive.",
			args:     args{neighbours: 2},
			want:     NewLiveCell(),
			expected: nil,
		},
		{
			name:     "Live cell with 3 alive neighbours stays alive.",
			args:     args{neighbours: 3},
			want:     NewLiveCell(),
			expected: nil,
		},
		{
			name:     "Live cell with less than 2 alive neighbours dies.",
			args:     args{neighbours: 1},
			want:     NewDeadCell(),
			expected: nil,
		},
		{
			name:     "Live cell with more than 3 alive neighbours dies.",
			args:     args{neighbours: 4},
			want:     NewDeadCell(),
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLiveCell().NextGeneration(tt.args.neighbours)
			assert.Equal(t, tt.expected, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
