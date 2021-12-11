package binary

import (
	"fmt"
	"math"
	"strconv"

	"github.com/doniacld/adventofcode/internal/filereader"
)

func ExtractDiag(filename string) ([][]int, error) {
	var diags [][]int
	tmp := make([]int, 0)

	err := filereader.ReadAndExtract(filename, func(line string) error {
			if len(diags) == 0 {
				diags = make([][]int, len(line))
			}

			for i := 0; i < len(line); i++ {
				v, err := strconv.Atoi(fmt.Sprintf("%c", line[i]))
				if err != nil {
					return err
				}

				tmp = diags[i]
				tmp = append(tmp, v)
				diags[i] = tmp
			}

		return nil
	})
	if err != nil {
		return diags, err
	}

	return diags, nil
}

func ComputeDiags(diags [][]int) (int, error) {
	var gamma, epsilon string

	for _, d := range diags {
		gamma += fmt.Sprintf("%d", max(d))

		eF := math.Abs(float64(max(d) - 1))
		epsilon += fmt.Sprintf("%d", int(eF))
	}

	gb, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return 0, err
	}
	ep, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return 0, err
	}

	return int(gb * ep), nil
}

func max(values []int) int {
	cnt1 := 0
	for _, v := range values {
		if v == 1 {
			cnt1++
		}
	}

	cnt0 := len(values) - cnt1

	if cnt1 > cnt0 {
		return 1
	}

	return 0
}
