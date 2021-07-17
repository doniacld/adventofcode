package game

import (
	"testing"

	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/stretchr/testify/assert"
)

func TestOperations_ComputeOperations(t *testing.T) {
	tt := []struct {
		description string
		input       string
		output      int
	}{
		{"nominal case", "./resources/input-test.txt", 5},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			var ops Operations
			err := filereader.ReadAndExtract(tc.input, func(line string) error {
				err := ops.ParseLine(line)
				return err
			})
			assert.NoError(t, err)
			idx, acc := ops.ComputeCorruptedAcc()
			assert.Equal(t, tc.output, acc)
			assert.Equal(t, 1, idx)
		})
	}
}
