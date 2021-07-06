package game

import (
	"testing"

	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/stretchr/testify/assert"
)

func TestOperations_ParseLine(t *testing.T) {
	tt := []struct {
		description string
		input       string
		output      Operations
	}{
		{"nominal case", "./resources/input-test-small.txt",
			Operations{{"nop", +0, false}, {"acc", +1, false}, {"jmp", -4, false}}},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			ops := NewOperations()
			err := filereader.ReadAndExtract(tc.input, func(line string) error {
				err := ops.ParseLine(line)
				return err
			})
			assert.NoError(t, err)
			assert.Equal(t, tc.output, ops)
		})
	}
}
