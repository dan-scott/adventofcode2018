package day7

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/dan-scott/adventofcode2018/domain"
)

// New Day 7
func New() domain.Day {
	return domain.Day{
		Part1: solvePart1,
		Part2: solvePart2,
	}
}

func solvePart1() interface{} {
	return getExecOrder(input)
}

func solvePart2() interface{} {
	return ""
}

type nodeMap map[string]*node

func (m nodeMap) getInOrder() []*node {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	nodes := make([]*node, len(keys), len(keys))
	for i, key := range keys {
		node, _ := m[key]
		nodes[i] = node
	}

	return nodes
}

type graph struct {
	roots nodeMap
	nodes nodeMap
}

func (g *graph) getNode(id string) *node {
	graphNode, ok := g.nodes[id]
	if !ok {
		graphNode = &node{
			id:       id,
			parents:  make(map[string]*node),
			children: make(map[string]*node),
		}
		g.nodes[id] = graphNode
	}

	return graphNode
}

type node struct {
	id       string
	parents  nodeMap
	children nodeMap
}

func getExecOrder(input string) string {
	nodeGraph := &graph{
		roots: make(map[string]*node),
		nodes: make(map[string]*node),
	}
	ordering := parseInput(input)
	for parent := range ordering {
		parentNode := nodeGraph.getNode(parent)
		nodeGraph.roots[parent] = parentNode
	}

	for parent, children := range ordering {
		parentNode := nodeGraph.getNode(parent)
		for _, child := range children {
			childNode := nodeGraph.getNode(child)
			parentNode.children[child] = childNode
			childNode.parents[parent] = parentNode
			_, isChildInRoot := nodeGraph.roots[child]
			if isChildInRoot {
				delete(nodeGraph.roots, child)
			}
		}
	}

	stack := nodeGraph.roots.getInOrder()
	var top *node

	order := ""

	for len(stack) > 0 {
		top, stack = stack[0], stack[1:]
		if len(top.parents) > 0 || strings.Contains(order, top.id) {
			continue
		}
		order = fmt.Sprintf("%s%s", order, top.id)

		for _, child := range top.children {
			delete(child.parents, top.id)
		}

		stack = append(top.children.getInOrder(), stack...)

		sort.Slice(stack, func(i, j int) bool {
			return stack[i].id < stack[j].id
		})
	}

	return order
}

var parser = regexp.MustCompile(`Step (\w) must be finished before step (\w) can begin.`)

func parseInput(input string) map[string][]string {
	lines := strings.Split(input, "\n")

	instructions := make(map[string][]string)

	for _, line := range lines {
		matches := parser.FindAllStringSubmatch(line, -1)
		a := matches[0][1]
		b := matches[0][2]

		aList, ok := instructions[a]
		if !ok {
			aList = make([]string, 0, 1)
		}

		instructions[a] = append(aList, b)
	}

	return instructions
}
