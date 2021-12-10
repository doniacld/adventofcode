package depth

import (
	"testing"

	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/stretchr/testify/assert"
)


func TestComputeIncreaseDepth(t *testing.T) {
	tests := []struct {
		name string
		file string
		want int
	}{
		{"nominal case", "./input-test.txt", 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			measurements, err := filereader.ExtractInt(tt.file)
			assert.Nil(t, err)
			if got := ComputeIncreaseDepth(measurements); got != tt.want {
				t.Errorf("ComputeIncreaseDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeIncreaseDepthSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		file string
		want int
	}{
		{"nominal case", "./input-test.txt", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			measurements, err := filereader.ExtractInt(tt.file)
			assert.Nil(t, err)
			if got := ComputeIncreaseDepthSlidingWindow(measurements); got != tt.want {
				t.Errorf("ComputeIncreaseDepthSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}