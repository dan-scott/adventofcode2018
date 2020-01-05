package day5

import (
	"fmt"
	"github.com/dan-scott/adventofcode2018/domain"
	"regexp"
	"strings"
	"unicode"
)

func New() domain.Day {
	return domain.Day{
		Part1: solvePart1,
		Part2: solvePart2,
	}
}

func solvePart1() interface{} {
	compressed := compress(input)

	return len(compressed)
}

func solvePart2() interface{} {
	return optimize(input)
}

func optimize(input string) int {
	minLen := len(compress(input))
	for c := 'a'; c <= 'z'; c++ {
		rStr := fmt.Sprintf("[%c%c]", c, unicode.ToUpper(c))
		reg := regexp.MustCompile(rStr)
		test := reg.ReplaceAllString(input, "")
		compressed := compress(test)
		compressedLen := len(compressed)
		if compressedLen < minLen {
			minLen = compressedLen
		}
	}

	return minLen
}

func compress(input string) string {

	current := input
	next := ""

	for {
		next = current[:1]
		for i := 1; i < len(current); i++ {
			if len(next) == 0 {
				next = current[i : i+1]
				continue
			}

			c := next[len(next)-1:]
			n := current[i : i+1]

			if c != n && strings.ToLower(c) == strings.ToLower(n) {
				next = next[:len(next)-1]
			} else {
				next += current[i : i+1]
			}

		}

		if next == current {
			break
		}
		current = next
	}

	return next
}
