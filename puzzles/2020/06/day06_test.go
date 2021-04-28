package day06

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tt := []struct {
		description           string
		input                 string
		expectedYesSum        int
		expectedCommonAnswers int
		expectedErr           string
	}{
		{"small sample", "input-test.txt", 11, 6, ""},
		{"large sample", "input.txt", 6437, 3229, ""},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			d := New(tc.input)

			o, err := d.Solve()
			assert.Nil(t, err)
			assert.Equal(t, fmt.Sprintf("Yes Sum: %d & Common number answers per group: %d", tc.expectedYesSum, tc.expectedCommonAnswers), o)
		})
	}
}
