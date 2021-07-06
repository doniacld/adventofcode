package game

import (
	"testing"

	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/stretchr/testify/assert"
)

func TestOperations_ComputeCorruptedAcc(t *testing.T) {
	tt := []struct {
		description string
		input       string
		output      int
	}{
		{"nominal case", "./resources/input-test.txt", 8},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			ops := NewOperations()
			err := filereader.ReadAndExtract(tc.input, func(line string) error {
				err := ops.ParseLine(line)
				return err
			})
			assert.NoError(t, err)
			acc := ops.ComputeFixedAcc()
			assert.Equal(t, tc.output, acc)
		})
	}
}
