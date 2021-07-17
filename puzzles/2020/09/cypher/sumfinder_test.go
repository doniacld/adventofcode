package cypher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCypher_FindSum(t *testing.T) {
	tt := []struct {
		description string
		file        string
		input       int
		output      int
	}{
		{
			"nominal case", "./resources/input-test.txt", 14, 62,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			c, err := ReadAndExtract(tc.file, 5)
			assert.NoError(t, err)
			out := c.FindSum(tc.input)
			assert.Equal(t, tc.output, out)
		})
	}

}
