package main

import (
	"fmt"

	"github.com/dan-scott/adventofcode2018/days/day1"
	"github.com/dan-scott/adventofcode2018/days/day2"
	"github.com/dan-scott/adventofcode2018/days/day3"
	"github.com/dan-scott/adventofcode2018/days/day4"
	"github.com/dan-scott/adventofcode2018/days/day5"
	"github.com/dan-scott/adventofcode2018/days/day6"
	"github.com/dan-scott/adventofcode2018/domain"
)

var solvers = []domain.Day{
	day1.New(),
	day2.New(),
	day3.New(),
	day4.New(),
	day5.New(),
	day6.New(),
}

func main() {
	for i, s := range solvers {
		fmt.Printf("Day %d\n", i+1)
		if s.Part1 != nil {
			fmt.Printf("\tSolving part 1... ")
			r := s.Part1()
			fmt.Println(r)
		}
		if s.Part2 != nil {
			fmt.Printf("\tSolving part 2... ")
			r := s.Part2()
			fmt.Println(r)
		}
		if s.Part1 == nil && s.Part2 == nil {
			fmt.Println("\tNo parts!")
		}
		fmt.Print("\n\n")
	}
}
