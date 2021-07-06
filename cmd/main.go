package main

import (
	"flag"
	"fmt"
	day01 "github.com/doniacld/adventofcode/puzzles/2020/01"
	day02 "github.com/doniacld/adventofcode/puzzles/2020/02"
	day03 "github.com/doniacld/adventofcode/puzzles/2020/03"
	day04 "github.com/doniacld/adventofcode/puzzles/2020/04"
	day05 "github.com/doniacld/adventofcode/puzzles/2020/05"
	day06 "github.com/doniacld/adventofcode/puzzles/2020/06"
	day07 "github.com/doniacld/adventofcode/puzzles/2020/07"
	"github.com/doniacld/adventofcode/puzzles/solver"
	"log"
)

const (
	flagParamYear = "year"
	flagParamDay  = "day"

	defaultYear = 2020

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

	switch *day {
	case 1:
		s = day01.New("./puzzles/2020/01/input.txt")
	case 2:
		s = day02.New("./puzzles/2020/02/input.txt")
	case 3:
		s = day03.New("./puzzles/2020/03/input.txt", [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}})
	case 4:
		s = day04.New("./puzzles/2020/04/input.txt")
	case 5:
		s = day05.New("./puzzles/2020/05/input.txt")
	case 6:
		s = day06.New("./puzzles/2020/06/input.txt")
	case 7:
		s = day07.New("./puzzles/2020/07/input.txt", "shiny gold")
	}

	out, err := s.Solve()
	if err != nil {
		log.Fatalf("Something wrong happened while computing day %d: %s", *day, err.Error())
	}

	fmt.Printf(">>>>>>>>>>>>>> %s <<<<<<<<<<<<<<<\n", out)
}
