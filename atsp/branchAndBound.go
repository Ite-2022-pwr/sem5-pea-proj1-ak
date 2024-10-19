package atsp

import "github.com/Ite-2022-pwr/sem5-pea-proj1-ak/graph"

type BranchAndBoundSolver struct {
	graph graph.Graph
}

func (atsp *BranchAndBoundSolver) GetGraph() graph.Graph {
	return atsp.graph
}

func NewBranchAndBoundSolver(g graph.Graph) *BranchAndBoundSolver {
	return &BranchAndBoundSolver{graph: g}
}

func (atsp *BranchAndBoundSolver) Solve(startVertex int) (int, []int) {
	return atsp.BranchAndBound(startVertex)
}

type Node struct {
	Vertex        int
	Level         int
	Path          []int
	Cost          int
	MatrixReduced [][]int
}

func (atsp *BranchAndBoundSolver) BranchAndBound(startVertex int) (int, []int) {
	return 0, nil
}
