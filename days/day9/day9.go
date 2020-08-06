package day9

import (
	"sort"

	"github.com/dan-scott/adventofcode2018/domain"
)

// New day 9
func New() domain.Day {
	return domain.Day{
		Part1: solvePart1,
		Part2: solvePart2,
	}
}

func solvePart1() interface{} {
	return getWinningScore(427, 70723)
}

func solvePart2() interface{} {
	return getWinningScore(427, 7072300)
}

type marbleCircle []int

func getWinningScore(players, worth int) int {
	marbles := make([]int, worth+1, worth+1)
	for w := 0; w < worth+1; w++ {
		marbles[w] = w
	}

	score := make([]int, players, players)
	for p := 0; p < players; p++ {
		score[p] = 0
	}

	circle := make([]int, 1, 1)

	circle[0] = marbles[0]
	marbles = marbles[1:]

	current := 0
	player := 0

	for len(marbles) > 0 {
		marble := marbles[0]

		if marble%23 == 0 {
			score[player] += marble
			current = (current - 7 + len(circle)) % len(circle)
			score[player] += circle[current]
			circle = append(circle[0:current], circle[current+1:]...)

		} else {

			nextIdx := (current + 2) % len(circle)

			if nextIdx == 0 {
				circle = append(circle, marble)
				current = len(circle) - 1
			} else {
				circle = append(circle, 0)
				copy(circle[nextIdx+1:], circle[nextIdx:])
				circle[nextIdx] = marble
				current = nextIdx
			}

		}

		marbles = marbles[1:]
		player = (player + 1) % players
	}

	sort.Ints(score)

	return score[len(score)-1]
}
