package cypher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAndExtract(t *testing.T) {
	tt := []struct {
		description string
		input       string
		idxPreamble int
		output      Cypher
	}{
		{"nominal case", "./resources/input-test-small.txt", 2,
			Cypher{Preamble: Preamble{1: {}, 2: {}}, Numbers: []int{1, 2, 3, 4}},
		}}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			c, err := ReadAndExtract(tc.input, tc.idxPreamble)
			assert.NoError(t, err)
			assert.Equal(t, tc.output, c)
		})
	}
}
