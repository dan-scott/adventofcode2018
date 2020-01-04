package domain

type Solver func() interface{}

type Day struct {
	Part1 Solver
	Part2 Solver
}
