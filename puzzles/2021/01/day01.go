package _1

import (
	"fmt"
	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/doniacld/adventofcode/puzzles/2021/01/depth"

	"github.com/doniacld/adventofcode/puzzles/solver"
)

func New(input string) solver.Solver {
	return day01{input}
}

// day01 is the implementation of the first day!
type day01 struct {
	fileName string
}

// Solve computes the day01 puzzle
func (d day01) Solve() (string, error) {
	measurements, err := filereader.ExtractInt(d.fileName)
	if err != nil {
		return "", err
	}

	part1 := depth.ComputeIncreaseDepth(measurements)
	part2 := depth.ComputeIncreaseDepthSlidingWindow(measurements)
	return fmt.Sprintf("Part One: %d & Part Two: %d", part1, part2), nil
}