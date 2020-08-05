package day8

import (
	"strconv"
	"strings"

	"github.com/dan-scott/adventofcode2018/domain"
)

// New day 8
func New() domain.Day {
	return domain.Day{
		Part1: solvePart1,
		Part2: solvePart2,
	}
}

func solvePart1() interface{} {
	return sumMetadata(input)
}

func solvePart2() interface{} {
	return ""
}

type node struct {
	children []*node
	metadata []int
}

func (n *node) sum() int {
	total := 0

	for _, v := range n.metadata {
		total += v
	}

	for _, c := range n.children {
		total += c.sum()
	}

	return total
}

func sumMetadata(input string) int {
	root := parseInput(input)

	return root.sum()
}

func parseInput(input string) *node {
	nums := parseNumbers(input)
	root, _ := parseNode(nums)
	return root
}

func parseNumbers(input string) []int {
	parts := strings.Split(input, " ")
	nums := make([]int, len(parts), len(parts))
	for i, s := range parts {
		val, _ := strconv.ParseInt(s, 10, 32)
		nums[i] = int(val)
	}
	return nums
}

func parseNode(nums []int) (*node, int) {

	nodeCt := nums[0]
	mdCt := nums[1]
	len := 2
	children := make([]*node, 0, 0)

	for i := 0; i < nodeCt; i++ {
		child, childLen := parseNode(nums[len:])
		children = append(children, child)
		len += childLen
	}

	newNode := &node{
		children: children,
		metadata: nums[len : len+mdCt],
	}

	len += mdCt

	return newNode, len
}
