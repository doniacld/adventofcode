package day03

import (
	"fmt"
	"github.com/doniacld/adventofcode/internal/filereader"
	"github.com/doniacld/adventofcode/puzzles/solver"
)

const (
	// tree is the char representing a tree in the map
	tree = "#"
)

func New(fileName string, slopePattern [][2]int) solver.Solver {
	return day03{fileName, slopePattern}
}

type day03 struct {
	fileName     string
	slopePattern [][2]int
}

func (d day03) Solve() (string, error) {
	res, err := GetEncounteredTrees(d.fileName, d.slopePattern)
	if err != nil {
		return "", err
	}
	out := fmt.Sprintf("Number of encountered trees: %d for pattern: %d", res, d.slopePattern)
	return out, nil
}

// GetEncounteredTrees returns the multiplication of the encountered trees per slope pattern
// complexity time: O(n²)
func GetEncounteredTrees(file string, slopes [][2]int) (int, error) {
	input, err := filereader.ExtractStrings(file, "")
	if err != nil {
		return 0, err
	}
	result := 1
	for _, s := range slopes {
		result *= countTrees(input, s[0], s[1])
	}
	return result, nil
}

// countTrees counts the encountered trees for each given slope pattern
func countTrees(input [][]string, right int, down int) int {
	var trees, line, col int

	for line = 0; line < len(input); line +=down {
		if input[line][col] == tree {
			trees++
		}
		col = (col + right) % len(input[line])
	}
	return trees
}
