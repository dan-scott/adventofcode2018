package day4

import (
	"github.com/dan-scott/adventofcode2018/domain"
	"regexp"
	"sort"
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
	timeline := parseInput(input)

	id, minute := strategy1(timeline)

	return id * minute
}

func solvePart2() interface{} {
	timeline := parseInput(input)

	id, minute := strategy2(timeline)

	return id * minute
}

type Entry struct {
	Time    int
	Message string
}

var guardParser = regexp.MustCompile(`Guard #(\d+) `)

func (e Entry) getGuardId() (int, bool) {
	matches := guardParser.FindAllStringSubmatch(e.Message, -1)
	if len(matches) == 0 {
		return 0, false
	}

	id, _ := strconv.ParseInt(matches[0][1], 10, 64)

	return int(id), true
}

var parser = regexp.MustCompile(`\[\d\d\d\d-\d\d-\d\d \d\d:(\d\d)] (.+)`)

func parseInput(input string) []Entry {
	rows := strings.Split(input, "\n")

	sort.Slice(rows, func(i, j int) bool {
		return rows[i] < rows[j]
	})

	var entries []Entry

	for _, row := range rows {
		matches := parser.FindAllStringSubmatch(row, -1)
		time, _ := strconv.ParseInt(matches[0][1], 10, 64)
		entries = append(entries, Entry{
			Time:    int(time),
			Message: matches[0][2],
		})
	}

	return entries
}

type guardStats struct {
	total int
	sleep map[int]int
}

func strategy1(timeline []Entry) (int, int) {

	guards := mapGuards(timeline)

	guardId := 0
	maxTime := 0
	for id, guard := range guards {
		if guard.total > maxTime {
			maxTime = guard.total
			guardId = id
		}
	}

	minute := 0
	maxTime = 0
	for t, l := range guards[guardId].sleep {
		if l > maxTime {
			maxTime = l
			minute = t
		}
	}

	return guardId, minute
}

func strategy2(timeline []Entry) (int, int) {

	guards := mapGuards(timeline)

	guardId := 0
	minute := 0
	maxTime := 0
	for id, guard := range guards {
		for t, l := range guard.sleep {
			if l > maxTime {
				maxTime = l
				minute = t
				guardId = id
			}
		}
	}

	return guardId, minute
}

func mapGuards(timeline []Entry) map[int]*guardStats {
	guards := map[int]*guardStats{}

	guardId := 0
	sleepStart := 0

	for _, entry := range timeline {
		if id, isId := entry.getGuardId(); isId {
			guardId = id
			if _, ok := guards[id]; !ok {
				guards[id] = &guardStats{
					total: 0,
					sleep: map[int]int{},
				}
			}
		} else if entry.Message == "falls asleep" {
			sleepStart = entry.Time
		} else if entry.Message == "wakes up" {
			duration := entry.Time - sleepStart
			guard := guards[guardId]
			guard.total += duration

			for i := sleepStart; i < entry.Time; i++ {
				if _, ok := guard.sleep[i]; !ok {
					guard.sleep[i] = 1
				} else {
					guard.sleep[i]++
				}
			}
		}
	}

	return guards
}
