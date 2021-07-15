package day09

import (
	"fmt"

	"github.com/doniacld/adventofcode/puzzles/2020/09/cypher"
	"github.com/doniacld/adventofcode/puzzles/solver"
)

func New(fileName string, idxPreamble int) solver.Solver {
	return day09{fileName: fileName, idxPreamble: idxPreamble}
}

type day09 struct {
	fileName    string
	idxPreamble int
}

func (d day09) Solve() (string, error) {
	c, err := cypher.ReadAndExtract(d.fileName, d.idxPreamble)
	if err != nil {
	    return err
	}

	noSumValue := c.FindIntruder(d.idxPreamble)
	out := fmt.Sprintf("First number to not respect the rule is: %d ", noSumValue)
	return out, nil
}
