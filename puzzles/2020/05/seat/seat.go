package seat

import (
	"strings"
)

const (
	minRow      = 1
	maxRow      = 128
	minRowLimit = 20
	maxRowLimit = 120

	minCol      = 1
	maxCol      = 8
	minColLimit = 1
	maxColLimit = 7
)

// boardingPass is represented by a row and a column of the plane
type boardingPass struct {
	row    int
	column int
}

// RetrieveBoardingPass returns the row and columns composing a boardingPass from the seat code
func RetrieveBoardingPass(seatCode string) boardingPass {
	seat := strings.Split(seatCode, "")

	row := decodeSeat("F", "B", seat[:7], minRow-1, maxRow-1, 0)
	column := decodeSeat("L", "R", seat[7:], minCol-1, maxCol-1, 0)
	return boardingPass{row, column}
}

// CreateAllSeatIds returns an array containing all the computed seatIds
func CreateAllSeatIds() []int {
	seatIds := make([]int, 0)
	// retrieves the seatId for the min and max of the array
	min := ComputeSeatId(boardingPass{minRow, minCol})
	max := ComputeSeatId(boardingPass{maxRow, maxCol})

	for counter := min; counter <= max; counter++ {
		seatIds = append(seatIds, counter)
	}
	return seatIds
}

// GetMissingSeatId returns the missing seatId
func GetMissingSeatId(allSeatIds []int, seenSeatIds map[int]struct{}) int {
	for _, sid := range allSeatIds {
		if _, ok := seenSeatIds[sid]; !ok {

			// check if the seatId is in the middle of the aircraft
			if isInMiddle(sid) {
				return sid
			}
		}
	}
	return 0
}

// ComputeSeatId computes the seatId from the boarding pass values with the given rule
func ComputeSeatId(bp boardingPass) int {
	return bp.row*8 + bp.column
}

// MaxSeatId returns the maximum value
func MaxSeatId(value1, value2 int) int {
	if value1 <= value2 {
		return value2
	}
	return value1
}

// decodeSeat computes the min or max range at each step depending on the limits
func decodeSeat(left, right string, letter []string, min, max, i int) int {
	if i == len(letter) {
		return min
	}
	// update the max value with the middle between the current min and max
	if letter[i] == left {
		max = max - (max-min)/2 - 1
	}

	// update the min value with the middle between the current min and max
	if letter[i] == right {
		min = min + (max-min)/2 + 1
	}
	i++
	return decodeSeat(left, right, letter, min, max, i)
}


// isInMiddle returns false if a seatId is at the very front or very back on the aircraft
func isInMiddle(seatId int) bool {
	minLimit := ComputeSeatId(boardingPass{minRowLimit, maxColLimit})
	maxLimit := ComputeSeatId(boardingPass{maxRowLimit, minColLimit})

	return seatId > minLimit && seatId < maxLimit
}
