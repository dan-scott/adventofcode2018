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

type marble struct {
	next  *marble
	prev  *marble
	value int
}

func (m *marble) insertNew(value int) *marble {
	n := &marble{
		next:  m.next,
		prev:  m,
		value: value,
	}
	m.next.prev = n
	m.next = n
	return n
}

func (m *marble) remove() *marble {
	next := m.next
	m.prev.next = m.next
	m.next.prev = m.prev
	m.next = nil
	m.prev = nil
	return next
}

type marbleCircle []int

func getWinningScore(players, worth int) int {

	score := make([]int, players, players)
	for p := 0; p < players; p++ {
		score[p] = 0
	}

	marble := &marble{
		value: 0,
	}
	marble.next = marble
	marble.prev = marble

	value := 0
	player := 0

	for value <= worth {
		value++
		if value%23 == 0 {
			score[player] += value
			marble = marble.prev.prev.prev.prev.prev.prev.prev
			score[player] += marble.value
			marble = marble.remove()
		} else {
			marble = marble.next.insertNew(value)
		}
		player = (player + 1) % players
	}

	sort.Ints(score)

	return score[len(score)-1]
}
