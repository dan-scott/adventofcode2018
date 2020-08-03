package day6

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/dan-scott/adventofcode2018/domain"
)

// New creates a copy of Day 6
func New() domain.Day {
	return domain.Day{
		Part1: solvePart1,
		Part2: solvePart2,
	}
}

type coord struct {
	x int64
	y int64
}

func (c coord) distTo(o coord) float64 {
	return math.Abs(float64(c.x-o.x)) + math.Abs(float64(c.y-o.y))
}

type coordField struct {
	coords []coord
	tl     coord
	br     coord
	field  map[coord]int
}

func (cf *coordField) print() {
	fmt.Println()
	for y := cf.tl.y; y <= cf.br.y; y++ {
		for x := cf.tl.x; x <= cf.br.x; x++ {
			v, ok := cf.field[coord{x, y}]
			if !ok {
				fmt.Print("-")
			} else {
				fmt.Printf("%c", 'a'+v)
			}
		}
		fmt.Println()
	}
}

func solvePart1() interface{} {
	return findMaxField(input)
}

func solvePart2() interface{} {
	return findCloseField(input, 10000)
}

func findMaxField(inputStr string) int {
	cf := parseInput(inputStr)

	counts := make(map[int]int)
	for i := range cf.coords {
		counts[i] = 0
	}

	for y := cf.tl.y; y <= cf.br.y; y++ {
		for x := cf.tl.x; x <= cf.br.x; x++ {
			current := coord{x, y}

			min := math.MaxFloat64
			for i, pos := range cf.coords {

				dist := pos.distTo(current)
				if dist < min {
					min = dist
					cf.field[current] = i
				} else if dist == min {
					cf.field[current] = -1
				}

			}
		}
	}

	for p, t := range cf.field {
		if counts[t] < 0 {
			continue
		}
		if t > -1 {
			counts[t]++
		}
		if p.x == cf.tl.x || p.x == cf.br.x || p.y == cf.tl.y || p.y == cf.br.y {
			counts[t] = -1
		}
	}

	max := 0
	for _, v := range counts {
		if v > -1 && v > max {
			max = v
		}
	}

	return max
}

func findCloseField(inputStr string, maxDist int) int {
	count := 0

	cf := parseInput(inputStr)

	for x := cf.tl.x; x < cf.br.x; x++ {
		for y := cf.tl.y; y < cf.br.y; y++ {
			pos := coord{x, y}
			dist := 0
			for _, l := range cf.coords {
				dist += int(pos.distTo(l))
			}

			if dist < maxDist {
				count++
			}

		}
	}

	return count
}

func parseInput(inputStr string) *coordField {
	min := coord{math.MaxInt64, math.MaxInt64}
	max := coord{math.MinInt64, math.MinInt64}
	lines := strings.Split(inputStr, "\n")
	coords := make([]coord, len(lines), len(lines))
	for i, line := range lines {
		vals := strings.Split(line, ", ")
		x, _ := strconv.ParseInt(vals[0], 10, 64)
		y, _ := strconv.ParseInt(vals[1], 10, 64)
		if x > max.x {
			max.x = x
		}
		if x < min.x {
			min.x = x
		}
		if y > max.y {
			max.y = y
		}
		if y < min.y {
			min.y = y
		}
		coords[i] = coord{x, y}
	}

	width := max.x - min.x
	height := max.y - min.y

	tl := coord{
		min.x - width/2,
		min.y - height/2,
	}

	br := coord{
		max.x + width/2,
		max.y + height/2,
	}

	field := make(map[coord]int)

	return &coordField{coords, tl, br, field}
}
