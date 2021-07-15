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
	}{
		{"nominal case", "./resources/input-test.txt", 5, 127},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			c, err := ReadAndExtract(tc.input, tc.idxPreamble)
			assert.NoError(t, err)
			v := c.(tc.idxPreamble)
			assert.Equal(t, tc.output, v)
		})
	}
}
