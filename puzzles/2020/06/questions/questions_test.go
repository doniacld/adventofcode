package questions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncCounterPartOne(t *testing.T) {
	tt := []struct{
		description string
		answers string
		expected int
	}{
		{"nominal case", "abc", 3},
		{"an answer occurs twice", "aa", 1},
		{"empty line, should not happened here", "", 0},

	}

	for _, tc := range tt {
		ys := NewYesesSum()
		t.Run(tc.description, func(t *testing.T) {
			ys.IncCounter([]byte(tc.answers))
			assert.Equal(t, tc.expected, ys.GetYeses())
		})
	}
}

func TestIncCounterPartTwo(t *testing.T) {
	tt := []struct{
		description string
		answers string
		expected int
	}{
		{"nominal case", "abc", 3},
		{"only one letter", "a", 1},
		{"empty line, should not happened here", "", 0},

	}

	for _, tc := range tt {
		ys := NewCommonAnswers()
		t.Run(tc.description, func(t *testing.T) {
			ys.IncCounter([]byte(tc.answers))
			assert.Equal(t, tc.expected, ys.GetYeses())
		})
	}
}
