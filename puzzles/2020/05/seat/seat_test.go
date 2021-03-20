package seat

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeSeatId(t *testing.T) {
	tt := []struct {
		description string
		seatCode    string
		expected    int
	}{

		{"valid case", "FBFBBFFRLR", 357},
		{"valid case", "BFFFBBFRRR", 567},
		{"valid case", "FFFBBBFRRR", 119},
		{"valid case", "BBFFBBFRLL", 820},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			bp := RetrieveBoardingPass(tc.seatCode)
			seatId := ComputeSeatId(bp)
			assert.Equal(t, tc.expected, seatId)
		})
	}
}

func TestDecodeSeat_Row(t *testing.T) {
	tt := []struct {
		description string
		seatCode    string
		expected    int
	}{
		{"valid case row", "FBFBBFF", 44},
		 {"valid case row", "BFFFBBF", 70},
		 {"valid case row", "FFFBBBF", 14},
		{"valid case row", "BBFFBBF", 102},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := decodeSeat("F", "B", strings.Split(tc.seatCode, ""), 0, 127, 0)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestDecodeSeat_Column(t *testing.T) {
	tt := []struct {
		description string
		seatCode    string
		expected    int
	}{
		{"valid case column", "RLR", 5},
		{"valid case column", "RRR", 7},
		{"valid case column", "RLL", 4},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := decodeSeat("L", "R", strings.Split(tc.seatCode, ""), 0, 7, 0)
			assert.Equal(t, tc.expected, res)
		})
	}
}
