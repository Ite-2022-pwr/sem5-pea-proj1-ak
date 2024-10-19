package atsp

import "github.com/Ite-2022-pwr/sem5-pea-proj1-ak/graph"

type ATSP interface {
	Solve(startVertex int) (int, []int)
	GetGraph() graph.Graph
}
