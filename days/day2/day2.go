package day2

import (
	"github.com/dan-scott/adventofcode2018/domain"
	"sort"
	"strings"
)

func New() domain.Day {
	return domain.Day{
		Part1: solvePart1,
		Part2: solvePart2,
	}
}

func solvePart1() interface{} {
	ct2, ct3 := 0, 0
	for _, id := range strings.Split(input, "\n") {
		has2, has3 := checkId(id)
		ct2 += has2
		ct3 += has3
	}

	return ct2 * ct3
}

func solvePart2() interface{} {
	return findAdjIds(input)
}

func findAdjIds(input string) string {
	ids := strings.Split(input, "\n")
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})

	offset := 2
	currentIdx := 0
Main:
	for offset < len(ids[0]) {
		pfx := ids[currentIdx][:offset]
		pst := ids[currentIdx][offset+1:]
		for i := currentIdx + 1; i < len(ids); i++ {
			if pfx != ids[i][:offset] {
				currentIdx++
				break
			}
			if pst == ids[i][offset+1:] {
				break Main
			}
		}

		if currentIdx == len(ids)-1 {
			currentIdx = 0
			offset++
		}
	}

	return ids[currentIdx][:offset] + ids[currentIdx][offset+1:]
}

func checkId(id string) (int, int) {
	cts := map[rune]int{}
	for _, r := range id {
		if _, ok := cts[r]; ok {
			cts[r]++
		} else {
			cts[r] = 1
		}
	}
	has2, has3 := 0, 0
	for _, c := range cts {
		if c == 2 {
			has2 = 1
			if has3 == 1 {
				break
			}
		} else if c == 3 {
			has3 = 1
			if has2 == 1 {
				break
			}
		}
	}

	return has2, has3
}
