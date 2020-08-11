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
	loc, size := findMaxDynamic(18)
	return fmt.Sprintf("%v,%v,%v", loc.x, loc.y, size)
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

func (v vec2) add(x, y int) vec2 {
	return vec2{v.x + x, v.y + y}
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

func findMaxDynamic(serial int) (vec2, int) {
	grid := calcGrid(serial)
	max := math.MinInt32
	maxSize := 0
	var maxLoc vec2

	for x := 1; x <= 299; x++ {
		for y := 1; y <= 299; y++ {
			tl := vec2{x, y}
			ms := int(math.Min(float64(300-x), float64(300-y)))
			val, size := calcPower(grid, tl, ms)
			if val > max {
				max = val
				maxLoc = tl
				maxSize = size
			}
		}
	}

	return maxLoc, maxSize
}

func calcPower(grid map[vec2]int, pos vec2, maxSize int) (int, int) {
	max := math.MinInt32
	ms := 0
	val := grid[pos]
	for s := 1; s <= maxSize; s++ {
		for p := 0; p < s; p++ {
			val += grid[pos.add(s, p)] + grid[pos.add(p, s)]
		}
		val += grid[pos.add(s, s)]
		if val > max {
			max = val
			ms = s
		}
	}

	return max, ms + 1
}

func printGrid(grid map[vec2]int) {
	for y := 1; y <= 300; y++ {
		fmt.Println()
		for x := 1; x <= 300; x++ {
			fmt.Printf("% d", grid[vec2{x, y}])
		}
	}
}
