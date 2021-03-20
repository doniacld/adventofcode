package day04

import (
	"fmt"
	"strings"

	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/doniacld/adventofcode/puzzles/2020/04/passport"
	"github.com/doniacld/adventofcode/puzzles/solver"
)

func New(fileName string) solver.Solver {
	return day04{fileName}
}

type day04 struct {
	fileName string
}

func (d day04) Solve() (string, error) {
	input, err := filereader.ExtractStrings(d.fileName)
	if err != nil {
		return "", err
	}

	res := retrieveValidPassports(input)
	out := fmt.Sprintf("Number of valid passports: %d", res)
	return out, nil
}

// retrieveValidPassports returns all valid passports depending on the established rules in the validator
func retrieveValidPassports(input []string) int {
	p := make(passport.Passport, 0)
	counter := 0

	for _, line := range input {
		// a blank line corresponds to a new passport
		if len(line) == 0 {
			if p.IsValid() {
				counter++
			}
			p = passport.NewPassport()
		} else {
			// split the passport fields line by spaces
			passportFields := strings.Split(line, " ")
			for _, field := range passportFields {
				// split the field into key, value of the field
				f := strings.Split(field, ":")
				p[f[0]] = f[1]
			}
		}
	}
	if p.IsValid() {
		counter++
	}
	return counter
}
