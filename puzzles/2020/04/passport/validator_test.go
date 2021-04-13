package passport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYearRule(t *testing.T) {
	tt := []struct {
		description string
		min         int
		max         int
		input       string
		expected    bool
	}{
		{"valid case", 1920, 2002, "1950", true},
		{"equal to min limit", 1920, 2002, "1920", true},
		{"equal to max limit", 1920, 2002, "2002", true},
		{"not in range", 1920, 2002, "1902", false},
		{"invalid int", 1920, 2002, "invalidInteger", false},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			f := yearRule(tc.min, tc.max)
			assert.Equal(t, tc.expected, f(tc.input))
		})
	}
}

func TestHeightRule(t *testing.T) {
	tt := []struct {
		description string
		input       string
		expected    bool
	}{
		{"valid case cm", "155cm", true},
		{"valid case in", "59in", true},
		{"invalid case cm", "149cm", false},
		{"invalid case cm", "194cm", false},
		{"invalid case in", "58in", false},
		{"invalid case in", "77in", false},
		{"invalid case regexp match", "invalidRegexp", false},
		{"invalid unit", "77xx", false},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expected, heightRule(tc.input))
		})
	}
}

func TestRegexpRule(t *testing.T) {
	tt := []struct {
		description string
		input       string
		regexpRule  string
		expected    bool
	}{
		{"valid eye color", "amb", "amb|blu|brn|gry|grn|hzl|oth", true},
		{"invalid eye color", "xxx", "amb|blu|brn|gry|grn|hzl|oth", false},
		{"valid hair color", "#123abc", "#[0-9a-f]{6}", true},
		{"invalid hair color", "#123abz", "#[0-9a-f]{6}", false},
		{"invalid hair color", "123abc", "#[0-9a-f]{6}", false},
		{"invalid hair color", "123abc", "#[0-9a-f]{6}", false},
		{"valid pid", "000000001", "[0-9]{9}", true},
		{"invalid pid", "0123456789", "[0-9]{9}", false},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			f := regexpRule(tc.regexpRule)
			assert.Equal(t, tc.expected, f(tc.input))
		})
	}
}
