package day07

import (
	"fmt"

	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/doniacld/adventofcode/puzzles/2020/07/bags"
	"github.com/doniacld/adventofcode/puzzles/solver"
)

func New(fileName string, target string) solver.Solver {
	return day07{fileName: fileName, target: target}
}

type day07 struct {
	fileName string
	target   string
}

func (d day07) Solve() (string, error) {
	r := bags.NewBags()

	err := filereader.ReadAndExtract(d.fileName, func(line string) error {
		r.ParseLine(line)
		return nil
	})

	// count the ways to achieve the root
	count := r.CountPossiblePaths(d.target)
	if err != nil {
		return "", err
	}

	nbBags := r.CountBags(d.target)
	out := fmt.Sprintf("Number of bags containing %s : %d & holds %d bags", d.target, count, nbBags)
	return out, nil
}
