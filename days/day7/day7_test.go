package day7

import (
	"testing"
)

func TestExecOrder(t *testing.T) {
	input := `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

	order := "CABDFE"

	actual := getExecOrder(input)

	if order != actual {
		t.Fatalf("Expected to get order %s but got %s", order, actual)
	}
}

func TestParalellExcecTime(t *testing.T) {
	input = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

	workers := 2

	step := 0

	time := 15

	actual := getParalellExecTime(input, workers, step)

	if time != actual {
		t.Fatalf("Expected to take %v but took %v", time, actual)
	}

}
