// atsp to pakiet, który zawiera implementacje algorytmów rozwiązujących problem ATSP (Asymmetric Travelling Salesman Problem).
package atsp

import "github.com/Ite-2022-pwr/sem5-pea-proj1-ak/graph"

// ATSP to interfejs, który muszą implementować wszystkie algorytmy rozwiązujące problem ATSP
type ATSP interface {
	Solve(startVertex int) (int, []int)
	GetGraph() graph.Graph
}
