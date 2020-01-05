package day3

import (
	"github.com/dan-scott/adventofcode2018/domain"
	"regexp"
	"strconv"
	"strings"
)

func New() domain.Day {
	return domain.Day{
		Part1: solvePart1,
		Part2: solvePart2,
	}
}

func solvePart1() interface{} {
	cuts := parseInput(input)
	squares := countSquares(cuts)

	ct := 0
	for _, c := range squares {
		if c > 1 {
			ct++
		}
	}
	return ct
}

func solvePart2() interface{} {
	cuts := parseInput(input)
	squares := countSquares(cuts)

Main:
	for _, c := range cuts {
		for x := c.Pos.X; x < c.Pos.X+c.W; x++ {
			for y := c.Pos.Y; y < c.Pos.Y+c.H; y++ {
				pos := Vec2d{x, y}
				if squares[pos] > 1 {
					continue Main
				}
			}
		}
		return c.Id
	}

	panic("nuts")
}

func countSquares(cuts []Cut) map[Vec2d]int {
	squares := map[Vec2d]int{}
	for _, c := range cuts {
		for x := c.Pos.X; x < c.Pos.X+c.W; x++ {
			for y := c.Pos.Y; y < c.Pos.Y+c.H; y++ {
				pos := Vec2d{x, y}
				if _, ok := squares[pos]; ok {
					squares[pos]++
				} else {
					squares[pos] = 1
				}
			}
		}
	}
	return squares
}

type Cut struct {
	Id   string
	Pos  Vec2d
	W, H int
}

type Vec2d struct {
	X, Y int
}

var reg = regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

func parseInput(input string) []Cut {
	var cuts []Cut
	for _, row := range strings.Split(input, "\n") {
		matches := reg.FindAllStringSubmatch(row, -1)
		cuts = append(cuts, Cut{
			Id: matches[0][1],
			Pos: Vec2d{
				X: parseInt(matches[0][2]),
				Y: parseInt(matches[0][3]),
			},
			W: parseInt(matches[0][4]),
			H: parseInt(matches[0][5]),
		})
	}
	return cuts
}

func parseInt(s string) int {
	n, _ := strconv.ParseInt(s, 10, 64)
	return int(n)
}
