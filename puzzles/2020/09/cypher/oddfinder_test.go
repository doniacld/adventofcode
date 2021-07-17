package cypher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCypher_FindOdd(t *testing.T) {
	tt := []struct {
		description string
		input       string
		idxPreamble int
		output      int
		expectedIdx int
	}{
		{"nominal case", "./resources/input-test.txt", 5, 127, 14},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			c, err := ReadAndExtract(tc.input, tc.idxPreamble)
			assert.NoError(t, err)
			v, i := c.FindOdd(tc.idxPreamble)
			assert.Equal(t, tc.output, v)
			assert.Equal(t, tc.expectedIdx, i)

		})
	}
}
