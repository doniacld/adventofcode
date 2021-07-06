package bags

import (
	"testing"

	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/stretchr/testify/assert"
)

func TestBags_CountBags(t *testing.T) {
	tt := []struct {
		description string
		file        string
		input       string
		expected    int
	}{
		{"several branches case", "./../resources/input-test.txt", "shiny gold", 32},
		{"linear case", "./../resources/input-test-part2.txt", "shiny gold", 126},
		{"does not exist", "./../resources/input-test-part2.txt", "bright gold", 0},

	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			r := NewBags()
			_ = filereader.ReadAndExtract(tc.file, func(line string) error {
				r.ParseLine(line)
				return nil
			})

			out := r.CountBags(tc.input)
			assert.Equal(t, tc.expected, out)
		})
	}
}
