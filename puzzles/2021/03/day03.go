package day03

import (
	"fmt"

	"github.com/doniacld/adventofcode/puzzles/2021/03/binary"
	"github.com/doniacld/adventofcode/puzzles/solver"
)

func New(input string) solver.Solver {
	return day03{input}
}

// day03 is the implementation of the third day!
type day03 struct {
	fileName string
}

func (d day03) Solve() (string, error) {
	diags, err := binary.ExtractDiag(d.fileName)
	if err != nil {
		return "", err
	}
	part1, err := binary.ComputeDiags(diags)
	if err != nil {
		return "", err
	}
	part2 := 0

	return fmt.Sprintf("Part One: %d & Part Two: %d", part1, part2), nil
}
