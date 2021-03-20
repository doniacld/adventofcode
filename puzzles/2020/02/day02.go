package day02

import (
	"fmt"
	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/doniacld/adventofcode/puzzles/2020/02/policypasswd"
	"github.com/doniacld/adventofcode/puzzles/solver"
)

func New(input string) solver.Solver {
	return day02{input}
}

// day02 is the implementation of the first day!
type day02 struct {
	fileName string
}

func (d day02) Solve() (string, error) {
	lines, err := filereader.ExtractStrings(d.fileName)
	if err != nil {
		return "", err
	}

	occ, err := policypasswd.RetrieveValidPasswordsOccurrencePolicy(lines)
	if err != nil {
		return "", err
	}
	pos, err := policypasswd.RetrieveValidPasswordsPositionPolicy(lines)
	if err != nil {
		return "", err
	}

	out := fmt.Sprintf("Number of valid passwords for occurrency policy: %d & position policy: %d", occ, pos)
	return out, nil
}
