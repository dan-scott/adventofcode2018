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

type graph struct {
	roots edges
	nodes edges
}

func newGraph() *graph {
	return &graph{
		nodes: newEdges(),
		roots: newEdges(),
	}
}

func (g *graph) getNode(id string) *node {
	graphNode, ok := g.nodes.get(id)
	if !ok {
		graphNode = newNode(id, g)
		g.nodes.addNode(graphNode)
	}

	return graphNode
}

type edges struct {
	nodes nodeMap
	order []string
}

func newEdges() edges {
	return edges{
		nodes: make(map[string]*node),
		order: make([]string, 0, 0),
	}
}

func (e *edges) updateOrder() {
	e.order = make([]string, 0, len(e.nodes))
	for id := range e.nodes {
		e.order = append(e.order, id)
	}
	sort.Strings(e.order)
}

func (e *edges) addNode(nodes ...*node) {
	for _, n := range nodes {
		e.nodes[n.id] = n
	}
	e.updateOrder()
}

func (e *edges) removeNode(n *node) {
	delete(e.nodes, n.id)
	e.updateOrder()
}

func (e *edges) get(id string) (*node, bool) {
	n, ok := e.nodes[id]
	return n, ok
}

func (e *edges) getInOrder() []*node {
	nodes := make([]*node, len(e.order), len(e.order))
	for i, id := range e.order {
		nodes[i] = e.nodes[id]
	}
	return nodes
}

func (e *edges) any() bool {
	return len(e.nodes) > 0
}

func (e *edges) pop() (*node, bool) {
	if !e.any() {
		return nil, false
	}
	n := e.nodes[e.order[0]]
	delete(e.nodes, n.id)
	e.updateOrder()
	return n, true
}

type node struct {
	id       string
	parents  edges
	children edges
	graph    *graph
}

func newNode(id string, graph *graph) *node {
	return &node{
		id:       id,
		parents:  newEdges(),
		children: newEdges(),
		graph:    graph,
	}
}

func (n *node) addChild(c *node) {
	n.children.addNode(c)
	c.parents.addNode(n)
	n.graph.roots.removeNode(c)
}

func (n *node) removeParent(p *node) {
	n.parents.removeNode(p)
	p.children.removeNode(n)
}

func getExecOrder(input string) string {
	nodeGraph := parseInput(input)

	stack := nodeGraph.roots.getInOrder()
	var top *node

	order := ""

	for len(stack) > 0 {
		top, stack = stack[0], stack[1:]
		if top.parents.any() || strings.Contains(order, top.id) {
			continue
		}
		order = fmt.Sprintf("%s%s", order, top.id)

		stack = append(top.children.getInOrder(), stack...)

		for _, child := range top.children.nodes {
			child.removeParent(top)
		}

		sort.Slice(stack, func(i, j int) bool {
			return stack[i].id < stack[j].id
		})
	}

	return order
}

var parser = regexp.MustCompile(`Step (\w) must be finished before step (\w) can begin.`)

func parseInput(input string) *graph {
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

	nodeGraph := newGraph()

	for parent := range instructions {
		parentNode := nodeGraph.getNode(parent)
		nodeGraph.roots.addNode(parentNode)
	}

	for parent, children := range instructions {
		parentNode := nodeGraph.getNode(parent)
		for _, child := range children {
			childNode := nodeGraph.getNode(child)
			parentNode.addChild(childNode)
		}
	}

	return nodeGraph
}
