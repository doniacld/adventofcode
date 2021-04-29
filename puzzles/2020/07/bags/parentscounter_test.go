package bags

import (
	"testing"

	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/stretchr/testify/assert"
)

func TestBags_CountPossiblePaths(t *testing.T) {
	tt := []struct {
		description string
		file        string
		input       string
		expected    int
	}{
		{"nominal case", "./../resources/input-test.txt", "shiny gold", 4},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			r := NewBags()
			_ = filereader.ReadAndExtract(tc.file, func(line string) error {
				r.ParseLine(line)
				return nil
			})

			out := r.CountPossiblePaths(tc.input)
			assert.Equal(t, tc.expected, out)
		})
	}
}
