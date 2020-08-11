package day11

import (
	"fmt"
	"math"

	"github.com/dan-scott/adventofcode2018/domain"
)

// New day 11
func New() domain.Day {
	return domain.Day{
		Part1: solvePart1,
		Part2: solvePart2,
	}
}

func solvePart1() interface{} {
	loc := findMax(7400)
	return fmt.Sprintf("%v,%v", loc.x, loc.y)
}

func solvePart2() interface{} {
	return ""
}

func powerLvl(x, y, serial int) int {

	rackID := x + 10
	powerLvl := ((y * rackID) + serial) * rackID
	powerLvl = (powerLvl % 1000) / 100

	return powerLvl - 5
}

type vec2 struct {
	x, y int
}

func calcGrid(serial int) map[vec2]int {
	grid := make(map[vec2]int)
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			grid[vec2{x, y}] = powerLvl(x, y, serial)
		}
	}
	return grid
}

func calcSubGrid(grid map[vec2]int) map[vec2]int {
	subGrid := make(map[vec2]int)
	for x := 1; x <= 297; x++ {
		for y := 1; y <= 297; y++ {
			val := grid[vec2{x, y}] + grid[vec2{x + 1, y}] + grid[vec2{x + 2, y}] + grid[vec2{x, y + 1}] + grid[vec2{x + 1, y + 1}] + grid[vec2{x + 2, y + 1}] + grid[vec2{x, y + 2}] + grid[vec2{x + 1, y + 2}] + grid[vec2{x + 2, y + 2}]
			subGrid[vec2{x, y}] = val
		}
	}
	return subGrid
}

func findMax(serial int) vec2 {
	grid := calcSubGrid(calcGrid(serial))

	max := math.MinInt32
	var maxVec vec2
	for p, v := range grid {
		if v > max {
			max = v
			maxVec = p
		}
	}
	return maxVec
}

func printGrid(grid map[vec2]int) {
	for y := 1; y <= 300; y++ {
		fmt.Println()
		for x := 1; x <= 300; x++ {
			fmt.Printf("% d", grid[vec2{x, y}])
		}
	}
}
