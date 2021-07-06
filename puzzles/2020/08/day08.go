package day08

import (
	"fmt"

	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/doniacld/adventofcode/puzzles/2020/08/game"
	"github.com/doniacld/adventofcode/puzzles/solver"
)

func New(fileName string) solver.Solver {
	return day08{fileName: fileName}
}

type day08 struct {
	fileName string
}

func (d day08) Solve() (string, error) {

	ops := game.NewOperations()
	err := filereader.ReadAndExtract(d.fileName, func(line string) error {
		err := ops.ParseLine(line)
		return err
	})
	if err != nil {
		return "", err
	}
	// store the initial ops before any modification by the Part one
	tmp := ops.ResetOps()

	_, accCounter := ops.ComputeCorruptedAcc()
	accCounterFixed := tmp.ComputeFixedAcc()
	out := fmt.Sprintf("acc value: %d, after fix acc value: %d", accCounter, accCounterFixed)

	return out, nil
}
