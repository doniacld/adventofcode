package dive

import (
	"strconv"
	"strings"

	"github.com/doniacld/adventofcode/internal/filereader"
)

type Move struct {
	op string
	nb int
}

func ExtractMove(filename string) ([]Move, error) {
	moves := make([]Move, 0)

	err := filereader.ReadAndExtract(filename, func(line string) error {
		l := strings.Split(line, " ")

		nb, err := strconv.Atoi(l[1])
		if err != nil {
			return err
		}

		move := Move{op: l[0], nb: nb}
		moves = append(moves, move)

		return nil
	})
	if err != nil {
		return moves, err
	}
	return moves, nil
}

func ComputePosition(moves []Move) int {
	x, y := 0, 0
	for _, m := range moves {
		switch m.op {
		case "forward":
			x += m.nb
		case "down":
			y += m.nb
		case "up":
			y -= m.nb
		}
	}
	return x * y
}

func ComputePositionWithAim(moves []Move) int {
	x, y, aim := 0, 0, 0
	for _, m := range moves {
		switch m.op {
		case "forward":
			x += m.nb
			y += m.nb * aim
		case "down":
			aim += m.nb
		case "up":
			aim -= m.nb
		}
	}
	return x*y
}
