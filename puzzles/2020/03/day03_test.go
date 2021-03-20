package day03

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRetrieveTreesNumb(t *testing.T) {
	tt := []struct {
		description string
		input       string
		slopes      [][2]int
		expectedRes int
	}{
		{"Nominal case", "./input-test.txt", [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}, 336},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res, err := GetEncounteredTrees(tc.input, tc.slopes)
			if err != nil {
				fmt.Println(err)
			}
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
