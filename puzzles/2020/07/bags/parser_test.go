package bags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRule(t *testing.T) {
	tt := []struct {
		description string
		input       string
		expected    Bags
	}{
		{"nominal case",
			"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			map[string]bagFamily{
				"light red":    {[]bag(nil), []bag{{1, "bright white"}, {2, "muted yellow"}}},
				"bright white": {[]bag{{1, "light red"}}, []bag(nil)},
				"muted yellow": {[]bag{{2, "light red"}}, []bag(nil)}},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			r := NewBags()
			r.ParseLine(tc.input)
			assert.Equal(t, tc.expected, r)
		})
	}
}
