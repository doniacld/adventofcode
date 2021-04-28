package day04

import (
	"fmt"
	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidPassports(t *testing.T) {
	tt := []struct {
		description string
		input       string
		expected    int
	}{
		{"nominal case", "./input-test.txt", 4},
	}
	for _, tc := range tt {
		in, err := filereader.ExtractString(tc.input)
		if err != nil {
			fmt.Println(err)
		}
		t.Run(tc.description, func(t *testing.T) {
			res := retrieveValidPassports(in)
			assert.Equal(t, tc.expected, res)
		})
	}
}
