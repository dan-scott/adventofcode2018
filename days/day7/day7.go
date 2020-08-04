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
	return getParalellExecTime(input, 5, 60)
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
		g.nodes.add(graphNode)
	}

	return graphNode
}

func (g *graph) popRoot() *node {
	n := g.roots.first()
	g.removeRootNode(n.id)
	return n
}

func (g *graph) removeRootNode(id string) {
	n, ok := g.roots.get(id)
	if !ok {
		return
	}

	for _, child := range n.children.nodes {
		child.removeParent(n)
	}

	g.roots.remove(n)
	g.nodes.remove(n)
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

func (e *edges) add(nodes ...*node) {
	for _, n := range nodes {
		e.nodes[n.id] = n
	}
	e.updateOrder()
}

func (e *edges) remove(n *node) {
	delete(e.nodes, n.id)
	e.updateOrder()
}

func (e *edges) get(id string) (*node, bool) {
	n, ok := e.nodes[id]
	return n, ok
}

func (e *edges) listInOrder() []*node {
	nodes := make([]*node, len(e.order), len(e.order))
	for i, id := range e.order {
		nodes[i] = e.nodes[id]
	}
	return nodes
}

func (e *edges) any() bool {
	return len(e.nodes) > 0
}

func (e *edges) first() *node {
	return e.nodes[e.order[0]]
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
	n.children.add(c)
	c.parents.add(n)
	n.graph.roots.remove(c)
}

func (n *node) removeParent(p *node) {
	n.parents.remove(p)
	p.children.remove(n)
	if !n.parents.any() {
		n.graph.roots.add(n)
	}
}

func (n *node) time() int {
	return int(n.id[0]-'A') + 1
}

func getExecOrder(input string) string {
	nodeGraph := parseInput(input)

	order := ""

	for nodeGraph.roots.any() {
		top := nodeGraph.popRoot()
		order = fmt.Sprintf("%s%s", order, top.id)
	}

	return order
}

type work struct {
	remaining int
	nodeID    string
}

func newWork(n *node, stepTime int) work {
	remaining := int(n.id[0]-'A') + 1
	return work{
		remaining: remaining,
		nodeID:    n.id,
	}
}

func getParalellExecTime(input string, workerCount, stepTime int) int {
	nodeGraph := parseInput(input)
	workers := make(map[string]int)

	time := 0

	for nodeGraph.nodes.any() || len(workers) > 0 {
		if len(workers) > 0 {
			time++
		}
		for nodeID := range workers {
			workers[nodeID]--
			if workers[nodeID] == 0 {
				delete(workers, nodeID)
				nodeGraph.removeRootNode(nodeID)
			}
		}
		availableCount := workerCount - len(workers)
		roots := make([]*node, 0, 0)
		for _, n := range nodeGraph.roots.listInOrder() {
			_, inProgress := workers[n.id]
			if !inProgress {
				roots = append(roots, n)
			}
		}

		for availableCount > 0 && len(roots) > 0 {
			availableCount--
			workers[roots[0].id] = roots[0].time() + stepTime
			roots = roots[1:]
		}
	}

	return time
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
		nodeGraph.roots.add(parentNode)
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
