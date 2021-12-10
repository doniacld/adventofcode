package day02

import (
	"fmt"

	"github.com/doniacld/adventofcode/puzzles/2021/02/dive"
	"github.com/doniacld/adventofcode/puzzles/solver"
)

func New(input string) solver.Solver {
	return day02{input}
}

// day02 is the implementation of the second day!
type day02 struct {
	fileName string
}

// Solve computes the second puzzle
func (d day02) Solve() (string, error) {
	moves, err := dive.ExtractMove(d.fileName)
	if err != nil {
		return "", err
	}

	part1 := dive.ComputePosition(moves)
	part2 := dive.ComputePositionWithAim(moves)

	return fmt.Sprintf("Part One: %d & Part Two: %d", part1, part2), nil
}

