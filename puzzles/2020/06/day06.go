package day06

import (
	"fmt"

	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/doniacld/adventofcode/puzzles/2020/06/questions"
	"github.com/doniacld/adventofcode/puzzles/solver"
)

func New(fileName string) solver.Solver {
	return day06{fileName: fileName}
}

type day06 struct {
	fileName string
}

func (d day06) Solve() (string, error) {
	var yesCounter, caCounter int
	ys := questions.NewYesesSum()
	ca := questions.NewCommonAnswers()

	err := filereader.ReadAndExtract(d.fileName, func(line string) error {
		// a blank line means it is a new group, reset values
		if line == "" {
			yesCounter += ys.GetYeses()
			ys.Reset()

			caCounter += ca.GetYeses()
			ca.Reset()
			return nil
		}
		// fill the answers and increment the counter
		l := []byte(line)
		ys.IncCounter(l)
		ca.IncCounter(l)
		return nil
	})
	if err != nil {
		return "", err
	}
	// handle the last value
	yesCounter += ys.GetYeses()
	caCounter += ca.GetYeses()

	out := fmt.Sprintf("Yes Sum: %d & Common number answers per group: %d", yesCounter, caCounter)
	return out, nil
}
