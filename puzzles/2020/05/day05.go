package day05

import (
	"fmt"

	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/doniacld/adventofcode/puzzles/2020/05/seat"
	"github.com/doniacld/adventofcode/puzzles/solver"
)

func New(fileName string) solver.Solver {
	return day05{fileName: fileName}
}

type day05 struct {
	fileName string
}

func (d day05) Solve() (string, error) {
	seats, err := filereader.ExtractStrings(d.fileName)
	if err != nil {
		return "", err
	}

	maxSeatId, missingSeatId := retrieveMaxAndMissingSeatIds(seats)
	out := fmt.Sprintf("max seatId: %d & missing seatId: %d", maxSeatId, missingSeatId)
	return out, nil
}

// retrieveMaxAndMissingSeatIds browses all seatsId and retrieves the max seatId value
// and the missing seatId
func retrieveMaxAndMissingSeatIds(seats []string) (int, int) {
	var max, seatId int
	seenSeatIds := make(map[int]struct{}, 0)

	for _, s := range seats {
		bp := seat.RetrieveBoardingPass(s)
		seatId = seat.ComputeSeatId(bp)

		// fill the map with all seatIds
		seenSeatIds[seatId] = struct{}{}

		// update the max
		max = seat.MaxSeatId(max, seatId)
	}

	// retrieve the missingSeatId
	allSeatIds := seat.CreateAllSeatIds()
	missingSeatId := seat.GetMissingSeatId(allSeatIds, seenSeatIds)

	return max, missingSeatId
}
