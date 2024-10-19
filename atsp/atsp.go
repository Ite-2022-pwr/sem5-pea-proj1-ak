package atsp

import "github.com/Ite-2022-pwr/sem5-pea-proj1-ak/graph"

type ATSP struct {
	graph graph.Graph
}

func NewATSP(g graph.Graph) *ATSP {
	return &ATSP{graph: g}
}

func (a *ATSP) GetGraph() graph.Graph {
	return a.graph
}
