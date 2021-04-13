package day01

import (
	"fmt"
	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/doniacld/adventofcode/puzzles/2020/01/targetsum"
	"github.com/doniacld/adventofcode/puzzles/solver"
)

const targetSum = 2020

func New(input string) solver.Solver {
	return day01{input}
}

// day01 is the implementation of the first day!
type day01 struct {
	fileName string
}

// Solve computes the day01 puzzle
func (d day01) Solve() (string, error) {
	expenses, err := filereader.ExtractInt(d.fileName)
	if err != nil {
		return "", err
	}
	part1 := targetsum.ComputeTargetSumWithPair(expenses, targetSum)
	part2 := targetsum.ComputeTargetSumWithTriplet(expenses, targetSum)

	return fmt.Sprintf("Part One: %d & Part Two: %d", part1, part2), nil
}
