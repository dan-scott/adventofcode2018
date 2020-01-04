package day1

import (
	"github.com/dan-scott/adventofcode2018/domain"
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
	frequency := 0
	for _, s := range parseInput() {
		frequency += s
	}

	return frequency
}

func solvePart2() interface{} {
	freq := 0
	freqMap := map[int]interface{}{}
	idx := 0
	inputs := parseInput()
	for {
		freq += inputs[idx]
		if _, ok := freqMap[freq]; ok {
			return freq
		}
		freqMap[freq] = nil
		idx = (idx + 1) % len(inputs)
	}
}

func parseInput() []int {
	var intList []int

	for _, s := range strings.Split(input, "\n") {
		val, _ := strconv.ParseInt(s, 10, 64)
		intList = append(intList, int(val))
	}

	return intList
}
