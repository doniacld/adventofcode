package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/doniacld/adventofcode/puzzles/solver"

	day0120 "github.com/doniacld/adventofcode/puzzles/2020/01"
	day0220 "github.com/doniacld/adventofcode/puzzles/2020/02"
	day0320 "github.com/doniacld/adventofcode/puzzles/2020/03"
	day0420 "github.com/doniacld/adventofcode/puzzles/2020/04"
	day0520 "github.com/doniacld/adventofcode/puzzles/2020/05"
	day0620 "github.com/doniacld/adventofcode/puzzles/2020/06"

	day0121 "github.com/doniacld/adventofcode/puzzles/2021/01"

)

const (
	flagParamYear = "year"
	flagParamDay  = "day"

	defaultYear = 2021

	defaultDay = 01
	minDay     = 1
	maxDay     = 25
)

func main() {
	// parse flags
	year := flag.Int(flagParamYear, defaultYear, "give me a year")
	day := flag.Int(flagParamDay, defaultDay, "give me a day")
	flag.Parse()

	if *year < 2020 {
		log.Fatalf("Maybe one day, I will solve the previous years !")
	}

	if *day < minDay || *day > maxDay {
		log.Fatalf("invalid given day: %q is out of the possible range [1-25]", *day)
	}

	var s solver.Solver

	if *year == 2020 {
		switch *day {
		case 1:
			s = day0120.New("./puzzles/2020/01/input.txt")
		case 2:
			s = day0220.New("./puzzles/2020/02/input.txt")
		case 3:
			s = day0320.New("./puzzles/2020/03/input.txt", [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}})
		case 4:
			s = day0420.New("./puzzles/2020/04/input.txt")
		case 5:
			s = day0520.New("./puzzles/2020/05/input.txt")
		case 6:
			s = day0620.New("./puzzles/2020/06/input.txt")
		}
	}

	if * year == 2021 {
		switch *day {
		case 1:
			s = day0121.New("./puzzles/2021/01/input.txt")
		}
	}

	out, err := s.Solve()
	if err != nil {
		log.Fatalf("Something wrong happened while computing day %d: %s", *day, err.Error())
	}

	fmt.Printf(">>>>>>>>>>>>>> %s <<<<<<<<<<<<<<<\n", out)
}
