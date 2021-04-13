package targetsum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeTargetSumWithPair(t *testing.T) {
	tt := []struct {
		description string
		value       []int
		expectedRes int
	}{
		{"Given use case", []int{1721, 979, 366, 299, 675, 1456}, 514579,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := ComputeTargetSumWithPair(tc.value, 2020)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}

func TestComputeTargetSumWithTriplet(t *testing.T) {
	tt := []struct {
		description string
		value       []int
		expectedRes int
	}{
		{"Given use case", []int{1721, 979, 366, 299, 675, 1456}, 241861950,
		},
	}
	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := ComputeTargetSumWithTriplet(tc.value, 2020)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
