package day05

import (
	"fmt"
	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay05_retrieveMaxAndMissingSeatIds(t *testing.T) {
	seats, err := filereader.ExtractStrings("./input-test.txt")
	if err != nil {
		fmt.Println(err)
	}

	maxSeatId, missingSeatId := retrieveMaxAndMissingSeatIds(seats)
	assert.Equal(t, 820, maxSeatId)
	assert.Equal(t, 168, missingSeatId)
}
