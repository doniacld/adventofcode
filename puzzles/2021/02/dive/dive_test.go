package dive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputePosition(t *testing.T) {
	tests := []struct {
		name string
		file string
		want int
	}{
		{"nominal case", "./input-test.txt", 150},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moves, err := ExtractMove(tt.file)
			assert.Nil(t, err)
			if got := ComputePosition(moves); got != tt.want {
				t.Errorf("ComputePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputePositionWithAim(t *testing.T) {
	tests := []struct {
		name string
		file string
		want int
	}{
		{"nominal case", "./input-test.txt", 900},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moves, err := ExtractMove(tt.file)
			assert.Nil(t, err)
			if got := ComputePositionWithAim(moves); got != tt.want {
				t.Errorf("ComputePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}