package day10

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/dan-scott/adventofcode2018/domain"
)

// New day 10
func New() domain.Day {
	return domain.Day{
		Part1: solvePart1,
		Part2: solvePart2,
	}
}

func solvePart1() interface{} {
	converge(input)
	return ""
}

func solvePart2() interface{} {

	return ""
}

type vec2 struct {
	x, y int
}

func (v vec2) add(o vec2) vec2 {
	return vec2{
		x: v.x + o.x,
		y: v.y + o.y,
	}
}

type light struct {
	pos, vel vec2
}

func (l light) move() light {
	return light{
		pos: l.pos.add(l.vel),
		vel: l.vel,
	}
}

func converge(input string) {
	next := parseInput(input)
	prev := make([]light, len(next), len(next))
	seconds := 0

	for true {
		next, prev = prev, next
		prevArea := area(prev)

		for i, l := range prev {
			next[i] = l.move()
		}

		nextArea := area(next)

		if nextArea > prevArea {
			fmt.Printf("\nshowed after %v seconds", seconds)
			printMap(prev)
			break
		}
		seconds++

	}

}

func area(lights []light) int {
	max := vec2{math.MinInt32, math.MinInt32}
	min := vec2{math.MaxInt32, math.MaxInt32}
	for _, l := range lights {
		if max.x < l.pos.x {
			max.x = l.pos.x
		}
		if max.y < l.pos.y {
			max.y = l.pos.y
		}
		if min.x > l.pos.x {
			min.x = l.pos.x
		}
		if min.y > l.pos.y {
			min.y = l.pos.y
		}
	}

	return (max.x - min.x) * (max.y - min.y)
}

func printMap(lights []light) {
	lightmap := make(map[vec2]interface{})
	max := vec2{math.MinInt32, math.MinInt32}
	min := vec2{math.MaxInt32, math.MaxInt32}
	for _, l := range lights {
		if max.x < l.pos.x {
			max.x = l.pos.x
		}
		if max.y < l.pos.y {
			max.y = l.pos.y
		}
		if min.x > l.pos.x {
			min.x = l.pos.x
		}
		if min.y > l.pos.y {
			min.y = l.pos.y
		}
		lightmap[l.pos] = true
	}

	fmt.Printf("\n%+v %+v", min, max)

	fmt.Println()
	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			_, has := lightmap[vec2{x, y}]
			if has {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

var parse = regexp.MustCompile(`position=<\s*(-?\d*),\s*(-?\d*)> velocity=<\s*(-?\d*),\s*(-?\d*)>`)

func parseInput(input string) []light {
	lines := strings.Split(input, "\n")
	lights := make([]light, len(lines), len(lines))
	for i, l := range lines {
		matches := parse.FindAllStringSubmatch(l, -1)
		lights[i] = light{
			pos: vec2{parseInt(matches[0][1]), parseInt(matches[0][2])},
			vel: vec2{parseInt(matches[0][3]), parseInt(matches[0][4])},
		}
	}
	return lights
}

func parseInt(input string) int {
	v, _ := strconv.ParseInt(input, 10, 32)
	return int(v)
}
