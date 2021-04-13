package policypasswd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveValidPasswordsOccurrencePolicy(t *testing.T) {
	tt := []struct {
		description string
		input       []string
		expectedRes int
	}{
		{"Given use case", []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}, 2},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res, err := RetrieveValidPasswordsOccurrencePolicy(tc.input)
			if err != nil {
				fmt.Println(err)
			}
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}

func TestRetrieveValidPasswordsPositionPolicy(t *testing.T) {
	tt := []struct {
		description string
		input       []string
		expectedRes int
	}{
		{"Given use case", []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}, 1},

	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res, err := RetrieveValidPasswordsPositionPolicy(tc.input)
			if err != nil {
				fmt.Println(err)
			}
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
